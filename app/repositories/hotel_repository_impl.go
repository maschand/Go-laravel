package repositories

import (
	"github.com/chand19-af/digitels-template/app/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HotelRepositoryImpl struct {
}

func NewHotelRepository() *HotelRepositoryImpl {
	return &HotelRepositoryImpl{}
}

func (repository *HotelRepositoryImpl) Save(context *fiber.Ctx, db gorm.DB, hotel models.Hotel) (models.Hotel, error) {
	if err := db.Create(&hotel).Error; err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (repository *HotelRepositoryImpl) Update(context *fiber.Ctx, db gorm.DB, hotel models.Hotel, id int) (models.Hotel, error) {
	if err := db.Model(&hotel).Where("ID = ?", id).Updates(&hotel).Error; err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (repository *HotelRepositoryImpl) Delete(context *fiber.Ctx, db gorm.DB, hotel models.Hotel) (bool, error) {
	if err := db.Delete(&hotel).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repository *HotelRepositoryImpl) Find(context *fiber.Ctx, db gorm.DB, hotel models.Hotel, id int) (models.Hotel, error) {
	if err := db.First(&hotel, id).Error; err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (repository *HotelRepositoryImpl) Get(context *fiber.Ctx, db gorm.DB, hotel models.Hotel) (*gorm.DB, error) {
	result := db.Find(&hotel)
	if err := result.Error; err != nil {
		return nil, err
	}
	return result, nil
}
