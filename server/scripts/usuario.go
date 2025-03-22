package scripts

import (
	"bufio"
	"os"
	"strings"

	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
	"github.com/spf13/cobra"
)

func CriaUsuario(cmd *cobra.Command, args []string) {
	cmd.Println("\n[ ADICIONAR USUARIO ]")

	reader := bufio.NewReader(os.Stdin)

	usuario := &schemas.Usuario{}

	cmd.Print("Digite um username para o usuario: ")
	login, _ := reader.ReadString('\n')
	login = strings.TrimSpace(login)
	for !validateLogin(login) {
		cmd.Print("Login já existe. Tente outro username: ")
		login, _ = reader.ReadString('\n')
		login = strings.TrimSpace(login)
	}
	usuario.Login = login

	cmd.Print("Digite uma senha para o usuario: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	for !validatePassword(password) {
		cmd.Print("Senha muito curta. Digite Novamente: ")
		password, _ = reader.ReadString('\n')
	}

	cmd.Print("Confirmar Senha: ")
	confirmPassword, _ := reader.ReadString('\n')
	confirmPassword = strings.TrimSpace(confirmPassword)
	for !validateConfirmPassword(password, confirmPassword) {

		cmd.Print("As senhas não correspondem. Tente novamente: ")
		password, _ = reader.ReadString('\n')
		password = strings.TrimSpace(password)
		for !validatePassword(password) {
			cmd.Print("Senha muito curta. Digite Novamente: ")
			password, _ = reader.ReadString('\n')
		}

		cmd.Print("Confirmar Senha: ")
		confirmPassword, _ = reader.ReadString('\n')
		confirmPassword = strings.TrimSpace(confirmPassword)
	}

	hashed, err := util.HashSenha(strings.TrimSpace(password))
	if err != nil {
		cmd.Println("Erro ao gerar Hash da senha")
		return
	}

	usuario.Senha = hashed

	if db.Create(usuario).Error != nil {
		cmd.Println("Erro ao criar o usuario")
		return
	}

	cmd.Println("Usuario criado com sucesso")

}

func RedefinirSenha(cmd *cobra.Command, args []string) {
	cmd.Println("\n[ REDEFINIR SENHA ]")

	reader := bufio.NewReader(os.Stdin)

	cmd.Print("Digite um username para o usuario: ")
	login, _ := reader.ReadString('\n')
	login = strings.TrimSpace(login)

	usuario := &schemas.Usuario{}
	if db.Where("login = ?", login).First(usuario).RowsAffected == 0 {
		cmd.Println("Usuario não encontrado")
		return
	}

	cmd.Println("Digite a senha do usuario:")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	for !util.VerificarSenha(password, usuario.Senha) {
		cmd.Println("Senha incorreta. Tente novamente")
		password, _ = reader.ReadString('\n')
		password = strings.TrimSpace(password)
	}

	cmd.Print("Digite uma nova senha para o usuario: ")
	newPassword, _ := reader.ReadString('\n')
	newPassword = strings.TrimSpace(newPassword)
	for !validatePassword(newPassword) {
		cmd.Print("Senha muito curta. Digite Novamente: ")
		newPassword, _ = reader.ReadString('\n')
	}

	cmd.Print("Confirmar Senha: ")
	confirmPassword, _ := reader.ReadString('\n')
	confirmPassword = strings.TrimSpace(confirmPassword)
	for !validateConfirmPassword(newPassword, confirmPassword) {

		cmd.Print("As senhas não correspondem. Tente novamente: ")
		newPassword, _ = reader.ReadString('\n')
		newPassword = strings.TrimSpace(newPassword)
		for !validatePassword(newPassword) {
			cmd.Print("Senha muito curta. Digite Novamente: ")
			newPassword, _ = reader.ReadString('\n')
		}

		cmd.Print("Confirmar Senha: ")
		confirmPassword, _ = reader.ReadString('\n')
		confirmPassword = strings.TrimSpace(confirmPassword)
	}

	hashed, err := util.HashSenha(strings.TrimSpace(newPassword))
	if err != nil {
		cmd.Println("Erro ao gerar Hash da senha")
		return
	}

	usuario.Senha = hashed

	if db.Save(usuario).Error != nil {
		cmd.Println("Erro ao salvar a nova senha")
		return
	}

	cmd.Println("Senha redefinida com sucesso")
}
