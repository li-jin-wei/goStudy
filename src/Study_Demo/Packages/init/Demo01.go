package main

//关键字：init
//init函数可以在所有程序执行开始前被调用，并且每个包下可以有多个init函数
//init函数先于main函数自动执行
//每个包中可以有多个init函数，每个包中的源文件中也可以有多个init函数
//init函数没有输入参数、返回值，也未声明，所以无法引用
//不同包的init函数按照包导入的依赖关系决定执行顺序
//无论包被导入多少次，init函数只会被调用一次，也就是只执行一次
//init函数在代码中不能被显示的调用，不能被引用（赋值给函数变量），否则会出现编译错误
//Go程序仅仅想要用一个package的init执行，我们可以这样使用：import _ “test_xxxx”，导入包的时候加上下划线就ok了
//init函数不应该依赖任何在main函数里创建的变量，因为init函数的执行是在main函数之前的

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func init() {
	dsn := "root:Ljinw1997.@tcp(121.40.224.156:3306)/curd_list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("链接数据库失败:", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
