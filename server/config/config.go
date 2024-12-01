package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	host := "localhost"
	user := "postgres"
	password := "postgres"
	dbName := "gaart"
	port := "5432"
	sslmode := "disable"
	timezone := "America/Fortaleza"

	db, err = InitializePostgres(host, user, password, dbName, port, sslmode, timezone)
	if err != nil {
		return fmt.Errorf("Error ao inicializar o banco: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
