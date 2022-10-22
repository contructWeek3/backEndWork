package repository

import (
	"commerce/features/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Name     string
	Email    string
	Phone    string
	Address  string
	Images   string
	Password string
}

func FromDomain(du domain.UserCore) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Username: du.Username,
		Email:    du.Email,
		Name:     du.Name,
		Phone:    du.Phone,
		Address:  du.Address,
		Images:   du.Images,
		Password: du.Password,
	}
}

func ToDomain(u User) domain.UserCore {
	return domain.UserCore{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Name,
		Phone:    u.Phone,
		Address:  u.Address,
		Images:   u.Images,
		Password: u.Password,
	}
}
