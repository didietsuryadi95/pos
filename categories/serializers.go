package categories

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Serializer struct {
	C               *gin.Context
	CategoriesModel []CategoryModel
}

type CategorySerializer struct {
	C *gin.Context
	CategoryModel
}

type CategoryDetailSerializer struct {
	C *gin.Context
	CategoryModel
}

type CategoryResponse struct {
	ID   uint   `json:"categoryId"`
	Name string `json:"name"`
}

type CategoryDetailResponse struct {
	ID        uint      `json:"categoryId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *Serializer) Response() (categories []CategoryResponse) {
	categories = make([]CategoryResponse, 0)
	for _, c := range s.CategoriesModel {
		categories = append(categories, CategoryResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	return
}

func (s *CategorySerializer) Response() CategoryResponse {
	return CategoryResponse{
		ID:   s.CategoryModel.ID,
		Name: s.CategoryModel.Name,
	}
}

func (s *CategoryDetailSerializer) Response() CategoryDetailResponse {
	return CategoryDetailResponse{
		ID:        s.CategoryModel.ID,
		Name:      s.CategoryModel.Name,
		CreatedAt: s.CategoryModel.CreatedAt,
		UpdatedAt: s.CategoryModel.UpdatedAt,
	}
}
