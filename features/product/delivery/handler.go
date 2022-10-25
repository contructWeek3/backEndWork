package delivery

import (
	cc "commerce/config"
	"commerce/features/product/domain"
	"commerce/utils/common"
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type productHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := productHandler{srv: srv}
	e.GET("/products", handler.ShowAllProducts())
	e.GET("/products/:id", handler.ShowSpesificProduct())
	e.GET("/products/me", handler.ShowMyProduct(), middleware.JWT([]byte(cc.JwtKey)))
	e.POST("/products", handler.CreateProduct(), middleware.JWT([]byte(cc.JwtKey)))
	e.PUT("/products/:id", handler.EditProduct(), middleware.JWT([]byte(cc.JwtKey)))
	e.DELETE("/products/:id", handler.DeleteProduct(), middleware.JWT([]byte(cc.JwtKey)))
}

func (ph *productHandler) ShowAllProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ph.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success get all products", ToResponses(res, "all")))
	}
}

func (ps *productHandler) ShowSpesificProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ps.srv.ShowSpesific(ID)

		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get detail product", ToResponses(res, "one")))
	}
}

func (ps *productHandler) ShowMyProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		res, err := ps.srv.ShowMy(userID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get my post", ToResponses(res, "my")))
	}
}

func (ph *productHandler) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ProductFormat
		userID := common.ExtractToken(c)
		input.UserID = int(userID)

		cfg, errDef := config.LoadDefaultConfig(context.TODO())
		if errDef != nil {
			var erroDef string = "Error: "
			erroDef += erroDef
			return c.JSON(http.StatusBadRequest, FailResponse(erroDef))
		}

		client := s3.NewFromConfig(cfg)
		uploader := manager.NewUploader(client)

		isSuccess := true
		file, er := c.FormFile("images")
		if er != nil {
			isSuccess = false
		} else {
			src, err := file.Open()
			if err != nil {
				isSuccess = false
			} else {
				result, errImg := uploader.Upload(context.TODO(), &s3.PutObjectInput{
					Bucket: aws.String("be12project3bucket"),
					Key:    aws.String(file.Filename),
					Body:   src,
					ACL:    "public-read",
				})

				if errImg != nil {
					return c.JSON(http.StatusBadRequest, FailResponse("Berhasil Upload Images"))
				}

				input.Images = result.Location
			}
			defer src.Close()
		}

		if isSuccess {
			if err := c.Bind(&input); err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}
			cnv := ToDomain(input)
			res, err := ph.srv.Create(cnv)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}

			return c.JSON(http.StatusCreated, SuccessResponse("Success create new product", ToResponses(res, "out")))
		}
		return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
	}
}

func (ph *productHandler) EditProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		var input ProductFormat
		userID := common.ExtractToken(c)
		input.UserID = int(userID)

		cfg, errDef := config.LoadDefaultConfig(context.TODO())
		if errDef != nil {
			var erroDef string = "Error: "
			erroDef += erroDef
			return c.JSON(http.StatusBadRequest, FailResponse(erroDef))
		}

		client := s3.NewFromConfig(cfg)
		uploader := manager.NewUploader(client)

		isSuccess := true
		file, er := c.FormFile("images")
		if er != nil {
			isSuccess = false
		} else {
			src, err := file.Open()
			if err != nil {
				isSuccess = false
			} else {
				result, errImg := uploader.Upload(context.TODO(), &s3.PutObjectInput{
					Bucket: aws.String("be12project3bucket"),
					Key:    aws.String(file.Filename),
					Body:   src,
					ACL:    "public-read",
				})

				if errImg != nil {
					return c.JSON(http.StatusBadRequest, FailResponse("Berhasil Upload Images"))
				}

				input.Images = result.Location
			}
			defer src.Close()
		}

		if isSuccess {
			if err := c.Bind(&input); err != nil {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
			}
			cnv := ToDomain(input)
			res, err := ph.srv.Edit(ID, cnv)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}

			return c.JSON(http.StatusCreated, SuccessResponse("Success update product", ToResponses(res, "out")))
		}
		return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
	}
}

func (ph *productHandler) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))
		err := ph.srv.Delete(ID)
		if err != nil {
			log.Error(err.Error())
			if strings.Contains(err.Error(), "table") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
		}
		return c.JSON(http.StatusOK, FailResponse("Success delete product"))
	}
}
