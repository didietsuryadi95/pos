package product

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"pos/common"
)

type Product struct {
	gorm.Model
	Sku      string
	Name     string
	Stock    int
	Price    int32
	Image    string
	Category int
	Discount int
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneArticle(condition interface{}) (Product, error) {
	db := common.GetDB()
	var model Product
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindManyArticle(tag, author, limit, offset, favorited string) ([]Product, int, error) {
	db := common.GetDB()
	var models []Product
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
	return models, count, err
}

func (model *Product) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}

func Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(Product{}).Error
	return err
}