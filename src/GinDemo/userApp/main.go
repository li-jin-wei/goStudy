package main

import (
	"github.com/gin-gonic/gin"
	"userApp/internal/web"
)

func main() {
	//创建gin的默认引擎
	server := gin.Default()

	//创建一个UserHandle的实例
	u := web.UserHandle{}

	//使用UserHandle 实例注册用户模块的路由
	u.RegisterRoutes(server)

	//在8080端口上启动gin服务
	server.Run()

}
