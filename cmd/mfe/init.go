package main

import (
	"fmt"
	"github.com/saranshabd/migrate-from-express/cmd/mfe/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/saranshabd/migrate-from-express/cmd/mfe/logger"
	"github.com/saranshabd/migrate-from-express/configs"
)

func main() {

	// Create a Gin server application with default configurations & middlewares
	app := gin.Default()

	// Load all the application HTTP endpoints
	http.InitRoutes(app)

	// Start listening to HTTP requests from the client
	logger.App.Info().Msgf("Firing it up on %d port number.", configs.Configs.Port)
	if err := app.Run(fmt.Sprintf(":%d", configs.Configs.Port)); err != nil {
		logger.App.Error().Err(err).Msg("Could not fire up Gin :/")
		os.Exit(1) // Kill the application
	}
}
