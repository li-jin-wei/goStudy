package main

import (
	"gorm.io/gorm"
)

//type User struct {
//	gorm.Model // 内置模型结构体，包含 ID、CreatedAt、UpdatedAt、DeletedAt 字段,用于记录记录的主键、创建时间、更新时间和软删除状态
//	Name       string
//	Age        int
//	Email      string `gorm:"unique"` //使用标签指定字段属性，表示Email字段哎数据库中是唯一的
//	Adress     string
//}

/*
gorm:"column:column_name"：指定字段在数据库中的列名。
gorm:"primaryKey"：指定字段为主键。
gorm:"autoIncrement"：指定字段为自增长。
gorm:"unique"：指定字段在数据库中唯一。
gorm:"not null"：指定字段不能为空。
gorm:"default:value"：指定字段的默认值。
gorm:"size:length"：指定字段的长度。
gorm:"index"：指定字段创建索引
*/

type Product1 struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"size:255,not null;"`
	Price    float64
	Category string `gorm:"index"`
}

//模型关联关系

type Order struct {
	ID          uint
	OrderNumber string
	TotalAmount float64
	UserID      uint
	User        user `gorm:"foreignKey:UserID"`
}

type User struct {
	ID    uint
	Name  string
	Email string
	Order Order
}
