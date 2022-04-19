package cashiers

import (
	"github.com/didietsuryadi95/pos/common"
	"github.com/gin-gonic/gin"
)

type CashierModelValidator struct {
	Name         string       `form:"name" json:"name" binding:"required"`
	Passcode     string       `form:"passcode" json:"passcode" binding:"required,numeric"`
	cashierModel CashierModel `json:"-"`
}

type AuthValidator struct {
	Passcode     string       `form:"passcode" json:"passcode" binding:"required,numeric"`
	cashierModel CashierModel `json:"-"`
}

func NewCashierModelValidator() CashierModelValidator {
	return CashierModelValidator{}
}

func AuthModelValidator() AuthValidator {
	return AuthValidator{}
}

func NewCashierModelValidatorFillWith(cashierModel CashierModel) CashierModelValidator {
	cashierModelValidator := NewCashierModelValidator()
	return cashierModelValidator
}

func (s *CashierModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	s.cashierModel.Name = s.Name
	s.cashierModel.Passcode = s.Passcode
	return nil
}

func (s *AuthValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, s)
	if err != nil {
		return err
	}
	return nil
}
