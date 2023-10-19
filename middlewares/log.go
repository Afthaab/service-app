package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// reqKey is a type used for context keys.
type reqKey string

// RequestIDKey is a key for 'request id' context value.
const RequestIDKey reqKey = "123"

func (m *Mid) Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		uuidStr := uuid.NewString()
		ctx = context.WithValue(ctx, RequestIDKey, uuidStr)
		c.Request = c.Request.WithContext(ctx)
		log.Info().Str("Trace Id", uuidStr).Str("Method", c.Request.Method).
			Str("URL Path", c.Request.URL.Path).Msg("request started")
		// After the request is processed by the next handler, logs the info again with status code
		defer log.Info().Str("Trace Id", uuidStr).Str("Method", c.Request.Method).
			Str("URL Path", c.Request.URL.Path).
			Int("status Code", c.Writer.Status()).Msg("Request processing completed")

		//we use c.Next only when we are using r.Use() method to assign middlewares

		c.Next()
	}
}
