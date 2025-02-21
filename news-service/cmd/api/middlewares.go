package main

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware checks headers has Authorization.
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		logsrus.Warn("Authorization header is missing")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		logsrus.Warn("Invalid Authorization header format")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	token := parts[1]
	if token != "my-secret-token" {
		logsrus.Warnf("Invalid token: %s", token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.Next()
}
