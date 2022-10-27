package delivery

import (
	"commerce/config"
	"commerce/features/checkout/domain"
	"commerce/utils/common"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type checkoutHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := checkoutHandler{srv: srv}
	e.POST("/checkouts", handler.CreateCheckout(), middleware.JWT([]byte(config.JwtKey)))
	e.GET("/purchase", handler.ShowMyPurchase(), middleware.JWT([]byte(config.JwtKey)))
	e.GET("/sell", handler.ShowMySell(), middleware.JWT([]byte(config.JwtKey)))
	e.DELETE("/checkout/:id", handler.CancelOrder(), middleware.JWT([]byte(config.JwtKey)))
}

var s snap.Client

func (ch *checkoutHandler) CreateCheckout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CheckoutFormat
		userID := common.ExtractToken(c)

		var v int
		rand.Seed(time.Now().UnixNano())
		v = rand.Intn(100000)
		invo := strconv.Itoa(v)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}
		cnv := ToDomain(input)

		// 1. Initiate Snap client
		s.New("SB-Mid-server-eKSCGMJJG-IEL_LscFEV9-nP", midtrans.Sandbox)

		// 2. Initiate Snap request param
		req := &snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  invo,
				GrossAmt: int64(cnv.TotalAllPrice),
			},
			CreditCard: &snap.CreditCardDetails{
				Secure: true,
			},
		}

		// 3. Execute request create Snap transaction to Midtrans Snap API
		snapResp, _ := s.CreateTransaction(req)

		cnv.PaymentToken = snapResp.Token
		cnv.PaymentLink = snapResp.RedirectURL
		cnv.UserID = int(userID)
		cnv.NoInvoice = v
		res, err := ch.srv.Create(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success checkout", ToResponses(res, "out")))
	}
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
