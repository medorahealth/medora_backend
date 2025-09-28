package main

import (
	"log"
	"github.com/gin-gonic/gin"

	// Adjust these import paths to match your project structure
	"github.com/medorahealth/medora_backend/internal/http/handler"
	"github.com/medorahealth/medora_backend/internal/config"
	"github.com/medorahealth/medora_backend/internal/http"
	"github.com/medorahealth/medora_backend/internal/service"
	"github.com/medorahealth/medora_backend/internal/repo"
)

// This is an EXAMPLE of how your main.go should look.
// You will need to adapt it to your actual application structure.
func main() {
	// 1. Initialize database connection using the new config package.
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// 2. Initialize your repositories.
	// The repository layer is created first, taking the DB connection.
	labRepo := repo.NewLabRepo(db)

	// 3. Initialize your services
	// This is a placeholder; you'll need a proper NewLabService function.
	labService := service.NewLabService(labRepo)

	// 4. Initialize your handlers, injecting the services they depend on.
	labHandler := handler.NewLabHandler(labService)
	// userHandler := handler.NewUserHandler(userService)
	// orderHandler := handler.NewOrderHandler(orderService)

	// 5. Initialize the Gin router
	router := gin.Default()

	// 6. Setup routes, passing the handler instances.
	http.SetupRoutes(router, labHandler /*, userHandler, orderHandler */)

	// 7. Run the server
	router.Run(":8080")
}

