package services

import (
	"commerce/features/cart/domain"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
)

type CartService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &CartService{
		qry: repo,
	}
}

func (cs *CartService) ShowMyCart(ID uint) ([]domain.Core, error) {
	res, err := cs.qry.MyCart(ID)
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return []domain.Core{}, errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return []domain.Core{}, errors.New("No Data")
		}
	}

	return res, nil
}

func (cs *CartService) AddCart(ProductID, Stock int) (domain.Core, error) {
	res, err := cs.qry.Insert(ProductID, Stock)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("Rejected from Database")
		}

		return domain.Core{}, errors.New("Some Problem on Database")
	}

	return res, nil
}

func (cs *CartService) EditCart(ProductID, Stock int) (domain.Core, error) {
	res, err := cs.qry.Update(ProductID, Stock)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Core{}, errors.New("Rejected from Database")
		}

		return domain.Core{}, errors.New("Some Problem on Database")
	}

	return res, nil
}

func (cs *CartService) Delete(ID int) error {
	err := cs.qry.Del(ID)
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
