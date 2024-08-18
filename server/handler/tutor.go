package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarTutores(contexto *gin.Context) {

	tutores := []schemas.Tutor{}

	if err := db.Find(&tutores).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Error ao buscar tutores")
		return
	}

	sendSucess(contexto, "listar-tutores", tutores)
}
