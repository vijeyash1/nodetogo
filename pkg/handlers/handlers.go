package handlers

import (
	"gorm.io/gorm"
)

type APIHandler struct {
	Db *gorm.DB
}

func NewAPIHandler(Db *gorm.DB) *APIHandler {
	return &APIHandler{Db}
}
