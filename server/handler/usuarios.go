package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarUsuarios(contexto *gin.Context) {

	usuarios := []schemas.Usuario{}

	if err := db.Find(&usuarios).Error; err != nil {
		sendError(contexto, http.StatusInternalServerError, "Erro ao buscar o usuários")
		return
	}

	sendSucess(contexto, "listar-usuarios", usuarios)
}

type ListarUsuarioResponse struct {
	Message string `json:"message"`
	Data    []schemas.UsuarioResponse
}
