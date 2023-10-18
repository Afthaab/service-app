package handlers

import (
	"net/http"

	"github.com/afthaab/service-app/middlewares"
	"github.com/gin-gonic/gin"
)

func API() *gin.Engine {

	router := gin.New()

	router.Use(gin.Recovery(), middlewares.Log())

	router.GET("/check", CheckPoint)
	return router
}

func CheckPoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Hello status ok",
	})
	panic(":hello")

}
