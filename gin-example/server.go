package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(context *gin.Context) {
		context.String(200, "%s", "pong")
	})

	server.Static("/resources","./resources")

	server.StaticFile("/myVideo","./resources/OperationShow.mp4")

	log.Fatalln(server.Run("localhost:8080"))
}