package main

import (
	"fmt"

	"github.com/souzera/GAART/config"
	"github.com/souzera/GAART/scripts"
)

func main() {

	fmt.Println("Running CLI")

	err := config.Init()
	if err != nil {
		fmt.Printf("Error initializing the application: %v", err)
	}

	scripts.InitializeScripts()
}
