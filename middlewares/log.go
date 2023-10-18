package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// reqKey is a type used for context keys.
type reqKey int

// RequestIDKey is a key for 'request id' context value.
const RequestIDKey reqKey = 123

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		uuidStr := uuid.NewString()
		ctx = context.WithValue(ctx, RequestIDKey, uuidStr)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
