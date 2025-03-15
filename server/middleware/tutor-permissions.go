package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
)

func TutorPermissions(contexto *gin.Context) {
	logger.Infof("[TUTOR-PERMISSIONS] Middleware de permissões de tutor")

	token := contexto.GetHeader("Authorization")

	user, err := util.ValidarToken(token)

	if err != nil {
		logger.Errorf("[TUTOR-PERMISSIONS] Token inválido: %v", err)
		contexto.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		contexto.Abort()
		return
	}

	tutor := schemas.Tutor{}

	if err := db.Where("usuario_id = ?", user).First(&tutor).Error; err != nil {
		logger.Errorf("[TUTOR-PERMISSIONS] Tutor não encontrado: %v", err)
		contexto.JSON(http.StatusUnauthorized, gin.H{"error": "Tutor não encontrado"})
		contexto.Abort()
		return
	}

	logger.Infof("[TUTOR-PERMISSIONS] Tutor %v autenticado", tutor.Nome)

	contexto.Next()
}