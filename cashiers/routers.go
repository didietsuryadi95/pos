package cashiers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/didietsuryadi95/pos/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Route(router *gin.RouterGroup) {
	router.GET("/:id", GetCashierDetail)
	router.GET("/", GetCashierList)
	router.POST("", CreateCashier)
	router.PUT("/:id", UpdateCashier)
	router.DELETE("/:id", DeleteCashier)
	router.GET("/:id/passcode", GetPassCode)
	router.POST("/:id/login", LoginCashier)
	router.POST("/:id/logout", LogoutCashier)
}

func LogoutCashier(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		idInt = 0
	}

	authModelValidator := AuthModelValidator()
	if err := authModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	cashierModel, err := FindOneCashier(&CashierModel{Model: gorm.Model{ID: uint(idInt)}, Passcode: authModelValidator.Passcode})

	if err != nil || cashierModel.ID == 0 {
		common.BaseResponseUnautorized(c, "Passcode Not Match")
		return
	}
	common.BaseResponseStatusOnly(c, true, "Success")

}

func LoginCashier(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		idInt = 0
	}

	authModelValidator := AuthModelValidator()
	if err := authModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	cashierModel, err := FindOneCashier(&CashierModel{Model: gorm.Model{ID: uint(idInt)}, Passcode: authModelValidator.Passcode})

	if err != nil || cashierModel.ID == 0 {
		common.BaseResponseUnautorized(c, "Passcode Not Match")
		return
	}
	UpdateContextUserModel(c, cashierModel.ID)
	common.BaseResponse(c, map[string]string{"token": common.GenToken(cashierModel.ID)})
}

func GetPassCode(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		idInt = 0
	}

	cashierModel, err := FindOneCashier(&CashierModel{Model: gorm.Model{ID: uint(idInt)}})
	if err != nil || cashierModel.ID == 0 {
		common.BaseResponseNotFound(c, "Cashier Not Found")
		return
	}

	common.BaseResponse(c, map[string]string{"passcode": cashierModel.Passcode})
}

func DeleteCashier(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		common.BaseResponseNotFound(c, "Cashier Not Found")
		return
	}
	err = Delete(&CashierModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		common.BaseResponseNotFound(c, "Cashier Not Found")
		return
	}
	common.BaseResponseStatusOnly(c, true, "Success")
}

func UpdateCashier(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		common.BaseResponseNotFound(c, "Cashier Not Found")
		return
	}

	cashierModel, err := FindOneCashier(&CashierModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		common.BaseResponseNotFound(c, "Cashier Not Found")
		return
	}
	cashierModelValidator := NewCashierModelValidatorFillWith(cashierModel)
	if err := cashierModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	if err := cashierModel.Update(cashierModel.ID, cashierModelValidator.cashierModel); err != nil {
		common.BaseResponseNotFound(c, "Cashier Not Found")
		return
	}

	common.BaseResponseStatusOnly(c, true, "Success")
}

func CreateCashier(c *gin.Context) {
	cashierModelValidator := NewCashierModelValidator()
	if err := cashierModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	if err := SaveOne(&cashierModelValidator.cashierModel); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("database", err))
		return
	}
	serializer := CashierDetailSerializer{c, cashierModelValidator.cashierModel}
	common.BaseResponse(c, serializer.Response())
}

func GetCashierList(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("skip")
	cashierModels, modelCount, limitInt, offsetInt, err := FindManyCashier(limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("cashiers", errors.New("Invalid param")))
		return
	}
	serializer := Serializer{c, cashierModels}
	meta := map[string]int{
		"total": modelCount,
		"skip":  offsetInt,
		"limit": limitInt,
	}
	data := map[string]interface{}{"cashiers": serializer.Response(), "meta": meta}
	common.BaseResponse(c, data)
}

func GetCashierDetail(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		idInt = 0
	}

	cashierModel, err := FindOneCashier(&CashierModel{Model: gorm.Model{ID: uint(idInt)}})
	if err != nil || cashierModel.ID == uint(0) {
		common.BaseResponseNotFound(c, "Cashier Not Found")
		return
	}
	siteSerializer := CashierSerializer{c, cashierModel}
	common.BaseResponse(c, siteSerializer.Response())
}
