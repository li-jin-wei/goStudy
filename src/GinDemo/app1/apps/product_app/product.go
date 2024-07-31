package product_app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
}

func (p *ProductHandler) RegisterRouters(server *gin.Engine) {
	productGroup := server.Group("/product")
	productGroup.GET("/", getProductList)
	productGroup.GET("/:id", getProductByID)
	productGroup.POST("/", createProduct)

}

func getProductList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get product list",
	})
}

func getProductByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get product by id",
		"id":      "id",
	})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "create product",
	})
}
