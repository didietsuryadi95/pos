package products

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Serializer struct {
	C             *gin.Context
	ProductsModel []ProductModel
}

type ProductSerializer struct {
	C *gin.Context
	ProductModel
}

type ProductDetailSerializer struct {
	C *gin.Context
	ProductModel
}

type ProductResponse struct {
	ID         uint   `json:"productId"`
	Name       string `json:"name"`
	Sku        string `json:"sku"`
	Image      string `json:"image"`
	Stock      int    `json:"stock"`
	Price      int32  `json:"price"`
	CategoryId uint   `json:"categoryId"`
}

type ProductDetailResponse struct {
	ID         uint      `json:"productId"`
	Name       string    `json:"name"`
	Sku        string    `json:"sku"`
	Image      string    `json:"image"`
	Stock      int       `json:"stock"`
	Price      int32     `json:"price"`
	CategoryId uint      `json:"categoryId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (s *Serializer) Response() (categories []ProductResponse) {
	categories = make([]ProductResponse, 0)
	for _, c := range s.ProductsModel {
		categories = append(categories, ProductResponse{
			ID:    c.ID,
			Name:  c.Name,
			Sku:   c.Sku,
			Image: c.Image,
			Stock: c.Stock,
			Price: c.Price,
		})
	}
	return
}

func (s *ProductSerializer) Response() ProductResponse {
	return ProductResponse{
		ID:         s.ProductModel.ID,
		Name:       s.ProductModel.Name,
		Sku:        s.ProductModel.Sku,
		Image:      s.ProductModel.Image,
		Stock:      s.ProductModel.Stock,
		Price:      s.ProductModel.Price,
		CategoryId: s.ProductModel.CategoryId,
	}
}

func (s *ProductDetailSerializer) Response() ProductDetailResponse {
	return ProductDetailResponse{
		ID:         s.ProductModel.ID,
		Name:       s.ProductModel.Name,
		Sku:        s.ProductModel.Sku,
		Image:      s.ProductModel.Image,
		Stock:      s.ProductModel.Stock,
		Price:      s.ProductModel.Price,
		CategoryId: s.ProductModel.CategoryId,
		CreatedAt:  s.ProductModel.CreatedAt,
		UpdatedAt:  s.ProductModel.UpdatedAt,
	}
}
