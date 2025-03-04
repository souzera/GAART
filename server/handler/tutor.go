package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
)

func ListarTutores(contexto *gin.Context) {

	tutores := []schemas.Tutor{}

	if err := db.Preload("Usuario").Preload("Endereco", "id IS NOT NULL").Find(&tutores).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Error ao buscar tutores")
		return
	}

	sendSucess(contexto, "listar-tutores", tutores)
}

func validarCriarTutorRequest(request schemas.CriarTutorRequest) error {
	if request.Usuario == "" {
		return errorParamRequired("usuarioId")
	}

	if request.Nome == "" {
		return errorParamRequired("nome")
	}

	return nil
}

func verificarEndereco(enderecoID string) error {
	if err := db.Where("id = ?", enderecoID).First(&schemas.Endereco{}).Error; err != nil {
		return errorCustomMessage("Endereço não encontrado")
	}
	return nil
}

func CriarTutor(contexto *gin.Context) {

	request := schemas.CriarTutorRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarTutorRequest(request); err != nil {
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	id, err := util.ParseStringToUUID(request.Usuario)
	if err != nil {
		sendError(contexto, http.StatusBadRequest, "Não foi possível converter o id do usuário")
		return
	}

	tutor := schemas.Tutor{
		UsuarioID: id,
		Nome:      request.Nome,
	}

	if request.Reputacao != nil {
		tutor.Reputacao = *request.Reputacao
	}

	if request.Endereco != nil {
		if err := verificarEndereco(*request.Endereco); err != nil {
			sendError(contexto, http.StatusBadRequest, err.Error())
			return
		}

		enderecoId, err := util.ParseStringToUUID(*request.Endereco)
		if err != nil {
			sendError(contexto, http.StatusBadRequest, "Não foi possível converter o id do endereço")
			return
		}

		tutor.EnderecoID = &enderecoId
	}

	if err := db.Create(&tutor).Error; err != nil {
		logger.Errorf("[CREATE-TUTOR] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Error ao adicionar tutor")
		return
	}

	sendSucess(contexto, "tutor-criado", tutor)
}
