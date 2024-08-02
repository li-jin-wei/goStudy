package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/path", func(c *gin.Context) {
		c.String(http.StatusOK, "这是静态路由")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "传递过来的名字：%s", name)
	})

	r.GET("/view/*.html", func(c *gin.Context) {
		path := c.Param(".html")
		c.String(http.StatusOK, "匹配的值是%s", path)
	})

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.POST("/test1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/foo")
		r.GET("/foo", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "foo",
			})
		})
	})

	r.GET("/test2", func(c *gin.Context) {
		c.Request.URL.Path = "/test3"
		r.HandleContext(c)
	})
	r.GET("/test3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test3",
		})
	})
	r.Run(":80")
}
