package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Nome      string    `json:"nome"`

	RacaID uuid.UUID `json:"racaId"`
	Raca   Raca      `json:"raca"`

	Tutor *Tutor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	DtNascimento time.Time `json:"nascimento"`
	Sexo         int       `json:"sexo" default:"0"`
	Castrado     bool      `json:"castrado" default:"false"`
	Vacinado     bool      `json:"vacinado" default:"false"`
	// TODO: Adicionar fotos
}

// REQUESTS

type CriarAnimalRequest struct {
}
