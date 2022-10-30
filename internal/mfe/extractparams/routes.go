package extractparams

import "github.com/gin-gonic/gin"

func InitRoutes(app *gin.Engine) {

	// Create HTTP router group
	r := app.Group("/extractParams/")

	r.POST("/", CreateExtractParams)
	r.GET("/", GetExtractParams)
	r.GET("/:specificPath/", GetSpecificExtractParams)
}
