package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	helloApi := r.Group("/hello")
	{
		helloApi.GET("/message", GetSample)
		helloApi.POST("/message", PostSample)
	}
	r.Run()
}
