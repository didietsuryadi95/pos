package product

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	router.GET("/:id", GetProductDetail)
	router.GET("/", GetProductList)
	router.POST("", CreateProduct)
	router.PUT("/:id", UpdateProduct)
	router.DELETE("/:id", DeleteProduct)
}

func DeleteProduct(context *gin.Context) {

}

func UpdateProduct(context *gin.Context) {

}

func CreateProduct(context *gin.Context) {

}

func GetProductList(context *gin.Context) {

}

func GetProductDetail(context *gin.Context) {

}
