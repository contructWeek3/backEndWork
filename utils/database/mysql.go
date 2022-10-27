package database

import (
	"commerce/config"
	cr "commerce/features/checkout/repository"
	pr "commerce/features/product/repository"
	ur "commerce/features/user/repository"
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(c *config.AppConfig) *gorm.DB {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser,
		c.DBPwd,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		log.Error("db config error :", err.Error())
		return nil
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&ur.User{})
	db.AutoMigrate(&pr.Product{})
	db.AutoMigrate(&cr.Checkout{})
	db.AutoMigrate(&cr.CheckoutDetail{})
}
