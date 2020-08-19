package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func someGetHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello Gin")
}

func SetupRoute() *gin.Engine {
	router := gin.Default()

	//router.GET("/someGet", func(c *gin.Context) {
	//	c.String(http.StatusOK, "Hello Golang")
	//})

	router.GET("/someGet", someGetHandler)

	return router
}
