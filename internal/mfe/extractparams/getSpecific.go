package extractparams

import (
	"github.com/gin-gonic/gin"
	"github.com/saranshabd/migrate-from-express/internal/mfe/pkg/http"
)

type GetSpecificExtractParamsArgs struct {
	SpecificPath string `uri:"specificPath" json:"specificPath" binding:"required"`
}

func GetSpecificExtractParams(c *gin.Context) {
	// Extract the parameters from the request.params & validate the object received
	var args GetSpecificExtractParamsArgs
	if err := c.BindUri(&args); err != nil {
		// Return HTTP 400 if invalid params are passed
		http.InvalidRequestParams(c)
		return
	}

	// Return the params object in response if everything is fine
	http.RespondOkay(c, "Request parameters extracted successfully", args)
}
