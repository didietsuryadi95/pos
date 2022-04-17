package order

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"pos/common"
)

type Order struct {
	gorm.Model
	CashierId      int `gorm:"unique_index"`
	PaymentTypesId int
	TotalPaid      int32
	TotalReturn    int32
	ReceiptId      string
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneArticle(condition interface{}) (Order, error) {
	db := common.GetDB()
	var model Order
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindManyArticle(tag, author, limit, offset, favorited string) ([]Order, int, error) {
	db := common.GetDB()
	var models []Order
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

func (model *Order) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}

func Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(Order{}).Error
	return err
}
