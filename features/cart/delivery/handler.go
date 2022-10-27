package delivery

import (
	cc "commerce/config"
	"commerce/features/cart/domain"
	"commerce/utils/common"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type cartHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := cartHandler{srv: srv}
	// e.GET("/cart", handler.ShowAllProducts())
	e.GET("/carts/me", handler.ShowMyCart(), middleware.JWT([]byte(cc.JwtKey)))
	e.POST("/carts", handler.AddProductCart(), middleware.JWT([]byte(cc.JwtKey)))
	e.PUT("/carts/:id", handler.EditProductCart(), middleware.JWT([]byte(cc.JwtKey)))
	e.DELETE("/carts/:id", handler.DeleteProductCart(), middleware.JWT([]byte(cc.JwtKey)))
}

// func (ph *productHandler) ShowAllProducts() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		res, err := ph.srv.ShowAll()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
// 		}

// 		return c.JSON(http.StatusOK, SuccessResponse("Success get all products", ToResponses(res, "all")))
// 	}
// }

func (ch *cartHandler) ShowMyCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		res, err := ch.srv.ShowMyCart(userID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get my cart", ToResponses(res, "my")))
	}
}

func (ch *cartHandler) AddProductCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input InsertCartFormat
		userID := common.ExtractToken(c)
		input.UserID = int(userID)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := ch.srv.Add(cnv.ProductID, cnv.Stock)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		} else {

		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success create new product", ToResponses(res, "one")))
	}
}

func (ch *cartHandler) EditProductCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input InsertCartFormat
		userID := common.ExtractToken(c)
		input.UserID = int(userID)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)
		res, err := ch.srv.Add(cnv.ProductID, cnv.Stock)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		} else {

		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success create new product", ToResponses(res, "one")))
	}
}

func (ch *cartHandler) DeleteProductCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := ch.srv.Delete(ID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, FailResponse("Success delete product from cart"))
	}
}