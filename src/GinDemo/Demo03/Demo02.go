package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func totalTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	end := time.Now()
	fmt.Println("视图函数运行时间为:", end.Sub(start))
}
func CustomMiddleware(c *gin.Context) {
	c.Set("userID", 1234)
	c.Next()
}
func main() {
	r := gin.Default()
	r.GET("/", totalTime, func(c *gin.Context) {
		time.Sleep(time.Second * 2)
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    time.Now().Unix(),
		})
	})

	r.Run()
}
