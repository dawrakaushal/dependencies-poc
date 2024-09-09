package runner

import (
	"bufio"
	"fmt"
	"github.com/kaushaldawra/example-artifacts/pkg/logger"
	"github.com/kaushaldawra/example-artifacts/pkg/supported"
	"go.uber.org/fx"
	"log/slog"
	"os"
	"slices"
	"strings"
)

type Runner struct {
	config *supported.Config
	logger *slog.Logger
}

func NewRunner(config *supported.Config, loggerProvider *logger.LoggerProvider) *Runner {
	return &Runner{
		config: config,
		logger: loggerProvider.GetConsoleLogger(),
	}
}

func Init(runner *Runner) {
	runner.Run()
}

func (r *Runner) Run() {
	r.logger.Warn("Printing supported markets")
	for _, market := range r.config.AllMarkets {
		r.logger.Info(fmt.Sprintf("Market %s is supported", market))
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		r.logger.Info("Enter market: ")
		inputMarket, _ := reader.ReadString('\n')
		inputMarket = strings.Trim(inputMarket, "\n")
		if !slices.Contains(r.config.AllMarkets, inputMarket) {
			r.logger.Error(fmt.Sprintf("Market %s is not supported", inputMarket))
		} else {
			r.logger.Info(fmt.Sprintf("Welcome to %s market", inputMarket))
		}
	}
}

var Module = fx.Options(
	fx.Provide(NewRunner),
)
