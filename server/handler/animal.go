package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarAnimais(contexto *gin.Context) {

	animals := []schemas.Animal{}

	logger.Info("ListarAnimais", animals)

	if err := db.Find(&animals).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Error ao buscar animais")
	}

	sendSucess(contexto, "listar-animais", animals)
}

func validarCriarAnimalRequest(request schemas.CriarAnimalRequest) error {
	if request.Nome == "" {
		return errorParamRequired("nome")
	}
	if request.Cidade == "" {
		return errorParamRequired("cidade")
	}
	return nil
}

func CriarAnimal(contexto *gin.Context) {
	request := schemas.CriarAnimalRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarAnimalRequest(request); err != nil {
		logger.Infof("Request: %v", request)
		logger.Errorf("Error: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	animal := schemas.Animal{
		Nome:         request.Nome,
		DtNascimento: request.DtNascimento,
		Especie:      request.Especie,
		Porte:        request.Porte,
		Foto:         request.Foto,
		Sexo:         request.Sexo,
		Adotado:      request.Adotado,
		Castrado:     request.Castrado,
		Cidade:       request.Cidade,
		Estado:       request.Estado,
	}

	if err := db.Create(&animal).Error; err != nil {
		logger.Errorf("[CREATE-ANIMAL] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Error ao criar animal")
		return
	}

	sendSucess(contexto, "criar-animal", animal)
}
