package config

import (
	"github.com/gin-gonic/gin"
	"xb.gin/controller"
)

func InitRouter(r *gin.Engine) {
	//r.Use(middleware.Demo())
	r.GET("/index", controller.Index)
	r.GET("/student/:ID", controller.GetStudent)

	r.POST("/auth/register", controller.Register)
}
