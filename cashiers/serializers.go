package cashiers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CashierDetailSerializer struct {
	C *gin.Context
	CashierModel
}

type Serializer struct {
	C             *gin.Context
	CashiersModel []CashierModel
}

type CashierSerializer struct {
	C *gin.Context
	CashierModel
}

type CashierDetailResponse struct {
	ID        uint      `json:"cashierId"`
	Passcode  string    `json:"passcode"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CashierResponse struct {
	ID   uint   `json:"cashierId"`
	Name string `json:"name"`
}

func (s *CashierDetailSerializer) Response() CashierDetailResponse {
	return CashierDetailResponse{
		ID:        s.CashierModel.ID,
		Passcode:  s.CashierModel.Passcode,
		Name:      s.CashierModel.Name,
		CreatedAt: s.CashierModel.CreatedAt,
		UpdatedAt: s.CashierModel.CreatedAt,
	}
}

func (s *Serializer) Response() (cashiers []CashierResponse) {
	cashiers = make([]CashierResponse, 0)
	for _, c := range s.CashiersModel {
		cashiers = append(cashiers, CashierResponse{
			ID:   c.ID,
			Name: c.Name,
		})
	}
	return
}

func (s *CashierSerializer) Response() CashierResponse {
	return CashierResponse{
		ID:   s.CashierModel.ID,
		Name: s.CashierModel.Name,
	}
}
