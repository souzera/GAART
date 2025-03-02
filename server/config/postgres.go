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

func MigrateAdmin(db *gorm.DB) error {
	return db.AutoMigrate(&schemas.Admin{})
}

func MigrateEndereco(db *gorm.DB) error {
	return db.AutoMigrate((&schemas.Endereco{}))
}

func MigrateTutor(db *gorm.DB) error {
	return db.AutoMigrate(&schemas.Tutor{})
}

func MigrateEspecie(db *gorm.DB) error {
	return db.AutoMigrate(&schemas.Especie{})
}

func MigrateRaca(db *gorm.DB) error {
	return db.AutoMigrate(&schemas.Raca{})
}

func MigrateAnimal(db *gorm.DB) error {
	return db.AutoMigrate(&schemas.Animal{})
}

func MigrateAdocao(db *gorm.DB) error {
	return db.AutoMigrate(&schemas.Adocao{})
}

func MigrateAll(db *gorm.DB) error {

	if err := MigrateUsuario(db); err != nil {
		return err
	}

	if err := MigrateAdmin(db); err != nil {
		return err
	}

	if err := MigrateEndereco(db); err != nil {
		return err
	}

	if err := MigrateTutor(db); err != nil {
		return err
	}

	if err := MigrateEspecie(db); err != nil {
		return err
	}

	if err := MigrateRaca(db); err != nil {
		return err
	}

	if err := MigrateAnimal(db); err != nil {
		return err
	}

	if err := MigrateAdocao(db); err != nil {
		return err
	}

	return nil
}
