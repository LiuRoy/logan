package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DbConnection *gorm.DB
)

func init() {
	db, err := gorm.Open("mysql", MysqlUrl)
	if err != nil {
		panic(err)
	}
	DbConnection = db
}
