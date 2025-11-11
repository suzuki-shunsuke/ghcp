package github

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v78/github"
	"github.com/suzuki-shunsuke/ghcp/pkg/git"
	"github.com/suzuki-shunsuke/ghcp/pkg/github/client/mock_client"
	testingLogger "github.com/suzuki-shunsuke/ghcp/pkg/logger/testing"
)

func TestGitHub_CreateCommit(t *testing.T) {
	ctx := context.TODO()
	repositoryID := git.RepositoryID{Owner: "owner", Name: "repo"}

	t.Run("SingleParent", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		gitHubClient := mock_client.NewMockInterface(ctrl)
		gitHubClient.EXPECT().
			CreateCommit(ctx, "owner", "repo", &github.Commit{
				Message: github.Ptr("message"),
				Parents: []*github.Commit{{SHA: github.Ptr("parentCommitSHA")}},
				Tree:    &github.Tree{SHA: github.Ptr("treeSHA")},
			}, nil).
			Return(&github.Commit{
				SHA: github.Ptr("commitSHA"),
			}, nil, nil)
		gitHub := GitHub{
			Client: gitHubClient,
			Logger: testingLogger.New(t),
		}
		commitSHA, err := gitHub.CreateCommit(ctx, git.NewCommit{
			Repository:      repositoryID,
			Message:         "message",
			ParentCommitSHA: "parentCommitSHA",
			TreeSHA:         "treeSHA",
		})
		if err != nil {
			t.Fatalf("CreateCommit returned error: %+v", err)
		}
		if commitSHA != "commitSHA" {
			t.Errorf("commitSHA wants commitSHA but %s", commitSHA)
		}
	})

	t.Run("NoParent", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		gitHubClient := mock_client.NewMockInterface(ctrl)
		gitHubClient.EXPECT().
			CreateCommit(ctx, "owner", "repo", &github.Commit{
				Message: github.Ptr("message"),
				Tree:    &github.Tree{SHA: github.Ptr("treeSHA")},
			}, nil).
			Return(&github.Commit{
				SHA: github.Ptr("commitSHA"),
			}, nil, nil)
		gitHub := GitHub{
			Client: gitHubClient,
			Logger: testingLogger.New(t),
		}
		commitSHA, err := gitHub.CreateCommit(ctx, git.NewCommit{
			Repository: repositoryID,
			Message:    "message",
			TreeSHA:    "treeSHA",
		})
		if err != nil {
			t.Fatalf("CreateCommit returned error: %+v", err)
		}
		if commitSHA != "commitSHA" {
			t.Errorf("commitSHA wants commitSHA but %s", commitSHA)
		}
	})
}
