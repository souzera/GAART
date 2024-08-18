package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarAdocoes(contexto *gin.Context) {

	adocoes := []schemas.Adocao{}

	if err := db.Find(&adocoes).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Error ao buscar adoções")
		return
	}

	sendSucess(contexto, "listar-adocoes", adocoes)
}
