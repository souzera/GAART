package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarRacas(contexto *gin.Context) {

	racas := []schemas.Raca{}

	if err := db.Preload("Especie").Find(&racas).Error; err != nil {
		logger.Errorf("[LISTAR-RACAS] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar as raças")
		return
	}

	sendSucess(contexto, "listar-racas", racas)
}

func validarCriarRacaRequest(request schemas.CriarRacaRequest) error {
	if request.Nome == "" {
		return errorParamRequired("nome")
	}
	if request.Porte == "" {
		return errorParamRequired("porte")
	}
	if request.Especie == "" {
		return errorParamRequired("especie")
	}

	if db.Where("id = ?", request.Especie).First(&schemas.Especie{}).Error != nil {
		return errorItemNotFound("especie")
	}

	return nil
}

func CriarRaca(contexto *gin.Context) {

	request := schemas.CriarRacaRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarRacaRequest(request); err != nil {
		logger.Infof("Request: %v", request)
		logger.Errorf("Error: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	especie := schemas.Especie{}
	db.Where("id = ?", request.Especie).First(&especie)

	raca := schemas.Raca{
		Nome:      request.Nome,
		Porte:     request.Porte,
		EspecieID: especie.ID,
	}

	if err := db.Create(&raca).Error; err != nil {
		logger.Errorf("[CREATE-RACA] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao criar a raça")
		return
	}

	sendSucess(contexto, "criar-raca", raca)
}
