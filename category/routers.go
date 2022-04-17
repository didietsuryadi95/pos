package category

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	router.GET("/:id", GetCategoryDetail)
	router.GET("/", GetCategoryList)
	router.POST("", CreateCategory)
	router.PUT("/:id", UpdateCategory)
	router.DELETE("/:id", DeleteCategory)
}

func DeleteCategory(context *gin.Context) {

}

func UpdateCategory(context *gin.Context) {

}

func CreateCategory(context *gin.Context) {

}

func GetCategoryList(context *gin.Context) {

}

func GetCategoryDetail(context *gin.Context) {

}
