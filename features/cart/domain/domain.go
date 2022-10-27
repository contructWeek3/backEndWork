package domain

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Core struct {
	gorm.Model
	ProductName string
	Images      string
	Stock       int
	Price       int
	Total_Item  int
	Total_Price int
	UserID      int
	ProductID   int
	Name        string
}

type Repository interface {
	Show() ([]Core, error)
	MyCart(ID uint) ([]Core, error)
	Insert(ProductID int, Stock int) (Core, error)
	Update(ProductID int, updateData Core) (Core, error)
	Del(ID int) error
}

type Service interface {
	ShowAll() ([]Core, error)
	ShowMyCart(ID uint) ([]Core, error)
	Add(ProductID, Stock int) (Core, error)
	Edit(ProductID int, updateData Core) (Core, error)
	Delete(ID int) error
}

type Handler interface {
	ShowAllCart() echo.HandlerFunc
	ShowMyCart(ID uint) echo.HandlerFunc
	AddProduct(ProductID, Stock int) echo.HandlerFunc
	EditProduct(ProductID, Stock int) echo.HandlerFunc
	DeletePost(ID int) echo.HandlerFunc
}
