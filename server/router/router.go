package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	router.Static("/uploads", "./uploads")
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	
	initializeRoutes(router)
	
	router.Run(":8000")
}
