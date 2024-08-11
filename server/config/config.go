package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
