package domain

import (
	"gorm.io/gorm"
)

type Core struct {
	gorm.Model
	NoInvoice     int
	TotalAllPrice int
	PaymentLink   string
	PaymentToken  string
	UserID        int
}

type Core2 struct {
	gorm.Model
	ProductID  int
	CheckoutID int
	TotalStock int
	TotalPrice int
}

type Repository interface {
	Checkout(newCheckout Core) (Core, error)
	Purchase(ID uint) ([]Core, error)
	Sell(ID uint) ([]Core, error)
	Cncl(ID int) error
}

type Service interface {
	Create(newCheckout Core) (Core, error)
	ShowPurchase(ID uint) ([]Core, error)
	ShowSell(ID uint) ([]Core, error)
	Cancel(ID int) error
}
