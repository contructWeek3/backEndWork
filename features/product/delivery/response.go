package delivery

import (
	"commerce/features/product/domain"
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
	ID          uint   `json:"id"`
	ProductName string `json:"name"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	UserID      int    `json:"id_seller"`
	Name        string `json:"seller"`
}

func ToResponses(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "one":
		val := core.(domain.Cores)
		res = Responses{ID: val.ID, ProductName: val.ProductName, Description: val.Description, Images: val.Images, Stock: val.Stock, Price: val.Price, UserID: val.UserID, Name: val.Name}
	case "out":
		val := core.(domain.Cores)
		res = Responses{ID: val.ID, ProductName: val.ProductName, Description: val.Description, Images: val.Images, Stock: val.Stock, Price: val.Price}
	case "all":
		var arr []Responses
		cnv := core.([]domain.Cores)
		for _, val := range cnv {
			arr = append(arr, Responses{ID: val.ID, ProductName: val.ProductName, Description: val.Description, Images: val.Images, Stock: val.Stock, Price: val.Price, UserID: val.UserID, Name: val.Name})
		}
		res = arr
	case "my":
		var arr []Responses
		cnv := core.([]domain.Cores)
		for _, val := range cnv {
			arr = append(arr, Responses{ID: val.ID, ProductName: val.ProductName, Description: val.Description, Images: val.Images, Stock: val.Stock, Price: val.Price})
		}
		res = arr
	}
	return res
}
