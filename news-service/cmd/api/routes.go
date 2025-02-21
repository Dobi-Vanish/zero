package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func (app *Config) routes(router *fiber.App) {
	router.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		c.Set("Access-Control-Expose-Headers", "Link")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set("Access-Control-Max-Age", "300")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	})

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "pong",
		})
	})

	router.Use(AuthMiddleware)
	router.Get("/list", app.GetNewsList)
	router.Post("/edit/:Id", app.EditNews)

	log.Println("Routes are registered")
}
