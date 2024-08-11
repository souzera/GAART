package schemas

import (
	"github.com/google/uuid"
)

/*
 * ADMIN, SUPERUSER
 * ORGANIZAÇÃO, ONGs
 */

type Admin struct {
	ID        uuid.UUID
	UsuarioID uuid.UUID
	Nome      string
	Contato   string
	Email     string
}
