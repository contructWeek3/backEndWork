package delivery

import (
	"commerce/features/cart/domain"
)

type InsertCartFormat struct {
	Stock     int `json:"stock" form:"stock"`
	ProductID int `json:"product_id" form:"product_id"`
	UserID    int `json:"user_id" form:"user_id"`
}

type CartFormat struct {
	ProductID   int    `json:"product_id" form:"product_id"`
	ProductName string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	UserID      int    `json:"user_id" form:"user_id"`
}

func ToDomain(i interface{}) domain.Core {
	switch i.(type) {
	case InsertCartFormat:
		cnv := i.(InsertCartFormat)
		return domain.Core{
			ProductID: cnv.ProductID,
			Stock:     cnv.Stock,
			UserID:    cnv.UserID}
	case CartFormat:
		cnv := i.(CartFormat)
		return domain.Core{
			ProductID:   cnv.ProductID,
			ProductName: cnv.ProductName,
			Stock:       cnv.Stock,
			UserID:      cnv.UserID}
	}
	return domain.Core{}
}
