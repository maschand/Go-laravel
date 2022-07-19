package controllers

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

func GetRequest(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": "Nice To Meet You and Me",
	})
}

func PostRequest(context *fiber.Ctx) error {
	body := &models.Post{}

	if err := context.BodyParser(body); err != nil {
		// Return status 400 and error message.
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Set initialized default data for book:
	body.ID = uuid.New()
	body.CreatedAt = time.Now()

	// Validate book fields.
	if err := validate.Struct(body); err != nil {
		// Return, if some fields are not valid.
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":  false,
		"result": body,
	})
}

func PutRequest(context *fiber.Ctx) error {
	body := &models.Post{}

	if err := context.BodyParser(body); err != nil {
		// Return status 400 and error message.
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Set initialized default data for book:
	body.UpdatedAt = time.Now()

	// Validate book fields.
	if err := validate.Struct(body); err != nil {
		// Return, if some fields are not valid.
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"result": body,
	})
}

func DeleteRequest(context *fiber.Ctx) error {
	body := &models.Post{}

	if err := context.BodyParser(body); err != nil {
		// Return status 400 and error message.
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Book model.
	validate := utils.NewValidator()

	// Validate book fields.
	if err := validate.Struct(body); err != nil {
		// Return, if some fields are not valid.
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"result": "Successfull Delete Your Data",
	})
}
