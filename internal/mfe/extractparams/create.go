package extractparams

import (
	"github.com/gin-gonic/gin"
	"github.com/saranshabd/migrate-from-express/internal/mfe/pkg/http"
)

type CreateExtractParamsArgs struct {
	Arg1 string `json:"arg1" binding:"required"`
	Arg2 int64  `json:"arg2" binding:"required"`
	Arg3 bool   `json:"arg3"` // optional
}

func CreateExtractParams(c *gin.Context) {
	// Extract the parameters from the request.body & validate the object received
	var args CreateExtractParamsArgs
	if err := c.BindJSON(&args); err != nil {
		// Return HTTP 400 if invalid params are passed
		http.InvalidRequestParams(c)
		return
	}

	// Return the params object in response if everything is fine
	http.RespondOkay(c, "Request parameters extracted successfully", args)
}
