package config

import (
	"fmt"

	"github.com/rezaahmadk/dts-Implementation-MVC-Golang/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("username:password@/db_name?charset=utf8&parseTime=True&loc=Local")), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect Database" + err.Error())
	}
	db.AutoMigrate(new(model.Account), new(model.Transaction))

	DB = db

	return db
}
