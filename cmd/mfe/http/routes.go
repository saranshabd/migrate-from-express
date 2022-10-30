package http

import (
	"github.com/gin-gonic/gin"
	"github.com/saranshabd/migrate-from-express/internal/mfe/extractparams"
	"github.com/saranshabd/migrate-from-express/internal/mfe/middlewares"
)

func InitRoutes(app *gin.Engine) {
	// Ensure that all the HTTP responses are by-default in the JSON format
	app.Use(middlewares.JSONResponse())

	// Load all the application HTTP routes
	extractparams.InitRoutes(app)
}
