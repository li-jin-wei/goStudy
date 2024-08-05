package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	//但文件上传并把内容返回给前端
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.SaveUploadedFile(file, "./"+file.Filename)
		c.Writer.Header().Add("Content-Disposition", "attachment; filename="+file.Filename)
		c.JSON(http.StatusOK, gin.H{
			"message": file.Filename,
			"data":    file,
		})
	})

	//多文件上传
	r.POST("/uploadFiles", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"message": err.Error(),
		//	})
		//}

		files := form.File["files"]
		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "./"+file.Filename)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data":    len(files),
		})
	})
	r.Run(":8081")
}
