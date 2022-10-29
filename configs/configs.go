package configs

import (
	"os"
	"strconv"

	"github.com/samber/lo"
	"github.com/saranshabd/migrate-from-express/cmd/mfe/logger"
)

type configurations struct {
	Environment string
	Port        int64
}

// Different environments in which the application can run
const (
	DevEnvironment     = "Development"
	ProdEnvironment    = "Production"
	TestingEnvironment = "Testing"
)

var Environments = []string{DevEnvironment, ProdEnvironment, TestingEnvironment}

var Configs *configurations

func init() {

	// Extract & validate the `Environment` variable
	env := os.Getenv("Environment")
	if !lo.Contains(Environments, env) {
		logger.App.Error().Msg("Invalid value for `Environment` passed")
		os.Exit(1) // Kill the application
	}

	// Extract & validate the `Port` variable
	portNumber, err := strconv.ParseInt(os.Getenv("Port"), 10, 64)
	if err != nil {
		logger.App.Error().Err(err).Msg("Could not parse port number from env")
		os.Exit(1) // Kill the application
	}

	Configs = &configurations{
		Environment: env,
		Port:        portNumber,
	}
}
