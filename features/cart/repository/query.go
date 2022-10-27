package repository

import (
	"commerce/features/cart/domain"
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

func (rq *repoQuery) MyCart(ID uint) ([]domain.Core, error) {
	var resQry []Cart
	if err := rq.db.Table("cart").Select("products.id", "products.product_name", "products.description", "products.images", "products.stock", "products.price", "users.id AS user_id").Joins("join carts on carts.product_id=products.id").Where("user.id = ?", ID).Model(&Cart{}).Find(&resQry).Error; err != nil {
		return []domain.Core{}, err
	}
	res := ToDomainArrayOut(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(ProductID, Stock int) (domain.Core, error) {
	var resQry Cart
	if err := rq.db.Exec("INSERT INTO Cart (product_id,stock, created_at, updated_at, deleted_at) values (?,?,?,?,?)",
		nil, time.Now(), time.Now(), nil, ProductID, Stock).Error; err != nil {
		return domain.Core{}, err
	}
	if er := rq.db.Table("cart").Select("product.id", "product_name", "description", "images", "stock", "price").Where("product.id = ?", ProductID).Model(&Cart{}).Find(&resQry).Error; er != nil {
		return domain.Core{}, er
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Update(ProductID, Stock int) (domain.Core, error) {
	var resQry Cart
	if err := rq.db.Exec("UPDATE cart SET updated_at = ?, product_id = ?, stock = ? WHERE id = ?",
		time.Now(), ProductID, Stock, ProductID).Error; err != nil {
		return domain.Core{}, err
	}
	if er := rq.db.Table("cart").Select("product.id", "product_stock").Where("product.id = ?", ProductID).Model(&Cart{}).Find(&resQry).Error; er != nil {
		return domain.Core{}, er
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Del(ID int) error {
	var resQry Cart
	if err := rq.db.Where("id = ?", ID).Delete(&resQry).Error; err != nil {
		return err
	}
	return nil
}
