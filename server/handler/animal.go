package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarAnimais(contexto *gin.Context) {

	animals := []schemas.Animal{}

	if err := db.Preload("Raca.Especie").Find(&animals).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Error ao buscar animais")
	}

	sendSucess(contexto, "listar-animais", animals)
}

func validarCriarAnimalRequest(request schemas.CriarAnimalRequest) error {
	if request.Nome == "" {
		return errorParamRequired("nome")
	}
	if request.Raca == "" {
		return errorParamRequired("raca")
	}
	if !(request.Sexo == 0 || request.Sexo == 1) {
		return errorCustomMessage("valor inválido para o campo sexo")
	}

	if db.Where("id = ?", request.Raca).First(&schemas.Raca{}).Error != nil {
		return errorItemNotFound("raça")
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

	raca := schemas.Raca{}
	db.Where("id = ?", request.Raca).First(&raca)

	animal := schemas.Animal{
		Nome:   request.Nome,
		RacaID: raca.ID,
		Sexo:   request.Sexo,
	}

	if err := db.Create(&animal).Error; err != nil {
		logger.Errorf("[CREATE-ANIMAL] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Error ao criar animal")
		return
	}

	sendSucess(contexto, "criar-animal", animal)
}
