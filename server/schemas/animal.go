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

// REQUESTS

type CriarAnimalRequest struct {
	Nome     string `json:"nome" binding:"required"`
	Idade    int    `json:"idade" binding:"required"`
	Especie  string `json:"especie" binding:"required"`
	Porte    string `json:"porte" binding:"required"`
	Foto     string `json:"foto"`
	Sexo     int    `json:"sexo" binding:"required"`
	Adotado  bool   `json:"adotado" default:"false"`
	Castrado bool   `json:"castrado" default:"false"`
	Cidade   string `json:"cidade" binding:"required"`
	Estado   string `json:"estado" binding:"required"`
}
