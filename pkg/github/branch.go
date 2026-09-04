package github

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v7"
	"github.com/shurcooL/githubv4"
	"github.com/suzuki-shunsuke/ghcp/pkg/git"
)

type QueryForCommitInput struct {
	ParentRepository git.RepositoryID
	ParentRef        git.RefName // optional
	TargetRepository git.RepositoryID
	TargetBranchName git.BranchName // optional
}

type QueryForCommitOutput struct {
	CurrentUserName              string
	ParentDefaultBranchCommitSHA git.CommitSHA
	ParentDefaultBranchTreeSHA   git.TreeSHA
	ParentRefCommitSHA           git.CommitSHA // empty if the parent ref does not exist
	ParentRefTreeSHA             git.TreeSHA   // empty if the parent ref does not exist
	TargetRepositoryNodeID       InternalRepositoryNodeID
	TargetBranchNodeID           InternalBranchNodeID
	TargetBranchCommitSHA        git.CommitSHA // empty if the branch does not exist
	TargetBranchTreeSHA          git.TreeSHA   // empty if the branch does not exist
}

func (q *QueryForCommitOutput) TargetBranchExists() bool {
	return q.TargetBranchCommitSHA != ""
}

// parentRefMaxElapsedTime is the maximum time to wait for the parent ref to become available.
const parentRefMaxElapsedTime = 30 * time.Second

// QueryForCommit returns the repository for creating or updating the branch.
// If the parent ref is given, it must exist.
// The GraphQL API is eventually consistent and a ref created moments ago is not
// always visible yet, so this waits for the ref if it is not found.
func (c *GitHub) QueryForCommit(ctx context.Context, in QueryForCommitInput) (*QueryForCommitOutput, error) {
	out, err := c.queryForCommit(ctx, in)
	if err != nil {
		return nil, err
	}
	if in.ParentRef == "" || out.ParentRefCommitSHA != "" {
		return out, nil
	}
	return c.waitUntilParentRefIsAvailable(ctx, in)
}

// waitUntilParentRefIsAvailable queries the repository until the parent ref is available.
// Without this, ghcp would create a commit which has no parent silently.
func (c *GitHub) waitUntilParentRefIsAvailable(ctx context.Context, in QueryForCommitInput) (*QueryForCommitOutput, error) {
	c.Logger.Debugf("The parent ref (%s) is not found. Waiting for the ref", in.ParentRef)
	operation := func() (*QueryForCommitOutput, error) {
		out, err := c.queryForCommit(ctx, in)
		if err != nil {
			return nil, backoff.Permanent(err)
		}
		if out.ParentRefCommitSHA == "" {
			return nil, fmt.Errorf("the parent ref (%s) is not found in the repository (%s)", in.ParentRef, in.ParentRepository)
		}
		return out, nil
	}
	out, err := backoff.Retry(ctx, operation,
		backoff.WithBackOff(backoff.NewExponentialBackOff()),
		backoff.WithMaxElapsedTime(parentRefMaxElapsedTime))
	if err != nil {
		return nil, fmt.Errorf("retry over: %w", err)
	}
	return out, nil
}

func (c *GitHub) queryForCommit(ctx context.Context, in QueryForCommitInput) (*QueryForCommitOutput, error) {
	var q struct {
		Viewer struct {
			Login string
		}

		ParentRepository struct {
			// default branch
			DefaultBranchRef struct {
				Name   string
				Target struct {
					Commit struct {
						Oid  string
						Tree struct {
							Oid string
						}
					} `graphql:"... on Commit"`
				}
			}

			// parent ref (optional)
			ParentRef struct {
				Prefix string
				Name   string
				Target struct {
					Commit struct {
						Oid  string
						Tree struct {
							Oid string
						}
					} `graphql:"... on Commit"`
				}
			} `graphql:"parentRef: ref(qualifiedName: $parentRef)"`
		} `graphql:"parentRepository: repository(owner: $parentOwner, name: $parentRepo)"`

		TargetRepository struct {
			ID  githubv4.ID
			Ref struct {
				ID     githubv4.ID
				Target struct {
					Commit struct {
						Oid  string
						Tree struct {
							Oid string
						}
					} `graphql:"... on Commit"`
				}
			} `graphql:"ref(qualifiedName: $targetRef)"`
		} `graphql:"targetRepository: repository(owner: $targetOwner, name: $targetRepo)"`
	}
	v := map[string]interface{}{
		"parentOwner": githubv4.String(in.ParentRepository.Owner),
		"parentRepo":  githubv4.String(in.ParentRepository.Name),
		"parentRef":   githubv4.String(in.ParentRef),
		"targetOwner": githubv4.String(in.TargetRepository.Owner),
		"targetRepo":  githubv4.String(in.TargetRepository.Name),
		"targetRef":   githubv4.String(in.TargetBranchName.QualifiedName().String()),
	}
	c.Logger.Debugf("Querying the repository with %+v", v)
	if err := c.Client.Query(ctx, &q, v); err != nil {
		return nil, fmt.Errorf("GitHub API error: %w", err)
	}
	c.Logger.Debugf("Got the result: %+v", q)
	out := QueryForCommitOutput{
		CurrentUserName:              q.Viewer.Login,
		ParentDefaultBranchCommitSHA: git.CommitSHA(q.ParentRepository.DefaultBranchRef.Target.Commit.Oid),
		ParentDefaultBranchTreeSHA:   git.TreeSHA(q.ParentRepository.DefaultBranchRef.Target.Commit.Tree.Oid),
		ParentRefCommitSHA:           git.CommitSHA(q.ParentRepository.ParentRef.Target.Commit.Oid),
		ParentRefTreeSHA:             git.TreeSHA(q.ParentRepository.ParentRef.Target.Commit.Tree.Oid),
		TargetRepositoryNodeID:       q.TargetRepository.ID,
		TargetBranchNodeID:           q.TargetRepository.Ref.ID,
		TargetBranchCommitSHA:        git.CommitSHA(q.TargetRepository.Ref.Target.Commit.Oid),
		TargetBranchTreeSHA:          git.TreeSHA(q.TargetRepository.Ref.Target.Commit.Tree.Oid),
	}
	c.Logger.Debugf("Returning the repository: %+v", out)
	return &out, nil
}

type CreateBranchInput struct {
	RepositoryNodeID InternalRepositoryNodeID
	BranchName       git.BranchName
	CommitSHA        git.CommitSHA
}

// CreateBranch creates a branch and returns nil or an error.
func (c *GitHub) CreateBranch(ctx context.Context, in CreateBranchInput) error {
	// https://docs.github.com/en/graphql/reference/mutations#createref
	v := githubv4.CreateRefInput{
		RepositoryID: in.RepositoryNodeID,
		Name:         githubv4.String(in.BranchName.QualifiedName().String()),
		Oid:          githubv4.GitObjectID(in.CommitSHA),
	}
	c.Logger.Debugf("Mutation createRef(%+v)", v)
	var m struct {
		CreateRef struct {
			Ref struct {
				Name string
			}
		} `graphql:"createRef(input: $input)"`
	}
	if err := c.Client.Mutate(ctx, &m, v, nil); err != nil {
		return fmt.Errorf("GitHub API error: %w", err)
	}
	c.Logger.Debugf("Got the result: %+v", m)
	return nil
}

type UpdateBranchInput struct {
	BranchRefNodeID InternalBranchNodeID
	CommitSHA       git.CommitSHA
	Force           bool
}

// UpdateBranch updates the branch and returns nil or an error.
func (c *GitHub) UpdateBranch(ctx context.Context, in UpdateBranchInput) error {
	// https://docs.github.com/en/graphql/reference/mutations#updateref
	v := githubv4.UpdateRefInput{
		RefID: in.BranchRefNodeID,
		Oid:   githubv4.GitObjectID(in.CommitSHA),
		Force: githubv4.NewBoolean(githubv4.Boolean(in.Force)),
	}
	c.Logger.Debugf("Mutation updateRef(%+v)", v)
	var m struct {
		UpdateRef struct {
			Ref struct {
				Name string
			}
		} `graphql:"updateRef(input: $input)"`
	}
	if err := c.Client.Mutate(ctx, &m, v, nil); err != nil {
		return fmt.Errorf("GitHub API error: %w", err)
	}
	c.Logger.Debugf("Got the result: %+v", m)
	return nil
}
