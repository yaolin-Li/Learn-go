package main

import (
	"github.com/gin-gonic/gin"
	"learn-go/gin-example/controllers"
	"learn-go/gin-example/middlewares"
	"log"
)

func main() {
	server := gin.Default()
	server.Use(middlewares.MyAuth())
	server.GET("/ping", func(context *gin.Context) {
		context.String(200, "%s", "pong")
	})

	server.Static("/resources","./resources")
	server.StaticFile("/myVideo","./resources/OperationShow.mp4")

	videoController := controllers.NewVideoController()
	videoGroup := server.Group("/videos")
	videoGroup.Use(middlewares.MyLogger())

	videoGroup.GET("/", videoController.GetAll)
	videoGroup.PUT("/:id", videoController.Update)
	videoGroup.POST("/", videoController.Create)
	videoGroup.DELETE("/:id", videoController.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}