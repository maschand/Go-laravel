package helpers

import "github.com/gofiber/fiber/v2"

func JsonResponse(context *fiber.Ctx, message string, code int, data interface{}, err error) error {
	return context.Status(code).JSON(fiber.Map{
		"result":  data,
		"error":   err,
		"message": message,
	})
}
