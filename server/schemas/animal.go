package schemas

import (
	_ "time"

	"github.com/google/uuid"
)

type Animal struct {
	ID       uuid.UUID
	Nome     string
	Idade    int
	Especie  string
	Porte    string
	Foto     string
	Sexo     int
	Adotado  bool
	Castrado bool
	Cidade   string
	Estado   string
}
