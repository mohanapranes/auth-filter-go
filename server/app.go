package server

import (
	"github/mohanapranes/auth-filter-go/auth/config"
	"github/mohanapranes/auth-filter-go/auth/middleware"
	"github/mohanapranes/auth-filter-go/auth/utils"
	"github/mohanapranes/auth-filter-go/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

// @title Fiber Auth API
// @version 1.0
// @description This is a sample server with authentication using Keycloak.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email mohanapraneswaran@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api
// @schemes http https

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func Start() {

	// Initialize Keycloak configuration
	keycloakConfig := config.NewKeycloakConfig()
	keycloakClient := utils.NewKeycloakClient(keycloakConfig)

	// Initialize auth middleware
	authMiddleware := middleware.NewAuthMiddleware(keycloakClient)

	app := fiber.New()

	// Register routes
	routes.RegisterRoutes(app, authMiddleware)

	// Start the server
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
