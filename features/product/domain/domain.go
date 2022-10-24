package domain

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Core struct {
	gorm.Model
	ProductName string
	Description string
	Images      string
	Stock       int
	Price       int
	UserID      int
}

type Cores struct {
	gorm.Model
	ProductName string
	Description string
	Images      string
	Stock       int
	Price       int
	UserID      int
	Name        string
}

type Repository interface {
	Show() ([]Cores, error)
	Spesific(ID int) (Cores, error)
	My(ID uint) ([]Cores, error)
	Insert(newProduct Core) (Cores, error)
	Update(ID int, updatePost Core) (Cores, error)
	Del(ID int) error
}

type Service interface {
	ShowAll() ([]Cores, error)
	ShowSpesific(ID int) (Cores, error)
	ShowMy(ID uint) ([]Cores, error)
	Create(newProduct Core) (Cores, error)
	Edit(ID int, updatePost Core) (Cores, error)
	Delete(ID int) error
}

type Handler interface {
	ShowAllPost() echo.HandlerFunc
	ShowSpesificProduct(ID int) echo.HandlerFunc
	ShowMyProduct(ID uint) echo.HandlerFunc
	CreateProduct() echo.HandlerFunc
	EditProduct(ID int) echo.HandlerFunc
	DeletePost(ID int) echo.HandlerFunc
}
