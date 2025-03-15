package scripts

import (
	"github.com/spf13/cobra"
)

func initializeCobra(rootCmd *cobra.Command) {

	addCommands(rootCmd)

	rootCmd.Execute()
}

func addCommands(parent *cobra.Command) {

	parent.AddCommand(
		&cobra.Command{
			Use:   "ping",
			Short: "Ping the GAART",
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Println("pong")
			},
		},

		&cobra.Command{
			Use:   "create-admin",
			Short: "Create an admin",
			Run:   CriarAdmin,
		},
	)

}
