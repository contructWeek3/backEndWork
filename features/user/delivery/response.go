package delivery

import "commerce/features/user/domain"

type Response struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Responses struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Images   string `json:"images"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "res":
		cnv := core.(domain.UserCore)
		res = Response{Username: cnv.Username, Token: cnv.Token}
	case "user":
		cnv := core.(domain.UserCore)
		res = Responses{Username: cnv.Username, Email: cnv.Email, Phone: cnv.Phone, Name: cnv.Name, Address: cnv.Address, Images: cnv.Images}
	}
	return res
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessLogin(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}
