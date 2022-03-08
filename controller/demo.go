package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xb.gin/global"
	"xb.gin/model"
)

func GetStudent(c *gin.Context) {
	id := c.Param("ID")
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		fmt.Println("错误", err)
	}

	global.XB_DB.Preload("Teachers").Preload("IDCard").Where("id = ?", id).First(&student)
	c.JSON(200, gin.H{"data": student})
}

func Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"success": "项目初始化",
	})
}
