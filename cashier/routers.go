package cashier

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	router.GET("/:id", GetCashierDetail)
	router.GET("/", GetCahierList)
	router.POST("", CreateCashier)
	router.PUT("/:id", UpdateCashier)
	router.DELETE("/:id", DeleteCashier)
	router.GET("/:id/passcode", GetPassCode)
	router.POST("/:id/login", LoginCashier)
	router.POST("/:id/logout", LogoutCashier)
}

func LogoutCashier(context *gin.Context) {

}

func LoginCashier(context *gin.Context) {

}

func GetPassCode(context *gin.Context) {

}

func DeleteCashier(context *gin.Context) {

}

func UpdateCashier(context *gin.Context) {

}

func CreateCashier(context *gin.Context) {

}

func GetCahierList(context *gin.Context) {

}

func GetCashierDetail(context *gin.Context) {

}
