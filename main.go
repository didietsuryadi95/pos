package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/didietsuryadi95/pos/cashiers"
	"github.com/didietsuryadi95/pos/categories"
	"github.com/didietsuryadi95/pos/common"
	"github.com/didietsuryadi95/pos/orders"
	"github.com/didietsuryadi95/pos/payments"
	"github.com/didietsuryadi95/pos/products"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	cashiers.AutoMigrate()
	categories.AutoMigrate()
	payments.AutoMigrate()
	products.AutoMigrate()
	orders.AutoMigrate()
}

func main() {
	// docker run -e MYSQL_HOST=172.17.0.1 -e MYSQL_USER=xxxx -e MYSQL_PASSWORD=xxxxx -e MYSQL_DBNAME=xxxxx -p 8090:3030 docker_username/name_of_app
	db := common.Init(common.DBConfiguration{
		User:     getEnv("MYSQL_USER", "pos"),
		Password: getEnv("MYSQL_PASSWORD", "123456"),
		Host:     getEnv("MYSQL_USER", "127.0.0.1"),
		Port:     getInt("MYSQL_PORT", 3306),
		Name:     getEnv("MYSQL_DBNAME", "pos"),
	})

	Migrate(db)
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1 := r.Group("")
	v1.Use(cashiers.AuthMiddleware(false))
	cashiers.Route(v1.Group("/cashiers"))
	categories.Route(v1.Group("/categories"))
	orders.Route(v1.Group("/orders"))
	payments.Route(v1.Group("/payments"))
	products.Route(v1.Group("/products"))
	orders.Report(v1.Group("/"))

	v1.Use(cashiers.AuthMiddleware(true))
	payments.RouteAuth(v1.Group("/payments"))
	categories.RouteAuth(v1.Group("/categories"))
	products.RouteAuth(v1.Group("/products"))

	r.Run(fmt.Sprintf("127.0.0.1:3030"))
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getInt(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	resultInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return resultInt
}
