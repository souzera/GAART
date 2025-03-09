package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequireAuth(contexto *gin.Context) {
	logger.Infof("[REQUIRE-AUTH] Middleware de autenticação")

	contexto.Next()
}