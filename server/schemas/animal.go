package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Animal struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Nome string `json:"nome"`

	RacaID uuid.UUID `json:"racaId"`
	Raca   Raca      `json:"raca"`

	TutorID *uuid.UUID `json:"tutorId"`

	DtNascimento time.Time `json:"nascimento"`
	Sexo         int       `json:"sexo" default:"0"`
	Castrado     bool      `gorm:"default:false" json:"castrado"`
	Vacinado     bool      `gorm:"default:false" json:"vacinado"`
	// TODO: Adicionar fotos
}

// REQUESTS

type CriarAnimalRequest struct {
	Nome         string  `json:"nome" binding:"required"`
	Sexo         int     `json:"sexo" binding:"required"`
	Raca         string  `json:"raca" binding:"required"`
	Tutor        *string `json:"tutor"`
	DtNascimento *string `json:"nascimento"`
	Castrado     *bool   `json:"castrado"`
	Vacinado     *bool   `json:"vacinado"`
}

type AtualizarAnimalRequest struct {
	Nome         *string `json:"nome"`
	Sexo         *int    `json:"sexo"`
	Raca         *string `json:"raca"`
	Tutor        *string `json:"tutor"`
	DtNascimento *string `json:"nascimento"`
	Castrado     *bool   `json:"castrado"`
	Vacinado     *bool   `json:"vacinado"`
}

// RESPONSES

type ListarAnimaisResponse struct {
	Data []Animal `json:"data"`
	Message string  `json:"message"`
}
