package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"xb.gin/global"
	"xb.gin/model"
)

func Register(ctx *gin.Context) {
	name := ctx.PostForm("name")
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	if len(name) == 0 {
		ctx.JSON(400, gin.H{"msg": "名称不能为空"})
		return
	}

	if len(phone) != 11 {
		ctx.JSON(400, gin.H{"msg": "手机号必须为11位"})
		return
	}

	if isPhoneExit(phone, global.XB_DB) {
		ctx.JSON(400, gin.H{"msg": "手机号已存在"})
		return
	}

	newUser := model.User{
		Name:     name,
		Phone:    phone,
		Password: password,
	}
	global.XB_DB.Create(&newUser)
	ctx.JSON(200, gin.H{"msg": "注册成功"})
}

func isPhoneExit(phone string, db *gorm.DB) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
