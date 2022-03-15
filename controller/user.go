package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
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

	if len(password) < 6 {
		ctx.JSON(400, gin.H{"mag": "密码不能少于6位"})
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
	//hasedPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {
	//	ctx.JSON(http.StatusServiceUnavailable, gin.H{"msg": "密码加密错误"})
	//	return
	//}
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

func Login(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	var user model.User
	global.XB_DB.Where("phone = ?", phone).First(&user)
	if user.ID == 0 {
		ctx.JSON(400, gin.H{"msg": "用户不存在"})
		return
	}

	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "密码错误"})
	//}
	if user.Password != password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"msg": "密码错误"})
		return
	}

	token := "tmp"

	ctx.JSON(200, gin.H{
		"msg":  "登陆成功",
		"data": token,
	})
}
