package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, statusCode int, success bool, message string, data any) {
	c.JSON(statusCode, &gin.H{
		"success": success,
		"message": message,
		"data":    data,
	})
}

func RespondOkay(c *gin.Context, message string, data any) {
	Respond(c, http.StatusOK, true, message, data)
}

func InvalidRequestParams(c *gin.Context) {
	Respond(c, http.StatusBadRequest, false, "Invalid request parameters", nil)
}
