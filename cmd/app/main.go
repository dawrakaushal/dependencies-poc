package main

import (
	"github.com/kaushaldawra/example-artifacts/cmd/runner"
	"github.com/kaushaldawra/example-artifacts/pkg/logger"
	"github.com/kaushaldawra/example-artifacts/pkg/supported"
	"go.uber.org/fx"
)

func dependencies() []fx.Option {
	deps := []fx.Option{
		supported.Module,
		logger.Module,
		runner.Module,
	}

	return append(deps, fx.Invoke(runner.Init))
}

func main() {
	fx.New(dependencies()...).Run()
}
