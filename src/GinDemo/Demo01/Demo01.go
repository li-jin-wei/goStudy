package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//DefaultQuery() 如果key没有制定value，则会返回默认值
	//Query() 若不存在参数，直接返回“”
	r.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "李华")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "OK",
			"username": username,
			"address":  address,
		})
	})

	//Param()获取路径上的参数，并返回
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "OK",
			"username": username,
			"address":  address,
		})
	})

	r.GET("/json", func(c *gin.Context) {
		//从c.Request.Body读取请求数据
		b, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
		}
		fmt.Printf("raw data:%s\n", string(b))
		//定义map或结构体
		var m map[string]interface{}
		//反序列化
		_ = json.Unmarshal(b, &m)
		c.JSON(http.StatusOK, m)
	})
	r.Run()
}
