package order

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	router.GET("/:id", GetOrderDetail)
	router.GET("/", GetOrderList)
	router.POST("", AddOrder)
	router.POST("/subtotal", SubTotalOrder)
	router.GET("/:id/download", DownloadInvoice)
	router.GET("/:id/check-download", CheckDownloadStatus)
}

func Report(router *gin.RouterGroup) {
	router.GET("/revenues", GetRevenues)
	router.GET("/solds", GetSolds)
}

func GetSolds(context *gin.Context) {

}

func GetRevenues(context *gin.Context) {

}

func SubTotalOrder(context *gin.Context) {

}

func CheckDownloadStatus(context *gin.Context) {

}

func DownloadInvoice(context *gin.Context) {

}

func AddOrder(context *gin.Context) {

}

func GetOrderList(context *gin.Context) {

}

func GetOrderDetail(context *gin.Context) {

}
