package router

import (
	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/handler"
)

func initializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	v1 := router.Group("/api/v1/")
	{
		v1.GET("/ping", func(contexto *gin.Context) {
			contexto.JSON(200, gin.H{
				"message": "pong",
			})
		})

		// Usuários
		v1.POST("/usuario", handler.CriarUsuario)

		// Espécies

		v1.GET("/especies", handler.ListarEspecies)
		v1.POST("/especie", handler.CriarEspecie)
		v1.PATCH("/especie", handler.AtualizarEspecie)

		// Raças

		v1.GET("/racas", handler.ListarRacas)
		v1.POST("/raca", handler.CriarRaca)

		// Animais

		v1.GET("/animais", handler.ListarAnimais)
		v1.GET("/animal", handler.BuscarAnimal)
		v1.POST("/animal", handler.CriarAnimal)
		v1.PATCH("/animal", handler.AtualizarAnimal)

		// Endereços

		v1.GET("/enderecos", handler.ListarEnderecos)
		v1.POST("/endereco", handler.CriarEndereco)

		// Tutores

		v1.GET("/tutores", handler.ListarTutores)
		v1.POST("/tutor", handler.CriarTutor)
	}

}
