package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	Login    string `gorm:"unique; not null"`
	Senha    string `gorm:"not null"`
	Email    string `gorm:"unique;default:null"`
	Telefone string `gorm:"unique;default:null"`
	Ativo    bool   `gorm:"default:true"`
}

// REQUESTS
type CriarUsuarioRequest struct {
	Login          string  `json:"login" binding:"required"`
	Senha          string  `json:"senha" binding:"required"`
	ConfirmarSenha string  `json:"confirmar_senha" binding:"required"`
	Email          *string `json:"email"`
	Telefone       *string `json:"telefone"`
}

type LoginUsuarioRequest struct {
	Login string `json:"login" binding:"required"`
	Senha string `json:"senha" binding:"required"`
}

type AtualizarUsuarioRequest struct {
	Login    *string `json:"login"`
	Senha    *string `json:"senha"`
	Email    *string `json:"email"`
	Telefone *string `json:"telefone"`
	Ativo    *bool   `json:"ativo"`
}

type RedefinirSenhaRequest struct {
	Token              string `json:"token" binding:"required"`
	NovaSenha          string `json:"nova" binding:"required"`
	ConfirmarNovaSenha string `json:"confirmar" binding:"required"`
}


// RESPONSES
type UsuarioResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Login     string    `json:"login"`
	Senha     string    `json:"-"`
}

type LoginUsuarioResponse struct {
	Token string `json:"token"`
}
