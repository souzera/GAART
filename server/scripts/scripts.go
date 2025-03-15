package scripts

import (
	"github.com/souzera/GAART/config"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
	cmd   *cobra.Command
)

func InitializeScripts() {
	logger = config.GetLogger("GAART")
	
	db = config.GetDB()
	
	cmd = &cobra.Command{
		Use:   "gaart",
		Short: "GAART is a tool to manage the GAART",
	}

	initializeCobra(cmd)
}


