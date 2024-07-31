package main

import (
	"app1/apps/product_app"
	"app1/apps/user_app"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	u := user_app.UserHandle{}
	u.RegisterRoutes(server)
	p := product_app.ProductHandler{}
	p.RegisterRouters(server)
	server.Run(":8080")

}
