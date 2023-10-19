package handlers

import (
	"github.com/rs/zerolog/log"

	"net/http"

	"github.com/afthaab/service-app/auth"
	"github.com/afthaab/service-app/middlewares"
	"github.com/gin-gonic/gin"
)

func API(a *auth.Auth) *gin.Engine {

	router := gin.New()

	m, err := middlewares.NewMid(a)
	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	router.Use(gin.Recovery(), m.Log())

	router.GET("/check", m.Authenticate(CheckPoint))
	return router
}

func CheckPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello status ok",
	})
	panic(":hello")

}
