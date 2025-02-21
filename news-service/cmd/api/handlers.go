package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

// GetNewsList gets list of all news
func (app *Config) GetNewsList(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid limit parameter",
		})
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid offset parameter",
		})
	}

	news, total, err := app.Repo.GetNewsList(limit, offset)
	if err != nil {
		logsrus.Errorf("Error during fetching the news: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch news",
		})
	}

	return c.JSON(fiber.Map{
		"data":   news,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

// EditNews edits the news with provided information
func (app *Config) EditNews(c *fiber.Ctx) error {
	var requestPayload struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		Categories []int  `json:"categories"`
	}

	if err := c.BodyParser(&requestPayload); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if len(requestPayload.Title) > 200 {
		logsrus.Errorf("Title must be less than 200 characters, title's length: %v", requestPayload.Title)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title must be less than 200 characters",
		})
	}

	if len(requestPayload.Content) > 1000 {
		logsrus.Errorf("Content must be less than 1000 characters, content's length: %v", requestPayload.Content)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Title must be less than 200 characters",
		})
	}
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logsrus.Errorf("Error converting ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	log.Println("Updating news with ID:", id)
	err = app.Repo.EditNews(id, requestPayload.Title, requestPayload.Content, requestPayload.Categories)
	if err != nil {
		logsrus.Errorf("Error updating news: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update news",
		})
	}

	return c.JSON(fiber.Map{
		"message": "News updated successfully",
	})
}
