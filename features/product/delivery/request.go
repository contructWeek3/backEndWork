package delivery

import (
	"commerce/features/product/domain"
)

type ProductFormat struct {
	ProductName string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Images      string `json:"images" form:"images"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	UserID      int    `json:"user_id" form:"user_id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case ProductFormat:
		cnv := i.(ProductFormat)
		return domain.Core{ProductName: cnv.ProductName, Description: cnv.Description, Images: cnv.Images, Stock: cnv.Stock, Price: cnv.Price, UserID: cnv.UserID}
	}
	return domain.Core{}
}
