package delivery

import (
	cc "commerce/config"
	"commerce/features/user/domain"
	"commerce/utils/common"
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.GET("/users", handler.MyProfile(), middleware.JWT([]byte(cc.JwtKey)))
	e.PUT("/users", handler.UpdateProfile(), middleware.JWT([]byte(cc.JwtKey)))
	e.DELETE("/users", handler.Deactivate(), middleware.JWT([]byte(cc.JwtKey)))
	e.POST("/register", handler.Register())
	e.POST("/login", handler.Login())
}

func (uh *userHandler) MyProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else {
			res, err := uh.srv.MyProfile(uint(userID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusOK, SuccessResponse("success get my profile", ToResponse(res, "user")))
		}
	}
}

func (uh *userHandler) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else {
			var input EditFormat
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
					return c.JSON(http.StatusBadRequest, FailResponse("cannot bind update data"))
				}

				cnv := ToDomain(input)
				res, err := uh.srv.UpdateProfile(cnv, uint(userID))
				if err != nil {
					return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
				}
				return c.JSON(http.StatusAccepted, SuccessResponse("Success update user", ToResponse(res, "user")))
			}
			return c.JSON(http.StatusBadRequest, FailResponse("fail upload file"))
		}
	}
}

func (uh *userHandler) Deactivate() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractToken(c)
		if userID == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		} else {
			err := uh.srv.Deactivate(userID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
			}
			return c.JSON(http.StatusAccepted, FailResponse("success deactivate account"))
		}
	}
}

func (uh *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UserFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToDomain(input)
		res, err := uh.srv.Register(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, SuccessResponse("Success create new user", ToResponse(res, "res")))
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToDomain(input)
		res, err := uh.srv.Login(cnv)
		fmt.Println(res.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		res.Token = common.GenerateToken(uint(res.ID))

		return c.JSON(http.StatusAccepted, SuccessLogin("Success to login", ToResponse(res, "res")))
	}
}
