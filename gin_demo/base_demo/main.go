package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func jsonHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "ke",
		"age":  18,
	})
}

func main01() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	router.GET("/json", jsonHandler)

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404")
	})

	router.Run()
}

func main() {

	router := SetupRoute()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
