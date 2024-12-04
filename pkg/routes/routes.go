package routes

import (
	"github/mohanapranes/auth-filter-go/auth/middleware"
	"github/mohanapranes/auth-filter-go/pkg/controller"

	_ "github/mohanapranes/auth-filter-go/docs" // This is where Swag will generate docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterRoutes(app *fiber.App, authMiddleware *middleware.AuthMiddleware) {

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api", authMiddleware.Authenticate())

	// @Summary User endpoint
	// @Description Protected endpoint for regular users
	// @Tags user
	// @Accept json
	// @Produce json
	// @Security Bearer
	// @Success 200 {object} Response
	// @Failure 401 {object} ErrorResponse
	// @Failure 403 {object} ErrorResponse
	// @Router /user [get]
	api.Get("/user", middleware.RequireRoles("user"), controller.UserController)

	// @Summary Admin endpoint
	// @Description Protected endpoint for admin users
	// @Tags admin
	// @Accept json
	// @Produce json
	// @Security Bearer
	// @Success 200 {object} Response
	// @Failure 401 {object} ErrorResponse
	// @Failure 403 {object} ErrorResponse
	// @Router /admin [get]
	api.Get("/admin", middleware.RequireRoles("admin"), controller.AdminController)
}
