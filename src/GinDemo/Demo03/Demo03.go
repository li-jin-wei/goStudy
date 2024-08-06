package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//server.Use(cors.New(cors.Config{
	//	AllowMethods:     []string{"PUT", "PATCH"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		if strings.HasPrefix(origin, "http://") || strings.HasPrefix(origin, "https://") {
	//			return true
	//		}
	//		return strings.HasPrefix(origin, "https://")
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://www.google.com"}
	server.Use(cors.New(config))
	server.Run(":80")
}
