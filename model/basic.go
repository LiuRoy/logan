package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"logan/config"
)

var (
	DbConnection *gorm.DB
)

func init() {
	db, err := gorm.Open("mysql", config.MysqlUrl)
	if err != nil {
		panic(err)
	}
	DbConnection = db
}
