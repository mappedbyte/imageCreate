package handler

import (
	"github.com/gin-gonic/gin"
	"image-Designer/internal/service"
	"net/http"
)

func SubmitHandler(c *gin.Context) {
	requestMsg := c.Query("message")
	id, err := service.Submit(requestMsg)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"id":      "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求已提交",
		"id":      id,
	})
}

func ResultHandler(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "id不可以为空,操作失败",
			"data":    make([]string, 0),
		})
		return
	}
	result, err := service.Result(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"data":    make([]string, 0),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "操作成功",
		"data":    result,
	})
}
