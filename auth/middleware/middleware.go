package middleware

import (
	"github/mohanapranes/auth-filter-go/auth/models"
	"github/mohanapranes/auth-filter-go/auth/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	keycloakClient *utils.KeycloakClient
}

func NewAuthMiddleware(keycloakClient *utils.KeycloakClient) *AuthMiddleware {
	return &AuthMiddleware{
		keycloakClient: keycloakClient,
	}
}

func (m *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "No authorization header",
			})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid authorization header format",
			})
		}

		token := tokenParts[1]
		tokenInfo, err := m.keycloakClient.IntrospectToken(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Failed to validate token",
			})
		}

		if !tokenInfo.Active {
			return c.Status(401).JSON(fiber.Map{
				"error": "Token is not active",
			})
		}

		// Store token info in context for later use
		c.Locals("user", tokenInfo)

		return c.Next()
	}
}

func RequireRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenInfo, ok := c.Locals("user").(*models.TokenIntrospection)
		if !ok {
			return c.Status(401).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}

		// Check if user has any of the required roles
		hasRole := false
		for _, requiredRole := range roles {
			for _, userRole := range tokenInfo.RealmRoles.Roles {
				if requiredRole == userRole {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		if !hasRole {
			return c.Status(403).JSON(fiber.Map{
				"error": "Insufficient permissions",
			})
		}

		return c.Next()
	}
}
