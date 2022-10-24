package main

import (
	"commerce/config"
	ud "commerce/features/user/delivery"
	ur "commerce/features/user/repository"
	us "commerce/features/user/services"
	"commerce/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)

	uRepo := ur.New(db)
	uService := us.New(uRepo)
	ud.New(e, uService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Logger.Fatal(e.Start(":8000"))
}
