package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
)

func ListarAnimais(contexto *gin.Context) {

	animals := []schemas.Animal{}

	// TODO: devo serializar o tutor?
	// verificar o funcionamento do Joins e criar um struct para o retorno
	if err := db.Preload("Raca.Especie").Find(&animals).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Error ao buscar animais")
	}

	sendSucess(contexto, "listar-animais", animals)
}

func BuscarAnimal(contexto *gin.Context) {
	id := contexto.Query("id")

	if id == "" {
		sendError(contexto, http.StatusBadRequest, "O id do animal é obrigatório")
		return
	}

	animal := schemas.Animal{}

	if err := db.Preload("Raca.Especie").Where("id = ?", id).First(&animal).Error; err != nil {
		sendError(contexto, http.StatusNotFound, "Animal não encontrado")
		return
	}

	sendSucess(contexto, "buscar-animal", animal)
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

	if request.DtNascimento != nil {
		logger.Infof("O usuario especificou a data de nascimento")
		animal.DtNascimento, _ = util.ParseStringToTime(*request.DtNascimento)
	}

	if request.Tutor != nil {
		logger.Infof("O usuario especificou o tutor")
		if db.Where("id = ?", request.Tutor).First(&schemas.Tutor{}).Error != nil {
			logger.Errorf("Tutor não encontrado")
			sendError(contexto, http.StatusBadRequest, "Tutor não encontrado")
			return
		}
		tutorID, _ := util.ParseStringToUUID(*request.Tutor)
		animal.TutorID = &tutorID
	}

	if request.Castrado != nil {
		logger.Infof("O usuario especificou %v no campo castrado", *request.Castrado)
		animal.Castrado = *request.Castrado
	}

	if request.Vacinado != nil {
		logger.Infof("O usuario especificou %v no campo vacinado", *request.Vacinado)
		animal.Vacinado = *request.Vacinado
	}

	logger.Infof("Animal: %v", animal)

	if err := db.Create(&animal).Error; err != nil {
		logger.Errorf("[CREATE-ANIMAL] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Error ao criar animal")
		return
	}

	sendSucess(contexto, "criar-animal", animal)
}

func AtualizarAnimal(contexto *gin.Context) {
	id := contexto.Query("id")

	if id == "" {
		logger.Errorf("O id do animal é obrigatório")
		sendError(contexto, http.StatusBadRequest, "O id do animal é obrigatório")
		return
	}

	request := schemas.AtualizarAnimalRequest{}

	if err := contexto.BindJSON(&request); err != nil {
		logger.Errorf("Error ao fazer bind do JSON: %v", err)
		sendError(contexto, http.StatusBadRequest, "Error ao fazer bind do JSON")
		return
	}

	animal := schemas.Animal{}

	if db.Where("id = ?", id).First(&animal).Error != nil {
		logger.Errorf("Animal não encontrado")
		sendError(contexto, http.StatusNotFound, "Animal não encontrado")
		return
	}

	if request.Nome != nil {
		animal.Nome = *request.Nome
	}

	if request.Sexo != nil {
		animal.Sexo = *request.Sexo
	}

	if request.Raca != nil {
		if db.Where("id = ?", *request.Raca).First(&schemas.Raca{}).Error != nil {
			logger.Errorf("Raça não encontrada")
		} else {
			animal.RacaID, _ = util.ParseStringToUUID(*request.Raca)
		}
	}

	if request.Tutor != nil {
		if db.Where("id = ?", *request.Tutor).First(&schemas.Tutor{}).Error != nil {
			logger.Errorf("Tutor não encontrado")
		} else {
			tutorID, _ := util.ParseStringToUUID(*request.Tutor)
			animal.TutorID = &tutorID
		}
	}

	if request.DtNascimento != nil {
		animal.DtNascimento, _ = util.ParseStringToTime(*request.DtNascimento)
	}

	if request.Castrado != nil {
		animal.Castrado = *request.Castrado
	}

	if request.Vacinado != nil {
		animal.Vacinado = *request.Vacinado
	}

	if err := db.Preload("Raca.Especie").Save(&animal).Error; err != nil {
		logger.Errorf("Error ao atualizar animal: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Error ao atualizar animal")
		return
	}

	sendSucess(contexto, "atualizar-animal", animal)
}
