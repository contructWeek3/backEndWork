package repository

import (
	"commerce/features/product/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string
	Description string
	Images      string
	Stock       int
	Price       int
	UserID      int
}

type ProductOut struct {
	gorm.Model
	ProductName string
	Description string
	Images      string
	Stock       int
	Price       int
	UserID      int
	Name        string
}

func ToDomain(p ProductOut) domain.Cores {
	return domain.Cores{
		Model:       gorm.Model{ID: p.ID},
		ProductName: p.ProductName,
		Description: p.Description,
		Images:      p.Images,
		Stock:       p.Stock,
		Price:       p.Price,
		UserID:      p.UserID,
		Name:        p.Name,
	}
}

func ToDomainArrayOut(ai []ProductOut) []domain.Cores {
	var res []domain.Cores
	for _, val := range ai {
		res = append(res, domain.Cores{Model: gorm.Model{ID: val.ID}, ProductName: val.ProductName, Description: val.Description, Images: val.Images, Stock: val.Stock, Price: val.Price, UserID: val.UserID, Name: val.Name})
	}

	return res
}
