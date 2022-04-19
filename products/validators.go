package products

import (
	"time"

	"github.com/didietsuryadi95/pos/categories"
	"github.com/didietsuryadi95/pos/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Discount struct {
	qty       int       `form:"qty" json:"qty" binding:"exists"`
	discType  string    `form:"type" json:"type" binding:"exists"`
	result    int       `form:"result" json:"result" binding:"exists"`
	expiredAt time.Time `form:"expiredAt" json:"expiredAt" binding:"exists" time_format:"unix"`
}

type ProductModelValidator struct {
	Name         string       `form:"name" json:"name" binding:"required"`
	Sku          string       `form:"sku" json:"sku"`
	Stock        int          `form:"stock" json:"stock" binding:"required"`
	Price        int32        `form:"price" json:"price" binding:"required"`
	Image        string       `form:"image" json:"image" binding:"required"`
	CategoryId   int          `form:"categoryId" json:"categoryId" binding:"required"`
	DiscountId   int          `form:"discountId" json:"discountId"`
	Discount     Discount     `json:"discount"`
	productModel ProductModel `json:"-"`
}

func NewProductModelValidator() ProductModelValidator {
	return ProductModelValidator{}
}

func NewProductModelValidatorFillWith(productModel ProductModel) ProductModelValidator {
	productModelValidator := NewProductModelValidator()
	productModelValidator.Sku = productModel.Sku
	productModelValidator.DiscountId = int(productModel.DiscountId)
	if (DiscountModel{}) != productModel.Discount {
		productModelValidator.Discount.qty = productModel.Discount.qty
		productModelValidator.Discount.discType = productModel.Discount.discType
		productModelValidator.Discount.result = productModel.Discount.result
		productModelValidator.Discount.expiredAt = productModel.Discount.expiredAt
	}

	return productModelValidator
}

func (model *ProductModel) setDiscount(id uint, product ProductModelValidator) error {
	db := common.GetDB()
	if 0 != id {
		discount, err := FindOneDiscount(&DiscountModel{Model: gorm.Model{ID: id}})
		if err != nil {
			return nil
		}

		model.Discount = discount
	} else if (Discount{}) != product.Discount {
		var discModel DiscountModel
		err := db.FirstOrCreate(DiscountModel{
			qty:       product.Discount.qty,
			discType:  product.Discount.discType,
			result:    product.Discount.result,
			expiredAt: product.Discount.expiredAt,
		}, DiscountModel{Model: gorm.Model{ID: id}}).Model(&discModel).Error

		if nil != err {
			return nil
		}
		model.Discount = discModel
	}
	return nil
}
func (s *ProductModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.productModel.Name = s.Name
	s.productModel.Sku = s.Sku
	s.productModel.Stock = s.Stock
	s.productModel.Price = s.Price
	s.productModel.Image = s.Image
	s.productModel.DiscountId = uint(s.DiscountId)
	s.productModel.setDiscount(uint(s.DiscountId), *s)

	categoryModel, err := categories.FindOneCategory(&categories.CategoryModel{Model: gorm.Model{ID: uint(s.CategoryId)}})
	if err != nil {
		return nil
	}
	s.productModel.CategoryId = categoryModel.ID
	s.productModel.Category = categoryModel
	return nil
}
