package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tutor struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UsuarioID uuid.UUID `gorm:"type:uuid;" json:"usuario_id"`
	Usuario   Usuario   `gorm:"foreignKey:UsuarioID" json:"usuario,omitempty"`

	EnderecoID *uuid.UUID `gorm:"type:uuid;" json:"endereco_id"`
	Endereco   Endereco  `gorm:"foreignKey: EnderecoID" json:"endereco"`

	Nome      string  `gorm:"not null" json:"nome"`
	Reputacao float32 `gorm:"default:5" json:"reputacao"`

	Animais []Animal `gorm:"foreignKey:TutorID" json:"animais,omitempty"`
}

// REQUESTS
type CriarTutorRequest struct {
	Usuario   string   `json:"usuario" binding:"required"`
	Nome      string   `json:"nome" binding:"required"`
	Endereco  *string  `json:"endereco"`
	Reputacao *float32 `json:"reputacao"`
}
