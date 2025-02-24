package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Raca struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`

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
