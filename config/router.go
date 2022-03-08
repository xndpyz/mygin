package router

import (
	"github.com/gin-gonic/gin"
	"xb.gin/api"
	"xb.gin/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Demo())
	r.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "项目初始化",
		})
	})

	r.GET("/demo/student/:ID", api.GetStudent)

	return r
}
