package common

import (
	"commerce/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GenerateToken(id uint) string {
	claim := make(jwt.MapClaims)
	claim["authorized"] = true
	claim["id"] = id
	claim["expired"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		log.Error("error on token signed string", err.Error())
		return "cannot generate token"
	}
	return str
}

func ExtractToken(c echo.Context) uint {
	token := c.Get("user").(*jwt.Token)
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		fmt.Print(uint(claim["id"].(float64)))
		return uint(claim["id"].(float64))
	}
	return 0
}
