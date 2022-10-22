package repository

import (
	"commerce/features/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) GetMyUser(userID uint) (domain.UserCore, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "id = ?", userID).Error; err != nil {
		log.Error("error on get my user", err.Error())
		return domain.UserCore{}, err
	}
	res := ToDomain(resQuery)
	return res, nil
}

func (rq *repoQuery) Update(updatedUser domain.UserCore, userID uint) (domain.UserCore, error) {
	var cnv User = FromDomain(updatedUser)
	if err := rq.db.Table("users").Where("id = ?", userID).Updates(&cnv).Error; err != nil {
		log.Error("error on updating user", err.Error())
		return domain.UserCore{}, err
	}

	res := ToDomain(cnv)
	return res, nil
}

func (rq *repoQuery) Delete(userID uint) error {
	var data User
	if err := rq.db.Delete(&data, "id = ?", userID).Error; err != nil {
		log.Error("error on deleting user", err.Error())
		return err
	}
	return nil
}

func (rq *repoQuery) AddUser(newUser domain.UserCore) (domain.UserCore, error) {
	var cnv User = FromDomain(newUser)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on adding user", err.Error())
		return domain.UserCore{}, err
	}
	newUser = ToDomain(cnv)
	return newUser, nil
}

func (rq *repoQuery) GetUser(existUser domain.UserCore) (domain.UserCore, error) {
	var resQuery User
	if err := rq.db.First(&resQuery, "username = ?", existUser.Username).Error; err != nil {
		log.Error("error on get user login", err.Error())
		return domain.UserCore{}, nil
	}
	res := ToDomain(resQuery)
	return res, nil
}
