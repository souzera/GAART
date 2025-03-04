package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
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

func validarCriarUsuarioRequest(request schemas.CriarUsuarioRequest) error {
	if request.Login == "" {
		return errorParamRequired("login")
	}
	if request.Senha == "" {
		return errorParamRequired("senha")
	}
	if request.ConfirmarSenha == "" {
		return errorParamRequired("confirmar_senha")
	}
	if request.Senha != request.ConfirmarSenha {
		return errorCustomMessage("senhas não conferem")
	}
	return nil
}

func CriarUsuario(contexto *gin.Context) {

	request := schemas.CriarUsuarioRequest{}

	contexto.BindJSON(&request)

	if err := validarCriarUsuarioRequest(request); err != nil {
		logger.Infof("Request: %v", request)
		logger.Errorf("Error: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	hash, err := util.HashSenha(request.Senha)
	if err != nil {
		logger.Errorf("Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao criar o usuário")
		return
	}

	usuario := schemas.Usuario{
		Login: request.Login,
		Senha: hash,
	}

	if request.Email != nil {
		usuario.Email = *request.Email
	}

	if request.Telefone != nil {
		usuario.Telefone = *request.Telefone
	}

	if err := db.Create(&usuario).Error; err != nil {
		logger.Errorf("[CREATE-USUARIO] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao criar o usuário")
		return
	}

	sendSucess(contexto, "criar-usuario", usuario)
}
