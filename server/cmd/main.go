package main

import (
	"github.com/souzera/GAART/config"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.NewLogger("GAART")

	logger.Infof("hello world")
}
