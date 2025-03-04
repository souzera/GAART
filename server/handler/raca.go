package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
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

func AtualizarRaca(contexto *gin.Context){

	id := contexto.Query("id")

	if id == "" {
		logger.Errorf("ID não informado")
		sendError(contexto, http.StatusBadRequest, "ID não informado")
		return
	}

	request := schemas.AtualizarRacaRequest{}

	if err := contexto.BindJSON(&request); err != nil {
		logger.Errorf("Erro ao dar bind no JSON: %v", err)
		sendError(contexto, http.StatusBadRequest, "Erro ao dar bind no JSON")
		return
	}

	raca := schemas.Raca{}
	if db.Where("id = ?", id).First(&raca).Error != nil {
		logger.Errorf("Raça não encontrada")
		sendError(contexto, http.StatusNotFound, "Raça não encontrada")
		return
	}

	if request.Nome != nil {
		raca.Nome = *request.Nome
	}

	if request.Porte != nil {
		raca.Porte = *request.Porte
	}

	if request.Especie != nil {
		if db.Where("id = ?", *request.Especie).First(&schemas.Especie{}).Error != nil {
			logger.Errorf("Espécie não encontrada")
		} else {
			especieId, _ := util.ParseStringToUUID(*request.Especie)
			raca.EspecieID = especieId
		}
	}

	if err := db.Save(&raca).Error; err != nil {
		logger.Errorf("[UPDATE-RACA] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao atualizar a raça")
		return
	}

	if err := db.Preload("Especie").First(&raca).Error; err != nil {
		logger.Errorf("[UPDATE-RACA] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar a raça atualizada")
		return
	}

	sendSucess(contexto, "atualizar-raca", raca)
}