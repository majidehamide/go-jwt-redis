package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", GetConfig("DB_USERNAME"), GetConfig("DB_PASSWORD"), GetConfig("DB_HOST"), GetConfig("DB_PORT"), GetConfig("DB_DATABASE"))
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	return db, err
}
