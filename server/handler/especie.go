package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarEspecies(contexto *gin.Context) {

	especies := []schemas.Especie{}

	logger.Info("[LISTAR-ESPECIES] Buscando as espécies", db.Find(&especies))

	if err := db.Find(&especies).Error; err != nil {
		logger.Errorf("[LISTAR-ESPECIES] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar as espécies")
		return
	}

	sendSucess(contexto, "listar-especies", especies)

}

func validarCriarEspecieRequest(request schemas.CriarEspecieRequest) error {
	if request.Nome == "" {
		return errorParamRequired("nome")
	}
	return nil
}

func CriarEspecie(contexto *gin.Context) {
	request := schemas.CriarEspecieRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarEspecieRequest(request); err != nil {
		logger.Infof("Request: %v", request)
		logger.Errorf("Error: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	especie := schemas.Especie{
		Nome: request.Nome,
	}

	if request.Genero != "" {
		especie.Genero = request.Genero
	}

	if err := db.Create(&especie).Error; err != nil {
		logger.Errorf("[CREATE-ESPECIE] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao criar a espécie")
		return
	}

	sendSucess(contexto, "criar-especie", especie)

}
