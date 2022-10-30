package extractparams

import (
	"github.com/gin-gonic/gin"
	"github.com/saranshabd/migrate-from-express/internal/mfe/middlewares"
)

func InitRoutes(app *gin.Engine) {

	// Create HTTP router group
	r := app.Group("/extractParams/")

	// Middleware to check if the incoming request is from a valid source
	r.Use(middlewares.IsVerifiedSource())

	r.POST("/", CreateExtractParams)
	r.GET("/", GetExtractParams)
	r.GET("/:specificPath/", GetSpecificExtractParams)
}
