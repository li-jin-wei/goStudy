package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func initDB() {
	//初始化数据库

	//连接数据库
	//  第一个参数: 指定要连接的数据库
	// 第二个参数: 指的是数据库的设置信息: 用户名:密码@tcp(ip:端口号)/数据库名字?charset = utf8 & parseTime = True & loc = Local
	// charset = utf8 设置字符集
	// parseTime = True 为了处理time.Time
	// loc = Local 时区设置，与本地时区保持一致
	dsn := "root:1232@@tcp(localhost:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("连接成功", db)
	}
}
