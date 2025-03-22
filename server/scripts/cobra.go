package scripts

import (
	"github.com/spf13/cobra"
)

var (
	cmdCriarAdmin     *cobra.Command
	cmdAtualizarAdmin *cobra.Command
)

func initializeCobra(rootCmd *cobra.Command) {

	addCommands(rootCmd)

	rootCmd.Execute()
}

func addCommands(parent *cobra.Command) {

	cmdCriarAdmin = &cobra.Command{
		Use:   "create-admin",
		Short: "Create an admin",
		Run:   CriarAdmin,
	}

	cmdAtualizarAdmin = &cobra.Command{
		Use:   "update-admin",
		Short: "Update an admin",
		Run:   AtualizarAdmin,
	}

	cmdAtualizarAdmin.Flags().StringP("login", "l", "", "Login para procurar o admin")
	cmdAtualizarAdmin.MarkFlagRequired("login")
	cmdAtualizarAdmin.Flags().StringP("nome", "n", "", "Novo nome do admin")

	cmdCriarUsuario := &cobra.Command{
		Use:   "create-user",
		Short: "Create an user",
		Run:   CriaUsuario,
	}

	cmdRedefinirSenha := &cobra.Command{
		Use:   "reset-password",
		Short: "Reset the password",
		Run:   RedefinirSenha,
	}

	parent.AddCommand(
		&cobra.Command{
			Use:   "ping",
			Short: "Ping the GAART",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println("pong")
			},
		},

		&cobra.Command{
			Use:  "version",
			Long: "Print the version of GAART",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println("GAART v0.0.1")
			},
		},

		&cobra.Command{
			Use:   "meta",
			Short: "Print the metadata",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println("GAART - Gerenciador de Acesso e Autenticação\n Autor: Matheus Barbosa")
			},
		},

		// ADMIN

		cmdCriarAdmin,
		cmdAtualizarAdmin,

		// USUARIO

		cmdCriarUsuario,
		cmdRedefinirSenha,
	)

}
