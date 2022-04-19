package categories

import (
	"strconv"

	"github.com/didietsuryadi95/pos/common"
	"github.com/jinzhu/gorm"
)

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&CategoryModel{})
}

type CategoryModel struct {
	gorm.Model
	Name string
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneCategory(condition interface{}) (CategoryModel, error) {
	db := common.GetDB()
	var model CategoryModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindManyCategory(limit, offset string) ([]CategoryModel, int, int, int, error) {
	db := common.GetDB()
	var models []CategoryModel
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

func (model *CategoryModel) Update(id uint, data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Where("id = ?", id).Update(data).Error
	return err
}

func Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(CategoryModel{}).Error
	return err
}
