package main

import (
	"github.com/gin-gonic/gin"
	"xb.gin/config"
	"xb.gin/global"
)

func main() {
	r := gin.Default()
	global.XB_DB = config.InitDB()
	config.InitRouter(r)

	//defer global.XB_DB.Close()

	panic(r.Run(":7777"))
}
