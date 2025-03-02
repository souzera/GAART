package schemas

import (

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Endereco struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	
	Logradouro  string         `json:"logradouro"`
	Numero      string         `json:"numero"`
	Complemento string         `json:"complemento"`
	Bairro      string         `json:"bairro"`
	Cidade      string         `json:"cidade"`
	Estado      string         `json:"estado"`
	Cep         string         `json:"cep"`
}

type CriarEnderecoRequest struct {
	Logradouro  string `json:"logradouro" binding:"required"`
	Numero      string `json:"numero" binding:"required"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"estado"`
	Cep         string `json:"cep" binding:"required"`
}
