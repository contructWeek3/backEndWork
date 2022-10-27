package delivery

import (
	"commerce/features/checkout/domain"
	"time"
)

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}

type Response struct {
	ID           uint   `json:"id"`
	NoInvoice    int    `json:"invoice"`
	PaymentLink  string `json:"payment_link"`
	PaymentToken string `json:"payment_token"`
}

type Responses struct {
	ID            uint      `json:"id"`
	NoInvoice     int       `json:"invoice"`
	TotalAllPrice int       `json:"total_price"`
	CreatedAt     time.Time `json:"date"`
}

func ToResponses(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "out":
		val := core.(domain.Core)
		res = Response{ID: val.ID, NoInvoice: val.NoInvoice, PaymentLink: val.PaymentLink, PaymentToken: val.PaymentToken}
	case "my":
		var arr []Responses
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, Responses{ID: val.ID, NoInvoice: val.NoInvoice, TotalAllPrice: val.TotalAllPrice, CreatedAt: val.CreatedAt})
		}
		res = arr
	}
	return res
}
