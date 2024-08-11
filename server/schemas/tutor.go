package schemas

import (
	"github.com/google/uuid"
)

type Tutor struct {
	ID        uuid.UUID
	UsuarioID uuid.UUID
	Nome      string
	Contato   string
	Email     string
	Endereco  string
	Reputacao float32
}
