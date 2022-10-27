package services

import (
	"commerce/features/checkout/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
)

type CheckoutService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &CheckoutService{
		qry: repo,
	}
}

func (cs *CheckoutService) ShowPurchase(ID uint) ([]domain.Core, error) {
	res, err := cs.qry.Purchase(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("No Data")
		}
	}

	return res, nil
}

func (cs *CheckoutService) ShowSell(ID uint) ([]domain.Core, error) {
	res, err := cs.qry.Sell(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("No Data")
		}
	}

	return res, nil
}

func (cs *CheckoutService) Cancel(ID int) error {
	err := cs.qry.Cncl(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return errors.New("No Data")
		}
	}
	return nil
}
