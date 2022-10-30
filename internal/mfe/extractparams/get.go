package extractparams

import (
	"github.com/gin-gonic/gin"
	"github.com/saranshabd/migrate-from-express/internal/mfe/pkg/http"
)

type GetExtractParamsArgs struct {
	Arg1 string `form:"arg1" json:"arg1" binding:"required"`
	Arg2 int64  `form:"arg2" json:"arg2" binding:"required"`
	Arg3 bool   `form:"arg3" json:"arg3,omitempty"` // optional
}

func GetExtractParams(c *gin.Context) {
	// Extract the parameters from the request.query & validate the object received
	var args GetExtractParamsArgs
	if err := c.BindQuery(&args); err != nil {
		// Return HTTP 400 if invalid params are passed
		http.InvalidRequestParams(c)
		return
	}

	// Return the params object in response if everything if fine
	http.RespondOkay(c, "Request parameres extracted successfully", args)
}
