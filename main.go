package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"xb.gin/config"
	"xb.gin/middleware"
)

func main() {
	r := gin.Default()

	XB_DB := config.InitDB()
	defer XB_DB.Close()

	r.Use(middleware.Demo())
	r.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "项目初始化",
		})
	})

	r.GET("/demo/student/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		var student Student
		_ = c.BindJSON(&student)

		XB_DB.Preload("Teachers").Preload("IDCard").Where("id = ?", id).First(&student)
		c.JSON(200, gin.H{"查询": student})
	})

	panic(r.Run(":7777"))
}

type Class struct {
	gorm.Model
	ClassName int
	Students  []Student
}

type Student struct {
	gorm.Model
	StudentName string
	ClassId     uint
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teachers;"`
}

type IDCard struct {
	gorm.Model
	Num       int
	StudentId uint
}

type Teacher struct {
	gorm.Model
	TeacherName string
	Students    []Student `gorm:"many2many:student_teachers;"`
}
