package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	err = godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error ao carregar o arquivo .env: %v",
			err)
	}

	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	port := os.Getenv("DATABASE_PORT")
	sslmode := os.Getenv("DATABASE_SSL_MODE")
	timezone := os.Getenv("DATABASE_TIMEZONE")

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
