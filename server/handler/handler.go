package handler

import (
	"github.com/souzera/GAART/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("GAART")
	db = config.GetDB()
}
