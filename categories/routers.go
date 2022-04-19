package categories

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/didietsuryadi95/pos/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RouteAuth(router *gin.RouterGroup) {
	router.GET("/:id", GetCategoryDetail)
	router.GET("/", GetCategoryList)
}

func Route(router *gin.RouterGroup) {
	router.POST("", CreateCategory)
	router.PUT("/:id", UpdateCategory)
	router.DELETE("/:id", DeleteCategory)
}

func DeleteCategory(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}
	err = Delete(&CategoryModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}
	common.BaseResponseStatusOnly(c, true, "Success")
}

func UpdateCategory(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	id := uint(id64)
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}

	categoryModel, err := FindOneCategory(&CategoryModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("categories", errors.New("Invalid slug")))
		return
	}
	categoryModelValidator := NewCategoryModelValidatorFillWith(categoryModel)
	if err := categoryModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	if err := categoryModel.Update(categoryModel.ID, categoryModelValidator.categoryModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	common.BaseResponseStatusOnly(c, true, "Success")
}

func CreateCategory(c *gin.Context) {
	categoryModelValidator := NewCategoryModelValidator()
	if err := categoryModelValidator.Bind(c); err != nil {
		common.BaseResponseErrors(c, common.NewValidatorError(err).Errors)
		return
	}

	if err := SaveOne(&categoryModelValidator.categoryModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	serializer := CategoryDetailSerializer{c, categoryModelValidator.categoryModel}
	common.BaseResponse(c, serializer.Response())
}

func GetCategoryList(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("skip")
	cashierModels, modelCount, limitInt, offsetInt, err := FindManyCategory(limit, offset)
	if err != nil {
		common.BaseResponseErrors(c, "Invalid Id")
		return
	}
	serializer := Serializer{c, cashierModels}
	meta := map[string]int{
		"total": modelCount,
		"skip":  offsetInt,
		"limit": limitInt,
	}
	data := map[string]interface{}{"categories": serializer.Response(), "meta": meta}
	common.BaseResponse(c, data)
}

func GetCategoryDetail(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		idInt = 0
	}

	cashierModel, err := FindOneCategory(&CategoryModel{Model: gorm.Model{ID: uint(idInt)}})
	if err != nil || cashierModel.ID == uint(0) {
		common.BaseResponseNotFound(c, "Category Not Found")
		return
	}
	siteSerializer := CategorySerializer{c, cashierModel}
	common.BaseResponse(c, siteSerializer.Response())
}
