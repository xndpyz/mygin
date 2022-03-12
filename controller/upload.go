package controller

import "github.com/gin-gonic/gin"

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, "./"+file.Filename)
	c.JSON(200, gin.H{
		"file": file,
	})
}
