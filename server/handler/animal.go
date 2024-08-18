package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarAnimais(contexto *gin.Context) {

	animais := []schemas.Animal{}

	if err := db.Find(&animais).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Error ao buscar animais")
	}

	sendSucess(contexto, "listar-animais", animais)
}
