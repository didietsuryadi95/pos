package payments

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Serializer struct {
	C             *gin.Context
	PaymentsModel []PaymentModel
}

type PaymentSerializer struct {
	C *gin.Context
	PaymentModel
}

type PaymentDetailSerializer struct {
	C *gin.Context
	PaymentModel
}

type PaymentResponse struct {
	ID   uint   `json:"categoryId"`
	Name string `json:"name"`
}

type PaymentDetailResponse struct {
	ID        uint      `json:"categoryId"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Logo      string    `json:"logo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *Serializer) Response() (payments []PaymentResponse) {
	payments = make([]PaymentResponse, 0)
	for _, c := range s.PaymentsModel {
		payments = append(payments, PaymentResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	return
}

func (s *PaymentSerializer) Response() PaymentResponse {
	return PaymentResponse{
		ID:   s.PaymentModel.ID,
		Name: s.PaymentModel.Name,
	}
}

func (s *PaymentDetailSerializer) Response() PaymentDetailResponse {
	return PaymentDetailResponse{
		ID:        s.PaymentModel.ID,
		Name:      s.PaymentModel.Name,
		Type:      s.PaymentModel.Type,
		CreatedAt: s.PaymentModel.CreatedAt,
		UpdatedAt: s.PaymentModel.UpdatedAt,
	}
}
