package cmd

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/suzuki-shunsuke/ghcp/pkg/git"
	"github.com/suzuki-shunsuke/ghcp/pkg/git/commitstrategy"
	"github.com/suzuki-shunsuke/ghcp/pkg/github/client"
	"github.com/suzuki-shunsuke/ghcp/pkg/logger"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/forkcommit"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/forkcommit/mock_forkcommit"
)

func TestCmd_Run_forkcommit(t *testing.T) {
	ctx := context.TODO()

	t.Run("BasicOptions", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		commitUseCase := mock_forkcommit.NewMockInterface(ctrl)
		commitUseCase.EXPECT().
			Do(ctx, forkcommit.Input{
				ParentRepository: git.RepositoryID{Owner: "owner", Name: "repo"},
				TargetBranchName: "topic",
				CommitStrategy:   commitstrategy.FastForward,
				CommitMessage:    "commit-message",
				Paths:            []string{"file1", "file2"},
			})
		r := Runner{
			NewLogger:         newLogger(t, logger.Option{}),
			NewGitHub:         newGitHub(t, client.Option{Token: "YOUR_TOKEN"}),
			Env:               newEnv(ctrl, map[string]string{envGitHubAPI: ""}),
			NewInternalRunner: newInternalRunner(InternalRunner{ForkCommitUseCase: commitUseCase}),
		}
		args := []string{
			cmdName,
			forkCommitCmdName,
			"--token", "YOUR_TOKEN",
			"-r", "owner/repo",
			"-b", "topic",
			"-m", "commit-message",
			"file1",
			"file2",
		}
		exitCode := r.Run(ctx, args, version)
		if exitCode != exitCodeOK {
			t.Errorf("exitCode wants %d but %d", exitCodeOK, exitCode)
		}
	})

	t.Run("--parent", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		commitUseCase := mock_forkcommit.NewMockInterface(ctrl)
		commitUseCase.EXPECT().
			Do(ctx, forkcommit.Input{
				TargetBranchName: "topic",
				ParentRepository: git.RepositoryID{Owner: "owner", Name: "repo"},
				CommitStrategy:   commitstrategy.RebaseOn("develop"),
				CommitMessage:    "commit-message",
				Paths:            []string{"file1", "file2"},
			})
		r := Runner{
			NewLogger:         newLogger(t, logger.Option{}),
			NewGitHub:         newGitHub(t, client.Option{Token: "YOUR_TOKEN"}),
			Env:               newEnv(ctrl, map[string]string{envGitHubAPI: ""}),
			NewInternalRunner: newInternalRunner(InternalRunner{ForkCommitUseCase: commitUseCase}),
		}
		args := []string{
			cmdName,
			forkCommitCmdName,
			"--token", "YOUR_TOKEN",
			"-u", "owner",
			"-r", "repo",
			"-m", "commit-message",
			"-b", "topic",
			"--parent", "develop",
			"file1",
			"file2",
		}
		exitCode := r.Run(ctx, args, version)
		if exitCode != exitCodeOK {
			t.Errorf("exitCode wants %d but %d", exitCodeOK, exitCode)
		}
	})
}
