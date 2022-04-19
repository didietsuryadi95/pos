package orders

import (
	"strconv"

	"github.com/didietsuryadi95/pos/cashiers"
	"github.com/didietsuryadi95/pos/common"
	"github.com/didietsuryadi95/pos/payments"
	"github.com/didietsuryadi95/pos/products"
	"github.com/jinzhu/gorm"
)

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&OrderModel{})
	db.AutoMigrate(&OrderProductModel{})
}

type OrderModel struct {
	gorm.Model
	CashierId      uint
	Cashier        cashiers.CashierModel
	PaymentTypesId uint
	PaymentType    payments.PaymentModel
	TotalPaid      int32
	TotalReturn    int32
	ReceiptId      string
	Products       []OrderProductModel `gorm:"ForeignKey:OrderId"`
}

type OrderProductModel struct {
	gorm.Model
	OrderId          uint
	ProductId        uint
	Name             cashiers.CashierModel
	DiscountId       uint
	Discount         products.DiscountModel
	Qty              payments.PaymentModel
	Price            int32
	TotalFinalPrice  int32
	TotalNormalPrice int32
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindOneArticle(condition interface{}) (OrderModel, error) {
	db := common.GetDB()
	var model OrderModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindManyArticle(tag, author, limit, offset, favorited string) ([]OrderModel, int, error) {
	db := common.GetDB()
	var models []OrderModel
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

func (model *OrderModel) Update(data interface{}) error {
	db := common.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}

func Delete(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(OrderModel{}).Error
	return err
}
