package repository

import (
	"commerce/features/checkout/domain"
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

func (rq *repoQuery) Checkout(newCheckout domain.Core) (domain.Core, error) {
	var resQry Checkout
	if err := rq.db.Exec("INSERT INTO checkouts (id, created_at, updated_at, deleted_at, no_invoice, total_all_price, payment_link, payment_token, user_id) values (?,?,?,?,?,?,?,?,?)",
		nil, time.Now(), time.Now(), nil, newCheckout.NoInvoice, newCheckout.TotalAllPrice, newCheckout.PaymentLink, newCheckout.PaymentToken, newCheckout.UserID).Error; err != nil {
		return domain.Core{}, err
	}
	if er := rq.db.Table("checkouts").Select("id", "no_invoice", "payment_link", "payment_token").Where("no_invoice = ?", newCheckout.NoInvoice).Model(&Checkout{}).Find(&resQry).Error; er != nil {
		return domain.Core{}, er
	}
	res := ToDomain(resQry)
	return res, nil
}

func (rq *repoQuery) Purchase(ID uint) ([]domain.Core, error) {
	var resQry []Checkout
	if err := rq.db.Table("checkouts").Select("id", "created_at", "no_invoice", "total_all_price").Where("user_id = ?", ID).Model(&Checkout{}).Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArrayOut(resQry)
	return res, nil
}

func (rq *repoQuery) Sell(ID uint) ([]domain.Core, error) {
	var resQry []Checkout
	if err := rq.db.Table("checkouts").Select("checkouts.id", "checkouts.created_at", "checkouts.no_invoice", "checkouts.total_all_price").Joins("join checkout_details on checkout_details.checkout_id=checkouts.id").Joins("join products on products.id=checkout_details.product_id").Where("products.user_id = ?", ID).Model(&Checkout{}).Find(&resQry).Error; err != nil {
		return nil, err
	}
	res := ToDomainArrayOut(resQry)
	return res, nil
}

func (rq *repoQuery) Cncl(ID int) error {
	var resQry Checkout
	if err := rq.db.Where("id = ?", ID).Delete(&resQry).Error; err != nil {
		return err
	}
	return nil
}
