package utils

import "github.com/gofiber/fiber/v2"

func ResponseSuccess(c *fiber.Ctx, message string, results interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
		"result":  results,
	})
}

func ResponseError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}
