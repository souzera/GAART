package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func Initialize() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Static("/uploads", "./uploads")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	
	initializeRoutes(router)
	
	router.Run("0.0.0.0:8000")
}
