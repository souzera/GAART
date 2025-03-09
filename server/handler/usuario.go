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

func LoginUsuario(contexto *gin.Context) {

	request := schemas.LoginUsuarioRequest{}

	contexto.BindJSON(&request)

	usuario := schemas.Usuario{}

	if err := db.Where("login = ?", request.Login).First(&usuario).Error; err != nil {
		logger.Errorf("[LOGIN-USUARIO] Usuario não encontrado: %v", err)
		sendError(contexto, http.StatusNotFound, "Dados inválidos")
		return
	}

	if !util.VerificarSenha(request.Senha, usuario.Senha) {
		logger.Errorf("[LOGIN-USUARIO] Senha inválida")
		sendError(contexto, http.StatusUnauthorized, "Dados inválidos")
		return
	}

	token := util.GerarToken(usuario.ID.String())
	if token == "" {
		logger.Errorf("[LOGIN-USUARIO] Erro ao gerar o token")
		sendError(contexto, http.StatusInternalServerError, "Erro ao gerar o token")
		return
	}

	response := schemas.LoginUsuarioResponse{
		Token: token,
	}

	sendSucess(contexto, "login-usuario", response)
}

func LogoutUsuario(contexto *gin.Context) {
	sendSucess(contexto, "logout-usuario", nil)
}

func validarRedefinirSenhaRequest(request schemas.RedefinirSenhaRequest) error {

	if request.Token == "" {
		return errorParamRequired("token")
	}

	if request.NovaSenha == "" {
		return errorParamRequired("nova")
	}

	if request.ConfirmarNovaSenha == "" {
		return errorParamRequired("confirmar")
	}

	if request.NovaSenha != request.ConfirmarNovaSenha {
		return errorCustomMessage("senhas não conferem")
	}

	//if validarToken(request.Token) {
	//	return errorCustomMessage("token inválido")
	//}

	return nil
}

func RedefinirSenhaUsuario(contexto *gin.Context) {

	// ISSO AQUI NÃO PRESTA kkkkk
	id := contexto.Query("id")
	if id == "" {
		sendError(contexto, http.StatusBadRequest, "ID não informado")
		return
	}
	// ajeita isso aqui pião
	// sugestões de melhoria:
	// - fazer uma busca do usuário pelo token
	// como fazer isso?
	// - criar um novo campo na tabela de usuários para armazenar o token
	// - os tokens podem ser temporarios, então criar um campo para armazenar a data de expiração
	// - criar um endpoint para gerar o token
	// - criar um endpoint para validar o token
	// - jwt??

	request := schemas.RedefinirSenhaRequest{}

	if err := contexto.BindJSON(&request); err != nil {
		logger.Errorf("[REDEFINIR-SENHA-USUARIO] Error: %v", err)
		sendError(contexto, http.StatusBadRequest, "Dados inválidos")
		return
	}

	if err := validarRedefinirSenhaRequest(request); err != nil {
		logger.Errorf("[REDEFINIR-SENHA-USUARIO] Error: %v", err)
		sendError(contexto, http.StatusBadRequest, err.Error())
		return
	}

	usuario := schemas.Usuario{}

	if err := db.Where("id = ?", id).First(&usuario).Error; err != nil {
		logger.Errorf("[REDEFINIR-SENHA-USUARIO] Usuario não encontrado: %v", err)
		sendError(contexto, http.StatusNotFound, "Dados inválidos")
		return
	}

	hash, err := util.HashSenha(request.NovaSenha)
	if err != nil {
		logger.Errorf("[REDEFINIR-SENHA-USUARIO] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao redefinir a senha")
		return
	}

	usuario.Senha = hash

	if err := db.Save(&usuario).Error; err != nil {
		logger.Errorf("[REDEFINIR-SENHA-USUARIO] Error: %v", err)
		sendError(contexto, http.StatusInternalServerError, "Erro ao redefinir a senha")
		return
	}

	sendSucess(contexto, "redefinir-senha-usuario", true)
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
