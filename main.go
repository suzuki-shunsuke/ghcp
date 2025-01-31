package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/suzuki-shunsuke/ghcp/pkg/di"
)

var version = "HEAD"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	os.Exit(di.NewCmd().Run(ctx, os.Args, version))

}
