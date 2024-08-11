package config

import (
	"fmt"

	"github.com/souzera/GAART/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres(host string, user string, password string, dbName string, port string, sslmode string, timezone string) (*gorm.DB, error) {

	logger := GetLogger("GAART")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=%s", host, user, password, dbName, port, sslmode, timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error ao conectar com o banco de dados: %v", err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	err = MigrateAll(db)
	if err != nil {
		logger.Errorf("Erros ao criar tabelas: %v", err)
		return nil, err
	}

	return db, err
}

func MigrateUsuario(db *gorm.DB) error {
	return db.AutoMigrate(&schemas.Usuario{})
}

func MigrateAll(db *gorm.DB) error {
	return nil
}
