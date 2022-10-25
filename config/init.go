package config

import (
	"os"

	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func GetConfig(key string) string {

	return os.Getenv(key)
}
