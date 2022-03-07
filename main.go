package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "项目初始化",
		})
	})

	panic(r.Run(":7777"))
}
