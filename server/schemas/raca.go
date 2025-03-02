package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Raca struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Nome  string `json:"nome"`
	Porte string `json:"porte"`

	EspecieID uuid.UUID `json:"especieId"`
	Especie   Especie   `json:"especie"`
}

// REQUESTs

type CriarRacaRequest struct {
	Nome    string `json:"nome" binding:"required"`
	Porte   string `json:"porte" binding:"required"`
	Especie string `json:"especie" binding:"required"`
}

// RESPONSEs

type RacaRensponse struct {
	ID      uuid.UUID `json:"id"`
	Nome    string    `json:"nome"`
	Porte   string    `json:"porte"`
	Especie Especie   `json:"especie"`
}
