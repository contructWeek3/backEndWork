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
	return &ProductService{
		qry: repo,
	}
}

func (cs *CartService) ShowAll() ([]domain.Cores, error) {
	res, err := ps.qry.Show()
	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "table") {
			return nil, errors.New("Database Error")
		} else if strings.Contains(err.Error(), "found") {
			return nil, errors.New("No Data")
		}
	}

	if len(res) == 0 {
		log.Info("No Data")
		return nil, errors.New("No Data")
	}
	return res, nil
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

func (ps *ProductService) ShowMy(ID uint) ([]domain.Cores, error) {
	res, err := ps.qry.My(ID)
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

func (ps *ProductService) Create(newProduct domain.Core) (domain.Cores, error) {
	res, err := ps.qry.Insert(newProduct)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Cores{}, errors.New("Rejected from Database")
		}

		return domain.Cores{}, errors.New("Some Problem on Database")
	}

	return res, nil
}

func (ps *ProductService) Edit(ID int, updateProduct domain.Core) (domain.Cores, error) {
	res, err := ps.qry.Update(ID, updateProduct)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return domain.Cores{}, errors.New("Rejected from Database")
		}

		return domain.Cores{}, errors.New("Some Problem on Database")
	}

	return res, nil
}

func (ps *ProductService) Delete(ID int) error {
	err := ps.qry.Del(ID)
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
