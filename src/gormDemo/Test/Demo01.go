package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func initDB() {
	dsn := "root:Ljinw1997@@tcp(iljw.top:65501)/Test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	DB = db
}

type User struct {
	gorm.Model
	ID      uint   `json:"ID"     gorm:"autoIncrement,unique"   binding:"required"`
	Name    string `json:"name"   gorm:"unique"                 binding:"required"`
	Age     uint   `json:"age"                                  binding:"required"`
	Sex     string `json:"sex"                                  binding:"required"`
	Address string `json:"address"                              binding:"required"`
	Email   string `json:"email"                                binding:"required"`
}

func main() {
	initDB()
	DB.AutoMigrate(User{})
	r := gin.Default()

	r.POST("/user/add", func(ctx *gin.Context) {
		var user User

		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(400, gin.H{
				"msg":  "添加失败",
				"data": gin.H{},
				"code": 400,
			})
		} else {
			err := DB.Create(&user).Error
			if err != nil {
				ctx.JSON(500, gin.H{
					"msg":  "数据库错误",
					"data": gin.H{},
					"code": 500,
				})
			} else {
				ctx.JSON(200, gin.H{
					"msg":  "添加成功",
					"data": user,
					"code": 200,
				})
			}

		}
	})

	r.Run(":8090")
}
