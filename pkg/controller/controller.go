package controller

import (
	"github/mohanapranes/auth-filter-go/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func UserController(ctx *fiber.Ctx) error {
	value, err := services.AccessByUser()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error accessing user data")
	}
	return ctx.SendString(value)
}

func AdminController(ctx *fiber.Ctx) error {
	value, err := services.AccessByAdmin()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Error accessing admin data")
	}
	return ctx.SendString(value)
}
