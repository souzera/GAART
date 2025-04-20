package main

import (
	"github.com/souzera/GAART/config"
	"github.com/souzera/GAART/router"
	_ "github.com/souzera/GAART/docs"
)

//@title           GAART API
//@version         1.0
//@description     This is a sample server for GAART API
//@termsOfService  http://swagger.io/terms/

//@contact.name  Matheus Barbosa

//@license.name  MIT License
//@license.url http://opensource.org/licenses/MIT
//@host      localhost:8000
//@BasePath  /api/v1

//@in header

var (
	logger config.Logger
)


func main() {
	logger = *config.GetLogger("GAART")

	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing the application: %v", err)
	}

	router.Initialize()
}
