package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Login string    `gorm:"unique; not null"`
	Senha string    `gorm:"not null"`
	Ativo bool      `gorm:"default:true"`
}

type CriarUsuarioRequest struct {
	Login          string `json:"login" binding:"required"`
	Senha          string `json:"senha" binding:"required"`
	ConfirmarSenha string `json:"confirmar_senha" binding:"required"`
}

// REQUESTS
type UsuarioResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Login     string    `json:"login"`
	Senha     string    `json:"-"`
}
