package main

import (
	"github.com/gin-gonic/gin"
	"webook/internal/web"
)

func main() {

	server := gin.Default()
	u := web.NewUserHandle()
	u.RegisterRoutes(server)
	server.Run("127.0.0.1:3000")
}
