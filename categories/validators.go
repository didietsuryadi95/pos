package categories

import (
	"github.com/didietsuryadi95/pos/common"
	"github.com/gin-gonic/gin"
)

type CategoryModelValidator struct {
	Name          string        `form:"name" json:"name" binding:"required"`
	categoryModel CategoryModel `json:"-"`
}

func NewCategoryModelValidator() CategoryModelValidator {
	return CategoryModelValidator{}
}

func NewCategoryModelValidatorFillWith(categoryModel CategoryModel) CategoryModelValidator {
	categoryModelValidator := NewCategoryModelValidator()
	categoryModelValidator.Name = categoryModel.Name
	return categoryModelValidator
}

func (s *CategoryModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.categoryModel.Name = s.Name
	return nil
}
