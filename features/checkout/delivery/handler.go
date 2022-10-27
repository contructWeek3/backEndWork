package delivery

import (
	"commerce/config"
	"commerce/features/checkout/domain"
	"commerce/utils/common"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type checkoutHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := checkoutHandler{srv: srv}
	e.GET("/purchase", handler.ShowMyPurchase(), middleware.JWT([]byte(config.JwtKey)))
	e.GET("/sell", handler.ShowMySell(), middleware.JWT([]byte(config.JwtKey)))
	e.DELETE("/checkout/:id", handler.CancelOrder(), middleware.JWT([]byte(config.JwtKey)))
}

func (ch *checkoutHandler) ShowMyPurchase() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		res, err := ch.srv.ShowPurchase(userID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get my purchase", ToResponses(res, "my")))
	}
}

func (ch *checkoutHandler) ShowMySell() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		res, err := ch.srv.ShowSell(userID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get my selling", ToResponses(res, "my")))
	}
}

func (ch *checkoutHandler) CancelOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := ch.srv.Cancel(ID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, FailResponse("Success cancel order"))
	}
}
