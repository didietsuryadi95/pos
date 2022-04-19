package payments

import (
	"net/http"
	"strconv"

	"github.com/didietsuryadi95/pos/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RouteAuth(router *gin.RouterGroup) {
	router.GET("/:id", GetPaymentDetail)
	router.GET("/", GetPaymentList)
}

func Route(router *gin.RouterGroup) {
	router.POST("", CreatePayment)
	router.PUT("/:id", UpdatePayment)
	router.DELETE("/:id", DeletePayment)
}

func DeletePayment(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}
	err = Delete(&PaymentModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}
	common.BaseResponseStatusOnly(c, true, "Success")
}

func UpdatePayment(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}

	paymentModel, err := FindOnePayment(&PaymentModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}
	paymentModelValidator := NewPaymentModelValidatorFillWith(paymentModel)
	if err := paymentModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	if err := paymentModel.Update(paymentModel.ID, paymentModelValidator.paymentModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	common.BaseResponseStatusOnly(c, true, "Success")
}

func CreatePayment(c *gin.Context) {
	paymentModelValidator := NewPaymentModelValidator()
	if err := paymentModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	if err := SaveOne(&paymentModelValidator.paymentModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := PaymentDetailSerializer{c, paymentModelValidator.paymentModel}
	common.BaseResponse(c, serializer.Response())
}

func GetPaymentList(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("skip")
	cashierModels, modelCount, limitInt, offsetInt, err := FindManyPayment(limit, offset)
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Param")
		return
	}
	serializer := Serializer{c, cashierModels}
	meta := map[string]int{
		"total": modelCount,
		"skip":  offsetInt,
		"limit": limitInt,
	}
	data := map[string]interface{}{"payments": serializer.Response(), "meta": meta}
	common.BaseResponse(c, data)
}

func GetPaymentDetail(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		idInt = 0
	}

	paymentModel, err := FindOnePayment(&PaymentModel{Model: gorm.Model{ID: uint(idInt)}})
	if err != nil || paymentModel.ID == uint(0) {
		common.BaseResponseNotFound(c, "Payment Not Found")
		return
	}
	siteSerializer := PaymentSerializer{c, paymentModel}
	common.BaseResponse(c, siteSerializer.Response())
}
