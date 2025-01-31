//go:build wireinject
// +build wireinject

// Package di provides dependency injection.
package di

import (
	"github.com/google/wire"
	"github.com/suzuki-shunsuke/ghcp/pkg/cmd"
	"github.com/suzuki-shunsuke/ghcp/pkg/env"
	"github.com/suzuki-shunsuke/ghcp/pkg/fs"
	"github.com/suzuki-shunsuke/ghcp/pkg/github"
	"github.com/suzuki-shunsuke/ghcp/pkg/github/client"
	"github.com/suzuki-shunsuke/ghcp/pkg/logger"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/commit"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/forkcommit"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/gitobject"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/pullrequest"
	"github.com/suzuki-shunsuke/ghcp/pkg/usecases/release"
)

func NewCmd() cmd.Interface {
	wire.Build(
		cmd.Set,
		logger.Set,
		client.Set,
		env.Set,

		wire.Value(cmd.NewInternalRunnerFunc(NewCmdInternalRunner)),
	)
	return nil
}

func NewCmdInternalRunner(logger.Interface, client.Interface) *cmd.InternalRunner {
	wire.Build(
		cmd.Set,
		fs.Set,
		github.Set,

		gitobject.Set,
		commit.Set,
		forkcommit.Set,
		pullrequest.Set,
		release.Set,
	)
	return nil
}
