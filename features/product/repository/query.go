package repository

import (
	"commerce/features/product/domain"
	"time"

	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(dbConn *gorm.DB) domain.Repository {
	return &repoQuery{
		db: dbConn,
	}
}

func (rq *repoQuery) Show() ([]domain.Cores, error) {
	var resQry []ProductOut
	if err := rq.db.Table("products").Select("products.id", "products.product_name", "products.description", "products.images", "products.stock", "products.price", "users.id AS user_id", "users.name").Joins("join users on users.id=products.user_id").Model(&ProductOut{}).Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArrayOut(resQry)
	return res, nil
}

func (rq *repoQuery) Spesific(ID int) (domain.Cores, error) {
	var resQry ProductOut
	if err := rq.db.Table("products").Select("products.id", "products.product_name", "products.description", "products.images", "products.stock", "products.price", "users.name").Joins("join users on users.id=products.user_id").Where("products.id = ?", ID).Model(&ProductOut{}).Find(&resQry).Error; err != nil {
		return domain.Cores{}, err
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) My(ID uint) ([]domain.Cores, error) {
	var resQry []ProductOut
	if err := rq.db.Table("products").Select("products.id", "products.product_name", "products.description", "products.images", "products.stock", "products.price", "users.name").Joins("join users on users.id=products.user_id").Where("users.id = ?", ID).Model(&ProductOut{}).Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArrayOut(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(newProduct domain.Core) (domain.Cores, error) {
	var resQry ProductOut
	if err := rq.db.Exec("INSERT INTO products (id, created_at, updated_at, deleted_at, product_name, description, images, stock, price, user_id) values (?,?,?,?,?,?,?,?,?,?)",
		nil, time.Now(), time.Now(), nil, newProduct.ProductName, newProduct.Description, newProduct.Images, newProduct.Stock, newProduct.Price, newProduct.UserID).Error; err != nil {
		return domain.Cores{}, err
	}
	if er := rq.db.Table("products").Select("products.id", "products.product_name", "products.description", "products.images", "products.stock", "products.price", "users.name").Joins("join users on users.id=products.user_id").Where("products.product_name = ? AND products.user_id = ?", newProduct.ProductName, newProduct.UserID).Model(&ProductOut{}).Find(&resQry).Error; er != nil {
		return domain.Cores{}, er
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Update(ID int, updateProduct domain.Core) (domain.Cores, error) {
	var resQry ProductOut
	if err := rq.db.Exec("UPDATE products SET updated_at = ?, product_name = ?, description = ?, images = ?, stock = ?, price = ? WHERE id = ?",
		time.Now(), updateProduct.ProductName, updateProduct.Description, updateProduct.Images, updateProduct.Stock, updateProduct.Price, ID).Error; err != nil {
		return domain.Cores{}, err
	}
	if er := rq.db.Table("products").Select("products.id", "products.product_name", "products.description", "products.images", "products.stock", "products.price", "users.name").Joins("join users on users.id=products.user_id").Where("products.id = ?", ID).Model(&ProductOut{}).Find(&resQry).Error; er != nil {
		return domain.Cores{}, er
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Del(ID int) error {
	var resQry Product
	if err := rq.db.Where("id = ?", ID).Delete(&resQry).Error; err != nil {
		return err
	}
	return nil
}
