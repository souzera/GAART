package middleware

import "github.com/souzera/GAART/config"

var (
	logger *config.Logger
)

func InitializeMiddleware() {
	logger = config.GetLogger("GAART")
}
