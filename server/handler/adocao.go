package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
)

func ListarAdocoes(contexto *gin.Context) {

	adocoes := []schemas.Adocao{}

	if err := db.Preload("Animal").Preload("Tutor").Find(&adocoes).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Error ao buscar adoções")
		return
	}

	sendSucess(contexto, "listar-adocoes", adocoes)
}

func validarAdocaoRequest(request schemas.AdocaoRequest) error {

	if request.AnimalID == "" {
		return errorParamRequired("AnimalID")
	}

	if request.TutorID == "" {
		return errorParamRequired("TutorID")
	}

	if db.Where("id = ?", request.AnimalID).First(&schemas.Animal{}).Error != nil {
		return errorItemNotFound("Animal")
	}

	if db.Where("id = ?", request.TutorID).First(&schemas.Tutor{}).Error != nil {
		return errorItemNotFound("Tutor")
	}

	return nil
}

func atualizarTutorDoAnimal(animalId, tutorId uuid.UUID) error {
	animal := schemas.Animal{}

	if err := db.Where("id = ?", animalId).First(&animal).Error; err != nil {
		return err
	}

	animal.TutorID = &tutorId

	if err := db.Save(&animal).Error; err != nil {
		logger.Errorf("Error ao tentar atualizar tutor do animal: %v", err)
		return err
	}

	return nil
}

func CriarAdocao(contexto *gin.Context) {

	request := schemas.AdocaoRequest{}

	if err := contexto.BindJSON(&request); err != nil {
		logger.Errorf("Error ao fazer bind do json: %v", err)
		sendError(contexto, http.StatusBadRequest, "Error ao fazer bind do json")
		return
	}

	if err := validarAdocaoRequest(request); err != nil {
		logger.Errorf("Error ao validar request: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	animalId, err := util.ParseStringToUUID(request.AnimalID)
	if err != nil {
		logger.Errorf("Error ao converter AnimalID: %v", err)
		sendError(contexto, http.StatusBadRequest, "Error ao converter AnimalID")
		return
	}

	tutorId, err := util.ParseStringToUUID(request.TutorID)
	if err != nil {
		logger.Errorf("Error ao converter TutorID: %v", err)
		sendError(contexto, http.StatusBadRequest, "Error ao converter TutorID")
		return
	}

	adocao := schemas.Adocao{
		AnimalID: animalId,
		TutorID:  tutorId,
	}

	if request.Descricao != nil {
		adocao.Descricao = *request.Descricao
	}

	if request.Status != nil {
		adocao.Status = *request.Status
	}

	if err := db.Create(&adocao).Error; err != nil {
		logger.Errorf("Error ao criar adocao: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Error ao criar adocao")
		return
	}

	// Atualizar o id do tutor do animal
	if err := atualizarTutorDoAnimal(animalId, tutorId); err != nil {
		sendError(contexto, http.StatusInternalServerError, err.Error())
	}

	db.Preload("Animal").Preload("Tutor").First(&adocao)

	sendSucess(contexto, "criar-adocao", adocao)
}
