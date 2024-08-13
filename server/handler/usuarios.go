package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
)

func ListarUsuarios(contexto *gin.Context) {

	usuarios := []schemas.Usuario{}

	if err := db.Find(&usuarios).Error; err != nil {

		return
	}

	sendSucess(contexto, "listar-usuarios", usuarios)
}
