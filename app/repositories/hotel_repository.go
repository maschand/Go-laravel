package repositories

import (
	"github.com/chand19-af/digitels-template/app/models"
	"github.com/gofiber/fiber/v2"
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
