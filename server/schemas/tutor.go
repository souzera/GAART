package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	ID        uuid.UUID
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	UsuarioID uuid.UUID `json:"usuarioId"`
	Nome      string    `json:"nome"`
	Contato   string    `json:"contato"`
	Email     string    `json:"email"`
	Endereco  string    `json:"endereco"`
	Reputacao float32   `json:"reputacao"`
}

// REQUESTS
type CriarTutorRequest struct {
	UsuarioID uuid.UUID `json:"usuarioId" binding:"required"`
	Nome      string    `json:"nome" binding:"required"`
	Contato   string    `json:"contato" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Endereco  string    `json:"endereco" binding:"required"`
	Reputacao float32   `json:"reputacao" binding:"required"`
}
