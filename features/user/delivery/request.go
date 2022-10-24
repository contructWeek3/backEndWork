package delivery

import "commerce/features/user/domain"

type UserFormat struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type EditFormat struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Name     string `json:"name" form:"name"`
	Address  string `json:"address" form:"address"`
	Images   string `json:"images" form:"images"`
}

func ToDomain(i interface{}) domain.UserCore {
	switch i.(type) {
	case UserFormat:
		cnv := i.(UserFormat)
		return domain.UserCore{Username: cnv.Username, Email: cnv.Email, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.UserCore{Username: cnv.Username, Password: cnv.Password}
	case EditFormat:
		cnv := i.(EditFormat)
		return domain.UserCore{Username: cnv.Username, Email: cnv.Email, Name: cnv.Name, Phone: cnv.Phone, Address: cnv.Address, Images: cnv.Images}
	}
	return domain.UserCore{}
}
