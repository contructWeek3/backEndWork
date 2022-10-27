package delivery

import (
	"commerce/features/checkout/domain"
)

type CheckoutFormat struct {
	NoInvoice     int    `json:"no_invoice" form:"no_invoice"`
	TotalAllPrice int    `json:"total_all_price" form:"total_all_price"`
	PaymentLink   string `json:"payment_link" form:"payment_link"`
	PaymentToken  string `json:"payment_token" form:"payment_token"`
	UserID        int    `json:"user_id" form:"user_id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case CheckoutFormat:
		cnv := i.(CheckoutFormat)
		return domain.Core{TotalAllPrice: cnv.TotalAllPrice}
	}
	return domain.Core{}
}
