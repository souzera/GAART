package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	ID    uuid.UUID
	Login string // unique valor único
	Senha string // TODO: ver como aplicar criptografia
	Ativo bool
}

type UsuarioResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Senha     string    `json:"-"`
}
