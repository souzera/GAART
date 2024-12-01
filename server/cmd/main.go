package main

import (
	"github.com/souzera/GAART/config"
	"github.com/souzera/GAART/router"
)

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
