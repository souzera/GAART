package middleware

import (
	"github.com/souzera/GAART/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeMiddleware() {
	logger = config.GetLogger("GAART")
	db = config.GetDB()
}
