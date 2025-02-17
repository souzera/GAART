package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()";primaryKey`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deteledAt,omitempty"`
	Nome         string    `json:"nome"`
	DtNascimento time.Time `json:"dtNascimento"`
	Especie      string    `json:"especie"`
	Porte        string    `json:"porte"`
	Foto         string    `json:"foto"`
	Sexo         int       `json:"sexo"`
	Adotado      bool      `json:"adotado"`
	Castrado     bool      `json:"castrado"`
	Cidade       string    `json:"cidade"`
	Estado       string    `json:"estado"`
}

// REQUESTS

type CriarAnimalRequest struct {
	Nome         string    `json:"nome" binding:"required"`
	DtNascimento time.Time `json:"nascimento" binding:"required"`
	Especie      string    `json:"especie" binding:"required"`
	Porte        string    `json:"porte" binding:"required"`
	Foto         string    `json:"foto"`
	Sexo         int       `json:"sexo" binding:"required"`
	Adotado      bool      `json:"adotado" default:"false"`
	Castrado     bool      `json:"castrado" default:"false"`
	Cidade       string    `json:"cidade" binding:"required"`
	Estado       string    `json:"estado" binding:"required"`
}
