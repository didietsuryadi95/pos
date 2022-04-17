package category

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"pos/common"
)

type Category struct {
	gorm.Model
	Name     string
	Products []Product `gorm:"many2many:article_tags;"`
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneArticle(condition interface{}) (Category, error) {
	db := common.GetDB()
	var model Category
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindManyArticle(tag, author, limit, offset, favorited string) ([]Category, int, error) {
	db := common.GetDB()
	var models []Category
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

func (model *Category) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}

func Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(Category{}).Error
	return err
}
