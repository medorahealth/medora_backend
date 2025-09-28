package router

import (
	"github.com/medorahealth/medora_backend/internal/http/handler" // Import your controllers

	"github.com/gin-gonic/gin"
)

// SetupLabRoutes now accepts a LabHandler instance to correctly register the route methods.
func SetupLabRoutes(router *gin.RouterGroup, labHandler *handler.LabHandler) {
	// Create a new group for lab-specific routes, e.g., /api/v1/labs
	labRoutes := router.Group("/labs")
	{
		// Define the routes and map them to the methods on the labHandler instance
		labRoutes.POST("/", labHandler.RegisterLab)

		// NOTE: You will need to implement GetAllLabs, GetLabByID, etc. as methods
		// on your LabHandler struct for these lines to work.
		// labRoutes.GET("/", labHandler.GetAllLabs)
		// labRoutes.GET("/:id", labHandler.GetLabByID)
		// labRoutes.PUT("/:id", labHandler.UpdateLab)
		// labRoutes.DELETE("/:id", labHandler.DeleteLab)
	}
}

