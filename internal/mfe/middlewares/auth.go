package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/saranshabd/migrate-from-express/internal/mfe/pkg/http"
)

type VerifiedSourcesArgs struct {
	VerifiedSource string `header:"verified-source" binding:"required"`
}

const (
	AcceptedSource = "accepted-source"
)

func IsVerifiedSource() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the header value from the HTTP request
		var args VerifiedSourcesArgs
		if err := c.ShouldBindHeader(&args); err != nil {
			http.InvalidHeaders(c)
			return
		}

		// Check if the source of the incoming request is accepted
		if args.VerifiedSource != AcceptedSource {
			http.InvalidHeaders(c)
			return
		}

		// Continue with the chain of middlewares/handlers
		c.Next()
	}
}
