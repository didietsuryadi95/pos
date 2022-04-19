package products

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/didietsuryadi95/pos/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RouteAuth(router *gin.RouterGroup) {
	router.GET("/:id", GetProductDetail)
	router.GET("/", GetProductList)
}

func Route(router *gin.RouterGroup) {
	router.POST("", CreateProduct)
	router.PUT("/:id", UpdateProduct)
	router.DELETE("/:id", DeleteProduct)
}

func DeleteProduct(c *gin.Context) {

	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		common.BaseResponseNotFound(c, "Product Not Found")
		return
	}
	err = Delete(&ProductModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		common.BaseResponseNotFound(c, "Product Not Found")
		return
	}
	common.BaseResponseStatusOnly(c, true, "Success")
}

func UpdateProduct(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		common.BaseResponseNotFound(c, "Product Not Found")
		return
	}

	productModel, err := FindOneProduct(&ProductModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		common.BaseResponseNotFound(c, "Product Not Found")
		return
	}
	productModelValidator := NewProductModelValidatorFillWith(productModel)
	if err := productModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	if err := productModel.Update(productModel.ID, productModelValidator.productModel); err != nil {
		common.BaseResponseNotFound(c, "Product Not Found")
		return
	}

	common.BaseResponseStatusOnly(c, true, "Success")
}

func CreateProduct(c *gin.Context) {
	productModelValidator := NewProductModelValidator()
	if err := productModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}
	err := SaveOne(&productModelValidator.productModel)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	productModelValidator.productModel.Sku = fmt.Sprintf("ID%03d", productModelValidator.productModel.ID)

	if err := productModelValidator.productModel.Update(productModelValidator.productModel.ID, productModelValidator.productModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := ProductDetailSerializer{c, productModelValidator.productModel}
	common.BaseResponse(c, serializer.Response())
}

func GetProductList(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("skip")
	keyword := c.Query("q")
	categoryId := c.Query("categoryId")
	productModels, modelCount, limitInt, offsetInt, err := FindManyProduct(limit, offset, keyword, categoryId)
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}
	serializer := Serializer{c, productModels}
	meta := map[string]int{
		"total": modelCount,
		"skip":  offsetInt,
		"limit": limitInt,
	}
	data := map[string]interface{}{"categories": serializer.Response(), "meta": meta}
	common.BaseResponse(c, data)
}

func GetProductDetail(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		idInt = 0
	}

	cashierModel, err := FindOneProduct(&ProductModel{Model: gorm.Model{ID: uint(idInt)}})
	if err != nil || cashierModel.ID == uint(0) {
		common.BaseResponseNotFound(c, "Product Not Found")
		return
	}
	siteSerializer := ProductSerializer{c, cashierModel}
	common.BaseResponse(c, siteSerializer.Response())
}
