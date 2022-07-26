package repositories

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/models"
	"gorm.io/gorm"
)

type HotelRepository interface {
	Save(context *fiber.Ctx, db gorm.DB, hotel models.Hotel) (models.Hotel, error)
	Update(context *fiber.Ctx, db gorm.DB, hotel models.Hotel, id int) (models.Hotel, error)
	Delete(context *fiber.Ctx, db gorm.DB, hotel models.Hotel) (bool, error)
	Find(context *fiber.Ctx, db gorm.DB, hotel models.Hotel, id int) (models.Hotel, error)
	Get(context *fiber.Ctx, db gorm.DB, hotel models.Hotel) (*gorm.DB, error)
}

//type HotelRepository struct {
//	Hotel
//}
//
//func NewHotelRepository(hotel Hotel) *HotelRepository {
//	return &HotelRepository{Hotel: hotel}
//}
