package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
)

func RequireAuth(contexto *gin.Context) {
	logger.Infof("[REQUIRE-AUTH] Middleware de autenticação")

	token := contexto.GetHeader("Authorization")

	user, err := util.ValidarToken(token)

	if err != nil {
		logger.Errorf("[REQUIRE-AUTH] Token inválido: %v", err)
		contexto.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		contexto.Abort()
		return
	}

	usuario := schemas.Usuario{}

	if err := db.Where("id = ?", user).First(&usuario).Error; err != nil {
		logger.Errorf("[REQUIRE-AUTH] Usuário não encontrado: %v", err)
		contexto.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
		contexto.Abort()
		return
	}

	logger.Infof("[REQUIRE-AUTH] Usuário %v autenticado", user)

	contexto.Next()
}
