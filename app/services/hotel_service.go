package services

import (
	"github.com/chand19-af/digitels-template/app/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HotelService interface {
	Get(context *fiber.Ctx, request models.Hotel) (*gorm.DB, error)
	Find(context *fiber.Ctx, id uint) (models.Hotel, error)
	Create(context *fiber.Ctx, request models.Hotel) (models.Hotel, error)
	Update(context *fiber.Ctx, request models.Hotel, id uint) (models.Hotel, error)
	Delete(context *fiber.Ctx, id uint) (bool, error)
}
