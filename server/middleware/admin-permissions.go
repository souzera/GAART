package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
)

func AdminPermissions(contexto *gin.Context) {
	logger.Infof("[ADMIN-PERMISSIONS] Middleware de permissões de administrador")

	token := contexto.GetHeader("Authorization")

	user, err := util.ValidarToken(token)

	if err != nil {
		logger.Errorf("[ADMIN-PERMISSIONS] Token inválido: %v", err)
		contexto.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		contexto.Abort()
		return
	}

	admin := schemas.Admin{}

	if err := db.Where("usuario_id = ?", user).First(&admin).Error; err != nil {
		logger.Errorf("[ADMIN-PERMISSIONS] Admin não encontrado: %v", err)
		contexto.JSON(http.StatusUnauthorized, gin.H{"error": "Admin não encontrado"})
		contexto.Abort()
		return
	}

	logger.Infof("[ADMIN-PERMISSIONS] Admin %v autenticado", admin.Nome)

	contexto.Status(http.StatusOK)
	contexto.Next()
}
