package routes

import (
	"github.com/anevaraz/job-opportunities-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handlers.InitializeHandler()
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.POST("/candidate", handlers.CreateCandidateHandler)
		v1.GET("/candidate", handlers.GetCandidateHandler)
		v1.POST("/opportunity", handlers.CreateOpportunityHandler)
		v1.GET("/opportunity", handlers.GetOpportunityHandler)
		v1.POST("/apply", handlers.ApplyOpportunityHandler)
	}
}
