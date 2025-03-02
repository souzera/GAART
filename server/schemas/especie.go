package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Especie struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Nome      string `json:"nome"`
	Genero    string `json:"genero"`
	Domesitco bool   `gorm:"default:true" json:"domestico"`
}

// REQUESTs

type CriarEspecieRequest struct {
	Nome      string `json:"nome" binding:"required"`
	Genero    string `json:"genero"`
	Domestico bool   `json:"domestico"`
}

type AtualizarEspecieRequest struct {
	Nome      *string `json:"nome"`
	Genero    *string `json:"genero"`
	Domestico *bool   `json:"domestico"`
}

// RESPONSEs
