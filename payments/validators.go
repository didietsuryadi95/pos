package payments

import (
	"github.com/didietsuryadi95/pos/common"
	"github.com/gin-gonic/gin"
)

type PaymentModelValidator struct {
	Name         string       `form:"name" json:"name" binding:"required"`
	Type         string       `form:"type" json:"type" binding:"required"`
	Logo         string       `form:"logo" json:"logo"`
	paymentModel PaymentModel `json:"-"`
}

func NewPaymentModelValidator() PaymentModelValidator {
	return PaymentModelValidator{}
}

func NewPaymentModelValidatorFillWith(paymentModel PaymentModel) PaymentModelValidator {
	paymentModelValidator := NewPaymentModelValidator()
	paymentModelValidator.Name = paymentModel.Name
	paymentModelValidator.Type = paymentModel.Type
	paymentModelValidator.Logo = paymentModel.Logo
	return paymentModelValidator
}

func (s *PaymentModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.paymentModel.Name = s.Name
	s.paymentModel.Type = s.Type
	s.paymentModel.Logo = s.Logo
	return nil
}
