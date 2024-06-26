package cmd

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/suzuki-shunsuke/ghcp/pkg/git"
	"github.com/suzuki-shunsuke/ghcp/pkg/github/client"
	"github.com/suzuki-shunsuke/ghcp/pkg/logger"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/pullrequest"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/pullrequest/mock_pullrequest"
)

func TestCmd_Run_pull_request(t *testing.T) {
	ctx := context.TODO()

	t.Run("BasicOptions", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := mock_pullrequest.NewMockInterface(ctrl)
		useCase.EXPECT().
			Do(ctx, pullrequest.Input{
				HeadRepository: git.RepositoryID{Owner: "owner", Name: "repo"},
				HeadBranchName: "feature",
				BaseRepository: git.RepositoryID{Owner: "owner", Name: "repo"},
				Title:          "commit-message",
			})
		r := Runner{
			NewLogger:         newLogger(t, logger.Option{}),
			NewGitHub:         newGitHub(t, client.Option{Token: "YOUR_TOKEN"}),
			Env:               newEnv(ctrl, map[string]string{envGitHubAPI: ""}),
			NewInternalRunner: newInternalRunner(InternalRunner{PullRequestUseCase: useCase}),
		}
		args := []string{
			cmdName,
			pullRequestCmdName,
			"--token", "YOUR_TOKEN",
			"-r", "owner/repo",
			"-b", "feature",
			"--title", "commit-message",
		}
		exitCode := r.Run(ctx, args, version)
		if exitCode != exitCodeOK {
			t.Errorf("exitCode wants %d but %d", exitCodeOK, exitCode)
		}
	})

	t.Run("--base-repo", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := mock_pullrequest.NewMockInterface(ctrl)
		useCase.EXPECT().
			Do(ctx, pullrequest.Input{
				HeadRepository: git.RepositoryID{Owner: "owner", Name: "repo"},
				HeadBranchName: "feature",
				BaseRepository: git.RepositoryID{Owner: "upstream-owner", Name: "upstream-repo"},
				Title:          "commit-message",
			})
		r := Runner{
			NewLogger:         newLogger(t, logger.Option{}),
			NewGitHub:         newGitHub(t, client.Option{Token: "YOUR_TOKEN"}),
			Env:               newEnv(ctrl, map[string]string{envGitHubAPI: ""}),
			NewInternalRunner: newInternalRunner(InternalRunner{PullRequestUseCase: useCase}),
		}
		args := []string{
			cmdName,
			pullRequestCmdName,
			"--token", "YOUR_TOKEN",
			"-r", "owner/repo",
			"--base-repo", "upstream-owner/upstream-repo",
			"-b", "feature",
			"--title", "commit-message",
		}
		exitCode := r.Run(ctx, args, version)
		if exitCode != exitCodeOK {
			t.Errorf("exitCode wants %d but %d", exitCodeOK, exitCode)
		}
	})

	t.Run("optional-flags", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		useCase := mock_pullrequest.NewMockInterface(ctrl)
		useCase.EXPECT().
			Do(ctx, pullrequest.Input{
				HeadRepository: git.RepositoryID{Owner: "owner", Name: "repo"},
				HeadBranchName: "feature",
				BaseRepository: git.RepositoryID{Owner: "upstream-owner", Name: "upstream-repo"},
				BaseBranchName: "develop",
				Title:          "commit-message",
				Body:           "body",
				Reviewer:       "the-reviewer",
				Draft:          true,
			})
		r := Runner{
			NewLogger:         newLogger(t, logger.Option{}),
			NewGitHub:         newGitHub(t, client.Option{Token: "YOUR_TOKEN"}),
			Env:               newEnv(ctrl, map[string]string{envGitHubAPI: ""}),
			NewInternalRunner: newInternalRunner(InternalRunner{PullRequestUseCase: useCase}),
		}
		args := []string{
			cmdName,
			pullRequestCmdName,
			"--token", "YOUR_TOKEN",
			"-u", "owner",
			"-r", "repo",
			"-b", "feature",
			"--base-owner", "upstream-owner",
			"--base-repo", "upstream-repo",
			"--base", "develop",
			"--title", "commit-message",
			"--body", "body",
			"--draft",
			"--reviewer", "the-reviewer",
		}
		exitCode := r.Run(ctx, args, version)
		if exitCode != exitCodeOK {
			t.Errorf("exitCode wants %d but %d", exitCodeOK, exitCode)
		}
	})
}
