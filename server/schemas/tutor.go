package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tutor struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	UsuarioID uuid.UUID `gorm:"type:uuid;not null" json:"usuario_id"`
	Usuario   Usuario   `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`

	EnderecoID uuid.UUID `gorm:"type:uuid;" json:"endereco_id"`
	Endereco   Endereco  `gorm:"foreignKey: EnderecoID" json:"endereco"`

	Nome      string  `gorm:"not null" json:"nome"`
	Reputacao float32 `gorm:"default:5" json:"reputacao"`

	Animais []Animal `gorm:"foreignKey:TutorID" json:"animais,omitempty"`
}

// REQUESTS
type CriarTutorRequest struct {
	UsuarioID uuid.UUID `json:"usuarioId" binding:"required"`
	Nome      string    `json:"nome" binding:"required"`
	Email     string    `json:"email"`
	Endereco  string    `json:"endereco"`
}
