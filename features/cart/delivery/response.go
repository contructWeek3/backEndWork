package delivery

import (
	"commerce/features/cart/domain"
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

type Responses struct {
	ID             uint   `json:"id"`
	Product_Name   string `json:"product_name"`
	Product_Images string `json:"images"`
	Product_Stock  int    `json:"product_stock"`
	Product_Price  int    `json:"product_price"`
	UserID         int    `json:"id_seller"`
	Name           string `json:"seller"`
}

func ToResponses(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "one":
		val := core.(domain.Core)
		res = Responses{
			ID:             val.ID,
			Product_Name:   val.ProductName,
			Product_Images: val.Images,
			Product_Stock:  val.Stock,
			Product_Price:  val.Price,
			UserID:         val.UserID,
			Name:           val.Name}
	case "my":
		var arr []Responses
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			arr = append(arr, Responses{
				ID:             val.ID,
				Product_Name:   val.ProductName,
				Product_Images: val.Images,
				Product_Stock:  val.Stock,
				Product_Price:  val.Price,
				UserID:         val.UserID,
				Name:           val.Name,
			})
		}
		res = arr
	}
	return res
}
