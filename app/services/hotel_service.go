package services

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/models"
	"gorm.io/gorm"
)

type HotelService interface {
	Get(context *fiber.Ctx, request models.Hotel) (*gorm.DB, error)
	Find(context *fiber.Ctx, request models.Hotel, id int) (models.Hotel, error)
	Create(context *fiber.Ctx, request models.Hotel) (models.Hotel, error)
	Update(context *fiber.Ctx, request models.Hotel, id int) (models.Hotel, error)
	Delete(context *fiber.Ctx, request models.Hotel) (bool, error)
}

//type HotelService struct {
//	Hotel
//}
//
//func NewHotelService(hotel Hotel) *HotelService {
//	return &HotelService{Hotel: hotel}
//}
