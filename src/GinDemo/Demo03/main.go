package main

import (
	//"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginServer := gin.Default()

	//ginServer.Use(favicon.New("./Users/dayu/go_project/favicon.ico"))

	//ginServer.LoadHTMLGlob("templates/*")
	//ginServer.Static("/static", "./static")

	ginServer.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"msg": "This data is from Go background",
		})
	})

	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")

		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	ginServer.POST("/json", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(data, &m)
		ctx.JSON(http.StatusOK, m)
	})

	ginServer.POST("/user/add", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})

	ginServer.GET("/test", func(ctx *gin.Context) {
		ctx.Redirect(301, "https://www.baidu.com")
	})

	ginServer.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(404, "404.html", nil)
	})
	err := ginServer.Run(":8090")
	if err != nil {
		return
	}
}
