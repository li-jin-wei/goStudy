package main

import (
	"DemoApp/web/user"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	u := user.User{}
	u.InitDB()
	u.InitRouter(server)
	server.Run(":80")
}
