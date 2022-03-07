package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//中间件
func Demo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("自定义中间件")
		c.Next()
		fmt.Println("自定义中间件end...")
	}
}
