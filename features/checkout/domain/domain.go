package domain

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Core struct {
	gorm.Model
	NoInvoice     string
	TotalAllPrice int
	PaymentLink   string
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
	Purchase(ID uint) ([]Core, error)
	Sell(ID uint) ([]Core, error)
	Cncl(ID int) error
}

type Service interface {
	ShowPurchase(ID uint) ([]Core, error)
	ShowSell(ID uint) ([]Core, error)
	Cancel(ID int) error
}

type Handler interface {
	ShowMyPurchase(ID uint) echo.HandlerFunc
	ShowMySell(ID uint) echo.HandlerFunc
	CancelOrder(ID int) echo.HandlerFunc
}
