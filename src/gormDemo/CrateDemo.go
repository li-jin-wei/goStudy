package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type Product struct {
	gorm.Model
	ID    uint   `gorm:"unique,autoIncrement"`
	Name  string `gorm:"unique"`
	Price float64
}

func InitDB() {
	dsn := "root:123456@tcp(localhost:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("数据库连接成功")
	}
	DB = db
}

func CrateTable() {
	//调用db.AutoMigrate()方法来自动创建或更新表。如果表不存在，它会创建一个新表；如果表已存在，它会根据模型的变化更新表结构。
	err := DB.AutoMigrate(&Product{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("表创建成功")
	}
}

func CrateDemo() {
	//实例化结构体并赋值，创建记录
	//product := Product{Name: "Laptop", Price: 999.99}

	//批量插入使用切片，Gorm会生成一条SQl来插入所有的数据，以返回所有主键值，并触发Hook方法。
	//product := []*Product{
	//	{Name: "Laptop", Price: 99.99},
	//	{Name: "Windows 10", Price: 9999.99},
	//	{Name: "Pencil", Price: 100.99},
	//}
	//通过指针创建数据
	//result := DB.Create(&product)

	product := []Product{
		{Name: "笔记本_1"},
		{Name: "笔记本_2"},
		{Name: "笔记本_3"},
		{Name: "笔记本_4"},
		{Name: "笔记本_5"},
		{Name: "笔记本_6"},
		{Name: "笔记本_7"},
		{Name: "笔记本_8"},
		{Name: "笔记本_9"},
		{Name: "笔记本_10"},
	}

	//数据被分割成多个批次时，GORM会开启一个事务</0>来处理，使用CreateInBatches来分批插入数据
	result := DB.CreateInBatches(&product, 10)
	//result := DB.Create(product)

	//指定字段创建表记录
	//result := DB.Select("Name", "ID").Create(&Product{Name: "Mac笔记本"})

	if result.Error != nil {
		log.Fatal(result.Error)
	} else {
		log.Println("Product created successfully")
	}
}
func main() {
	InitDB()
	CrateTable()
	CrateDemo()
}
