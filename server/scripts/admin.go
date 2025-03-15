package scripts

import (
	"bufio"
	"os"
	"strings"

	"github.com/souzera/GAART/schemas"
	"github.com/souzera/GAART/util"
	"github.com/spf13/cobra"
)

func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	return true
}

func validateConfirmPassword(password, confirmPassword string) bool {
	if password != confirmPassword {
		return false
	}
	return true
}

func validateLogin(login string) bool {
	if len(login) < 3 {
		return false
	}

	if db.Where("login = ?", login).First(&schemas.Usuario{}).RowsAffected > 0 {
		return false
	}

	return true
}

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

	cmd.Print(password, confirmPassword)

	hashed, err := util.HashSenha(strings.TrimSpace(password))
	if err != nil {
		cmd.Println("Error hashing password")
		return
	}

	usuario.Senha = hashed

	if db.Create(usuario).Error != nil {
		cmd.Println("Error creating user")
		return
	}

	admin.UsuarioID = usuario.ID

	cmd.Print("Enter the name of the admin: ")
	name, _ := reader.ReadString('\n')
	admin.Nome = strings.TrimSpace(name)

	if db.Create(&admin).Error != nil {

		// TODO: delete user if admin creation fails

		cmd.Println("Error creating admin")
		return
	}

	cmd.Println("Admin created successfully")
}
