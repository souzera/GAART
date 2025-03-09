package router

import (
	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/handler"
	"github.com/souzera/GAART/middleware"
)

func initializeRoutes(router *gin.Engine) {

	middleware.InitializeMiddleware()
	handler.InitializeHandler()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/ping", func(contexto *gin.Context) {
			contexto.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/version", func(contexto *gin.Context) {
			contexto.JSON(200, gin.H{
				"version": "1.0.0",
			})
		})

		// Usuários
		v1.POST("/usuario", handler.CriarUsuario)
		v1.POST("/login", handler.LoginUsuario)
		v1.POST("/logout", handler.LogoutUsuario)
		v1.PATCH("/redefinir-senha", handler.RedefinirSenhaUsuario)

		// Espécies

		v1.GET("/especies", handler.ListarEspecies)
		v1.POST("/especie", handler.CriarEspecie)
		v1.PATCH("/especie", handler.AtualizarEspecie)

		// Raças

		v1.GET("/racas", handler.ListarRacas)
		v1.POST("/raca", handler.CriarRaca)
		v1.PATCH("/raca", handler.AtualizarRaca)

		// Animais

		v1.GET("/animais", middleware.RequireAuth, handler.ListarAnimais)
		v1.GET("/animal", handler.BuscarAnimal)
		v1.POST("/animal", handler.CriarAnimal)
		v1.PATCH("/animal", handler.AtualizarAnimal)

		// Endereços

		v1.GET("/enderecos", handler.ListarEnderecos)
		v1.POST("/endereco", handler.CriarEndereco)
		v1.PATCH("/endereco", handler.AtualizarEndereco)

		// Tutores

		v1.GET("/tutores", handler.ListarTutores)
		v1.POST("/tutor", handler.CriarTutor)

		// Adoções

		v1.GET("/adocoes", handler.ListarAdocoes)
		v1.POST("/adocao", handler.CriarAdocao)
	}

}
