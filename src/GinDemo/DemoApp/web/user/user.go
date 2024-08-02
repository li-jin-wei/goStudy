package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
)

type User struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"  binding:"required"`
	Sex  string `json:"sex"  binding:"required"`
	City string `json:"city" binding:"required"`
}

// DB 数据库
var DB *gorm.DB

// InitDB 初始化数据库
func (u *User) InitDB() {
	dsn := "root:Ljinw1997@@tcp(iljw.top:65501)/gin_Demo?charset=utf8&parseTime=True&loc"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	DB = db

	//自动迁移并创建表
	DB.AutoMigrate(&User{})
}

// InitRouter 初始化路由
func (u *User) InitRouter(r *gin.Engine) {
	r.POST("/add/user", u.InsertUser)
	r.GET("/get/lists", u.GetUserList)
}

// InsertUser 新增用户
func (u *ser) InsertUser(c *gin.Context) {
	var user []User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"msg":  "添加失败",
			"data": gin.H{},
			"code": 400,
		})
	} else {
		DB.Create(&user)
		c.JSON(200, gin.H{
			"msg":  "添加成功",
			"data": user,
			"code": 200,
		})
	}

}

// GetUserList 获取用户列表
func (u *User) GetUserList(c *gin.Context) {
	var user []User
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	offsetVal := (pageNum - 1) * pageSize

	if pageSize == -1 && pageNum == -1 {
		offsetVal = 0
	}

	var total int64

	//查询总数
	DB.Model(user).Count(&total)

	//分页查询
	DB.Model(user).Limit(pageSize).Offset(offsetVal).Find(&user)

	if len(user) == 0 {
		c.JSON(200, gin.H{
			"msg":  "没有查询到数据",
			"code": 400,
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "查询成功",
			"code": 200,
			"data": gin.H{
				"list":     user,
				"total":    total,
				"pageNum":  pageNum,
				"pageSize": pageSize,
			},
		})
	}

}
