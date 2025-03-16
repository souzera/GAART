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

	parent.AddCommand(
		&cobra.Command{
			Use:   "ping",
			Short: "Ping the GAART",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println("pong")
			},
		},
		cmdCriarAdmin,
		cmdAtualizarAdmin,
	)

}
