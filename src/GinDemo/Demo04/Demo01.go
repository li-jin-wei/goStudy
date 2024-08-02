package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
	"log"
	"os"
)

type User struct {
	//binding:"required" 这个字段不允许为空值，如果为会报错
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,mustBig"`
	Sex  string `json:"sex" binding:"required"`
}

// age字段自定义验证规则绑定到binding中
// 如果年龄小于等于18，报错
func mustBig(fl validator.FieldLevel) bool {

	//断言接口是int类型
	age, _ := fl.Field().Interface().(int)
	if age > 18 {
		return true
	}
	return false
}

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("测试前")
		c.Next()
		fmt.Println("测试后")
	}
}

func main() {
	r := gin.Default()
	v1 := r.Group("/v1").Use(Test())
	v1.GET("path/:id", func(c *gin.Context) {
		{
			id := c.Param("id")
			//DefaultQuery() 如果没有没有传值，则会返回设定的默认值
			user := c.DefaultQuery("user", "李华")
			pwd := c.Query("pwd")
			c.JSON(200, gin.H{
				"id":   id,
				"user": user,
				"pwd":  pwd,
			})
			fmt.Println("测试中间")
		}
	})

	// // 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	//gin.DisableConsoleColor()
	//强化日志颜色
	gin.ForceConsoleColor()

	//创建日志文件
	f, _ := os.Create("gin.log")

	//同时将日志输出到文件和和控制台
	gin.DefaultWriter = io.MultiWriter(f)

	r.POST("/path/post", func(c *gin.Context) {

		//注册验证规则
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("mustBig", mustBig)
		}

		var user User
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": err.Error(),
			})
			log.Println(err)
		} else {
			c.JSON(200, gin.H{
				"code":    200,
				"message": user,
			})
		}

	})

	//单文件上传
	r.POST("/testUpload", func(c *gin.Context) {
		//设置上传文件大小
		r.MaxMultipartMemory = 8 << 20 // 8 GB

		//form-data 获取值
		fileName := c.PostForm("fileName")

		//获取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "./"+file.Filename)

		//文件内容返回给前端
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file.Filename))
		c.File("./" + file.Filename)
		c.JSON(200, gin.H{
			"fileName": fileName,
			"code":     200,
			"message":  file,
		})
	})

	//多文件上传
	r.POST("/testUploadFiles", func(c *gin.Context) {

		//r.MaxMultipartMemory = 8 << 20

		form, _ := c.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "./"+file.Filename)
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, file.Filename))
			c.File("./" + file.Filename)
		}
		c.JSON(200, gin.H{
			"code":    200,
			"message": len(files),
			"files":   files,
		})
	})

	r.DELETE("/path/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "true",
			"message": "已删除",
		})
	})

	r.PUT("/path/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "data",
		})
	})
	r.Run(":8081")
}
