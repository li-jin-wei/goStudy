package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandle struct {
}

func (u *UserHandle) RegisterRoutes(server *gin.Engine) {

	//用户路由组
	userGroup := server.Group("/user")
	userGroup.GET("/", getUserList)
	userGroup.GET("/:id", getUserList)
	userGroup.POST("/", createUser)

}

// 获取用户列表
func getUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get user list"})
}

// 获取用户ID
func getUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Get user by id"})
}

// 修改用户
func createUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Create user"})
}
