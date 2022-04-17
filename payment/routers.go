package payment

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	router.GET("/:id", GetPaymentDetail)
	router.GET("/", GetPaymentList)
	router.POST("", CreatePayment)
	router.PUT("/:id", UpdatePayment)
	router.DELETE("/:id", DeletePayment)
}

func DeletePayment(context *gin.Context) {

}

func UpdatePayment(context *gin.Context) {

}

func CreatePayment(context *gin.Context) {

}

func GetPaymentList(context *gin.Context) {

}

func GetPaymentDetail(context *gin.Context) {

}
