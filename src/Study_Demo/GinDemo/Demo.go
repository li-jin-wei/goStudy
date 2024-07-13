package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
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

	type List struct {
		gorm.Model        //主键
		Name       string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
		State      string `gorm:"type:varchar(20); not null" json:"state" binding:"required"`
		Phone      string `gorm:"type:varchar(20); not null" json:"phone" binding:"required"`
		Email      string `gorm:"type:varchar(40); not null" json:"email" binding:"required"`
		Address    string `gorm:"type:varchar(200); not null" json:"address" binding:"required"`
	}
	err = db.AutoMigrate(&List{})
	if err != nil {
		return
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Gin+Gorm测试crud接口!")
	})

	r.POST("/user/add", func(ctx *gin.Context) {
		var data List

		err := ctx.ShouldBindJSON(&data)
		if err != nil {
			ctx.JSON(200, gin.H{
				"msg":  "添加失败",
				"data": gin.H{},
				"code": 400,
			})
		} else {
			db.Create(&data)
			ctx.JSON(200, gin.H{
				"msg":  "添加成功",
				"data": data,
				"code": 200,
			})
		}
	})
	r.DELETE("/user/delete/:id", func(ctx *gin.Context) {
		var data []List
		id := ctx.Param("id")
		db.Where("id =?", id).Find(&data)
		if len(data) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "id没有找到，删除失败",
				"code": 400,
			})
		} else {
			db.Delete(&data)
			ctx.JSON(200, gin.H{
				"msg":  "删除成功",
				"code": 200,
			})
		}
	})

	r.PUT("/user/update/:id", func(ctx *gin.Context) {
		var data List
		id := ctx.Param("id")
		db.Select("id").Where("id =?", id).Find(&data)
		if data.ID == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "用户id没有找到",
				"code": 400,
			})
		} else {
			err := ctx.ShouldBindJSON(&data)
			if err != nil {
				ctx.JSON(200, gin.H{
					"msg":  "修改失败",
					"code": 400,
				})
			} else {
				db.Where("id =?", id).Updates(&data)
				ctx.JSON(200, gin.H{
					"msg":  "修改成功",
					"code": 200,
				})
			}
		}
	})

	r.GET("/User/list/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		var dataList []List
		db.Where("name =?", name).Find(&dataList)
		if len(dataList) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "没有查询到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			ctx.JSON(200, gin.H{
				"msg":  "查询成功",
				"code": 200,
				"data": dataList,
			})
		}

	})

	//r.GET("/User/list/:id", func(ctx *gin.Context) {
	//	id := ctx.Param("id")
	//	var dataList []List
	//	db.Select("id").Where("id =?", id).Find(&dataList)
	//	if dataList.ID == 0 {
	//		ctx.JSON(200, gin.H{
	//			"msg":  "用户id没有找到",
	//			"code": 400,
	//		})
	//	} else {
	//		ctx.JSON(200, gin.H{
	//			"msg":  "查询成功",
	//			"code": 200,
	//			"data": dataList,
	//		})
	//	}
	//
	//})

	r.GET("/user/list", func(ctx *gin.Context) {
		var dataList []List

		pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
		pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = -1
		}
		offsetVal := (pageNum - 1) * pageSize
		if pageNum == -1 && pageSize == -1 {
			offsetVal = -1
		}

		var total int64
		db.Model(dataList).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&dataList)

		if len(dataList) == 0 {
			ctx.JSON(200, gin.H{
				"msg":  "没有查询到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			ctx.JSON(200, gin.H{
				"msg":  "查询成功",
				"code": 200,
				"data": gin.H{
					"list":     dataList,
					"total":    total,
					"pageNum":  pageNum,
					"pageSize": pageSize,
				},
			})
		}
	})
	PORT := "3001"
	err = r.Run(":" + PORT)
	if err != nil {
		return
	}
}
