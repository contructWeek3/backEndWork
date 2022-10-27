package repository

import (
	"commerce/features/cart/domain"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ProductID   int
	ProductName string
	Stock       int
	Price       int
	UserID      int
}

func ToDomain(c Cart) domain.Core {
	return domain.Core{
		Model:  gorm.Model{ID: c.ID},
		Stock:  c.Stock,
		UserID: c.UserID,
	}
}

func ToDomainArrayOut(ai []Cart) []domain.Core {
	var res []domain.Core
	for _, val := range ai {
		res = append(res, domain.Core{
			Model:       gorm.Model{ID: val.ID},
			ProductName: val.ProductName,
			Stock:       val.Stock,
			Price:       val.Price,
			UserID:      val.UserID,
		})
	}

	return res
}
