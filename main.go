package main

import (
	"github.com/gin-gonic/gin"
	"image-Designer/internal/handler"
	"log"
)

func main() {

	r := gin.Default()
	r.GET("/submit", handler.SubmitHandler)
	r.GET("/result/:id", handler.ResultHandler)
	err := r.Run(":9999")
	if err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
