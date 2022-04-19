package payments

import (
	"strconv"

	"github.com/didietsuryadi95/pos/common"
	"github.com/jinzhu/gorm"
)

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&PaymentModel{})
}

type PaymentModel struct {
	gorm.Model
	Name string
	Type string
	Logo string
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOnePayment(condition interface{}) (PaymentModel, error) {
	db := common.GetDB()
	var model PaymentModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindManyPayment(limit, offset string) ([]PaymentModel, int, int, int, error) {
	db := common.GetDB()
	var models []PaymentModel
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

func (model *PaymentModel) Update(id uint, data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Where("id = ?", id).Update(data).Error
	return err
}

func Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(PaymentModel{}).Error
	return err
}
