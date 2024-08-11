package schemas

import (
	"github.com/google/uuid"
)

type Usuario struct {
	ID    uuid.UUID
	Login string // unique valor único
	Senha string // TODO: ver como aplicar criptografia
	Ativo bool
}
