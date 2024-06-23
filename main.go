package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := getRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func getRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
