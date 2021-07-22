package main

import (
	"io"
	"net/http"
	"os"

	"github.com/daviddonikapanjaitan/belajar-gin-gonic/controller"
	"github.com/daviddonikapanjaitan/belajar-gin-gonic/middlewares"
	"github.com/daviddonikapanjaitan/belajar-gin-gonic/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
		}
	})

	server.Run(":8080")
}
