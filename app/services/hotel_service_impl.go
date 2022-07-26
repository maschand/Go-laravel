package services

import (
	"github.com/chand19-af/digitels-template/app/models"
	"github.com/chand19-af/digitels-template/app/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HotelServiceImpl struct {
	HotelRepository repositories.HotelRepository
	DB              *gorm.DB
}

func NewHotelService(hotelRepository repositories.HotelRepository, DB *gorm.DB) *HotelServiceImpl {
	return &HotelServiceImpl{
		HotelRepository: hotelRepository,
		DB:              DB,
	}
}

func (repository *HotelServiceImpl) Get(context *fiber.Ctx, request models.Hotel) (*gorm.DB, error) {
	result, err := repository.HotelRepository.Get(context, *repository.DB, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repository *HotelServiceImpl) Create(context *fiber.Ctx, request models.Hotel) (models.Hotel, error) {
	result, err := repository.HotelRepository.Save(context, *repository.DB, request)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (repository *HotelServiceImpl) Update(context *fiber.Ctx, request models.Hotel, id int) (models.Hotel, error) {
	result, err := repository.HotelRepository.Update(context, *repository.DB, request, id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (repository *HotelServiceImpl) Delete(context *fiber.Ctx, request models.Hotel) (bool, error) {
	result, err := repository.HotelRepository.Delete(context, *repository.DB, request)
	if err != nil {
		return result, err
	}

	return true, nil
}

func (repository *HotelServiceImpl) Find(context *fiber.Ctx, request models.Hotel, id int) (models.Hotel, error) {
	result, err := repository.HotelRepository.Find(context, *repository.DB, request, id)
	if err != nil {
		return result, err
	}

	return result, nil
}
