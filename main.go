package main

import (
	"os"

	"github.com/suzuki-shunsuke/ghcp/pkg/di"
)

var version = "HEAD"

func main() {
	os.Exit(di.NewCmd().Run(os.Args, version))
}
