package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"pos/cashier"
	"pos/category"
	"pos/order"
	"pos/payment"
	"pos/product"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
}

func main() {

	r := gin.Default()
	v1 := r.Group("")
	cashier.Route(v1.Group("/cashier"))
	category.Route(v1.Group("/category"))
	order.Route(v1.Group("/order"))
	payment.Route(v1.Group("/payment"))
	product.Route(v1.Group("/product"))
	order.Report(v1.Group("/"))

	r.Run(fmt.Sprintf("127.0.0.1:3030"))
}
