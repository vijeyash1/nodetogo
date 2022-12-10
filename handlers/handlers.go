package handlers

import (
	"gorm.io/gorm"
)

type Config struct {
	Db *gorm.DB
}

func NewConfig(Db *gorm.DB) *Config {
	return &Config{Db}
}
