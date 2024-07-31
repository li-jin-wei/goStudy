package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type QueryParams struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}
type FormData struct {
	Name string `form:"name"`
	Age  int    `form:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//路由组
	apiGroup := r.Group("/api")
	apiGroup.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "Get Users")
	})
	apiGroup.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, func(c *gin.Context) {
			c.String(http.StatusOK, "Create Users")
		})
	})

	r.GET("/test", func(c *gin.Context) {
		//重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		r.HandleContext(c)
	})

	r.GET("/test3", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.jd.com")
	})

	//使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
	r.GET("/someJson", func(c *gin.Context) {
		data := map[string]string{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/apiuser", func(c *gin.Context) {
		var queryparams QueryParams

		if err := c.ShouldBind(&queryparams); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "参数绑定失败",
			})
		} else {
			c.JSON(200, gin.H{
				"name": queryparams.Name,
				"age":  queryparams.Age,
			})
		}
	})

	r.POST("/user1", func(c *gin.Context) {
		var formData FormData

		//使用ShouldBind与ShouldBindJson方法可以将POST请求的表单数据或JSON数据绑定到结构体中
		err := c.ShouldBind(&formData)
		if err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": "参数绑定失败",
			})
		} else {
			c.JSON(200, gin.H{
				"name": formData.Name,
				"age":  formData.Age,
			})
		}
	})
	r.Run()

}
