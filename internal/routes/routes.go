package routes

import (
	"github.com/anevaraz/go-jobs-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handlers.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/candidate", handlers.GetCandidate)
		v1.POST("/candidate", handlers.Create)
	}
}
