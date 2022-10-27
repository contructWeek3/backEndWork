package repository

import (
	"commerce/features/checkout/domain"

	"gorm.io/gorm"
)

type Checkout struct {
	gorm.Model
	NoInvoice     string
	TotalAllPrice int
	PaymentLink   string
	UserID        int
}

type CheckoutDetail struct {
	gorm.Model
	ProductID  int
	CheckoutID int
	TotalStock int
	TotalPrice int
}

func ToDomainArrayOut(ci []Checkout) []domain.Core {
	var res []domain.Core
	for _, val := range ci {
		res = append(res, domain.Core{Model: gorm.Model{ID: val.ID, CreatedAt: val.CreatedAt}, NoInvoice: val.NoInvoice, TotalAllPrice: val.TotalAllPrice})
	}

	return res
}
