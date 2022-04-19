package products

import (
	"strconv"
	"time"

	"github.com/didietsuryadi95/pos/categories"
	"github.com/didietsuryadi95/pos/common"
	"github.com/jinzhu/gorm"
)

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&DiscountModel{})
	db.AutoMigrate(&ProductModel{})
}

type DiscountModel struct {
	gorm.Model
	qty       int
	discType  string
	result    int
	expiredAt time.Time
}

type ProductModel struct {
	gorm.Model
	Sku        string
	Name       string
	Stock      int
	Price      int32
	Image      string
	Category   categories.CategoryModel
	CategoryId uint
	Discount   DiscountModel
	DiscountId uint
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneProduct(condition interface{}) (ProductModel, error) {
	db := common.GetDB()
	var model ProductModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindOneDiscount(condition interface{}) (DiscountModel, error) {
	db := common.GetDB()
	var model DiscountModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindManyProduct(limit, offset string) ([]ProductModel, int, int, int, error) {
	db := common.GetDB()
	var models []ProductModel
	var count int

	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 20
	}

	tx := db.Begin()
	db.Model(&models).Count(&count)
	db.Offset(offset_int).Limit(limit_int).Find(&models)
	err = tx.Commit().Error
	return models, count, limit_int, offset_int, err
}

func (model *ProductModel) Update(id uint, data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Where("id = ?", id).Update(data).Error
	return err
}

func Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(ProductModel{}).Error
	return err
}
