package scripts

import (
	"bufio"
	"os"
	"strings"

	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
	"github.com/spf13/cobra"
)

func CriarAdmin(cmd *cobra.Command, args []string) {
	cmd.Println("\n[ ADICIONAR ADMIN ]")

	reader := bufio.NewReader(os.Stdin)

	usuario := &schemas.Usuario{}
	admin := &schemas.Admin{}

	cmd.Print("Digite um username para o admin: ")
	login, _ := reader.ReadString('\n')
	login = strings.TrimSpace(login)
	for !validateLogin(login) {
		cmd.Print("Login já existe. Tente outro username: ")
		login, _ = reader.ReadString('\n')
		login = strings.TrimSpace(login)
	}
	usuario.Login = login

	cmd.Print("Digite uma senha para o admin: ")
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
		cmd.Println("Erro ao gerar o hash da senha")
		return
	}

	usuario.Senha = hashed

	if db.Create(usuario).Error != nil {
		cmd.Println("Erro ao cria o usuário")
		return
	}

	admin.UsuarioID = usuario.ID

	cmd.Print("Digite o nome do admin: ")
	name, _ := reader.ReadString('\n')
	admin.Nome = strings.TrimSpace(name)

	if db.Create(&admin).Error != nil {

		db.Delete(&usuario)

		cmd.Println("Erro ao criar o admin")
		return
	}

	cmd.Println("Admin criado com sucesso")
}

func AtualizarAdmin(cmd *cobra.Command, args []string) {
	cmd.Println("\n[ ATUALIZAR ADMIN ]")

	login := cmd.Flag("login").Value.String()
	usuario := &schemas.Usuario{}

	if db.Where("login = ?", login).First(&usuario).RowsAffected == 0 {
		cmd.Println("Usuário não encontrado")
		return
	}

	name := cmd.Flag("nome").Value.String()
	if name == "" {
		cmd.Println("Valor inválido para o nome")
		return
	}

	admin := &schemas.Admin{}

	if db.Where("usuario_id = ?", usuario.ID).First(&admin).RowsAffected == 0 {
		cmd.Println("Admin não encontrado")
		return
	}

	admin.Nome = name

	if db.Save(&admin).Error != nil {
		cmd.Println("Error updating admin")
		return
	}

	cmd.Println("Admin updated successfully")
}
