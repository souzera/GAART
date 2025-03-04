package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarEnderecos(contexto *gin.Context) {

	enderecos := []schemas.Endereco{}

	if err := db.Find(&enderecos).Error; err != nil {
		logger.Errorf("")
		sendError(contexto, http.StatusNotFound, "Erro ao listar endereços")
	}

	sendSucess(contexto, "listar-enderecos", enderecos)
}

func validarCriarEnderecoRequest(request schemas.CriarEnderecoRequest) error {
	if request.Logradouro == "" {
		return errorParamRequired("logradouro")
	}
	if request.Numero == "" {
		return errorParamRequired("numero")
	}
	if request.Cep == "" {
		return errorParamRequired("cep")
	}
	return nil
}

func CriarEndereco(contexto *gin.Context) {

	request := schemas.CriarEnderecoRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarEnderecoRequest(request); err != nil {
		logger.Infof("Request: %v", request)
		logger.Errorf("Error: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: Async Consultar api de CEP (VIA CEP)
	// tentar generalizar essa função para não depender de um serviço específico

	endereco := schemas.Endereco{
		Logradouro: request.Logradouro,
		Numero:     request.Numero,
		Cep:        request.Cep,
	}

	if request.Complemento != nil {
		endereco.Complemento = *request.Complemento
	}

	if request.Bairro != nil {
		endereco.Bairro = *request.Bairro
	}

	if request.Cidade != nil {
		endereco.Cidade = *request.Cidade
	}

	if request.Estado != nil {
		endereco.Estado = *request.Estado
	}

	if err := db.Create(&endereco).Error; err != nil {
		logger.Errorf("[CREATE-ENDERECO] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao criar o endereço")
		return
	}

	sendSucess(contexto, "criar-endereco", endereco)
}

func AtualizarEndereco(contexto *gin.Context) {
	id := contexto.Query("id")

	if id == "" {
		logger.Errorf("ID não informado")
		sendError(contexto, http.StatusBadRequest, "ID não informado")
		return
	}

	request := schemas.AtualizarEnderecoRequest{}

	if err := contexto.BindJSON(&request); err != nil {
		logger.Errorf("Erro ao dar bind no JSON: %v", err)
		sendError(contexto, http.StatusBadRequest, "Erro ao dar bind no JSON")
		return
	}

	endereco := schemas.Endereco{}

	if db.Where("id = ?", id).First(&endereco).Error != nil {
		logger.Errorf("Endereço não encontrado")
		sendError(contexto, http.StatusNotFound, "Endereço não encontrado")
		return
	}

	if request.Logradouro != nil {
		endereco.Logradouro = *request.Logradouro
	}

	if request.Numero != nil {
		endereco.Numero = *request.Numero
	}

	if request.Cep != nil {
		endereco.Cep = *request.Cep
	}

	if request.Complemento != nil {
		endereco.Complemento = *request.Complemento
	}

	if request.Bairro != nil {
		endereco.Bairro = *request.Bairro
	}

	if request.Cidade != nil {
		endereco.Cidade = *request.Cidade
	}

	if request.Estado != nil {
		endereco.Estado = *request.Estado
	}

	if err := db.Save(&endereco).Error; err != nil {
		logger.Errorf("Erro ao atualizar o endereço: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao atualizar o endereço")
		return
	}

	sendSucess(contexto, "atualizar-endereco", endereco)
}
