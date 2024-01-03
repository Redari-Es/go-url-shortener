package main

import (
	"fmt"
	"go-url-shortener/handler"
	"go-url-shortener/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			// "message": "Hey Go URL Shortener !",
			"message": "Welcome to the URL Shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandlerShortUrlRedirect(c)
	})
	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(":3030")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: % v", err))
	}
	fmt.Printf("Hello Go URL Shortener!🚀")
}
