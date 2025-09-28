package http

import (
	"github.com/gin-gonic/gin"
	"github.com/medorahealth/medora_backend/internal/http/handler"
	approuter "github.com/medorahealth/medora_backend/internal/http/router"

	
)

// SetupRoutes now accepts handler instances to pass them to the specific route setups.
// NOTE: You will need to create and pass UserHandler and OrderHandler as well.
func SetupRoutes(engine *gin.Engine, labHandler *handler.LabHandler /*, userHandler *handler.UserHandler, orderHandler *handler.OrderHandler */) {
	// Create a new router group for the API, versioned as v1.
	// This helps in versioning your API in the future (e.g., /api/v2).
	apiV1 := engine.Group("/api/v1")
	{
		// Pass the specific handler instance to its corresponding route setup function.
		approuter.SetupLabRoutes(apiV1, labHandler)

		// You would do the same for users and orders once they follow the same handler pattern
		// SetupUserRoutes(apiV1, userHandler)
		// SetupOrderRoutes(apiV1, orderHandler)
	}
}
