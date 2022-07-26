package controllers

import (
	"github.com/chand19-af/digitels-template/app/helpers"
	"github.com/chand19-af/digitels-template/app/models"
	"github.com/chand19-af/digitels-template/app/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type HotelControllerImp struct {
	HotelService services.HotelService
}

func NewHotelController(hotelService services.HotelService) *HotelControllerImp {
	return &HotelControllerImp{HotelService: hotelService}
}

func (service HotelControllerImp) GetHotel(context *fiber.Ctx) error {
	hotel := &models.Hotel{}

	if err := context.BodyParser(hotel); err != nil {
		helpers.JsonResponse(
			context,
			"Server is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	result, err := service.HotelService.Get(context, models.Hotel{})

	if err != nil {
		helpers.JsonResponse(
			context,
			"Service is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	return helpers.JsonResponse(
		context,
		"Service is not available",
		fiber.StatusInternalServerError,
		result,
		nil)
}

func (service HotelControllerImp) FindHotel(context *fiber.Ctx) error {
	hotel := models.Hotel{}
	id, err := strconv.Atoi(context.Query("id"))

	if err := context.BodyParser(hotel); err != nil {
		helpers.JsonResponse(
			context,
			"Server is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	result, err := service.HotelService.Find(context, hotel, id)

	if err != nil {
		helpers.JsonResponse(
			context,
			"Service is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	return helpers.JsonResponse(
		context,
		"Service is not available",
		fiber.StatusInternalServerError,
		result,
		nil)
}

func (service HotelControllerImp) CreateHotel(context *fiber.Ctx) error {
	hotel := models.Hotel{}

	if err := context.BodyParser(hotel); err != nil {
		helpers.JsonResponse(
			context,
			"Server is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	result, err := service.HotelService.Create(context, hotel)

	if err != nil {
		helpers.JsonResponse(
			context,
			"Service is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	return helpers.JsonResponse(
		context,
		"Service is not available",
		fiber.StatusInternalServerError,
		result,
		nil)
}

func (service HotelControllerImp) UpdateHotel(context *fiber.Ctx) error {
	hotel := models.Hotel{}
	id, err := strconv.Atoi(context.Query("id"))

	if err := context.BodyParser(hotel); err != nil {
		helpers.JsonResponse(
			context,
			"Server is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	result, err := service.HotelService.Update(context, hotel, id)

	if err != nil {
		helpers.JsonResponse(
			context,
			"Service is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	return helpers.JsonResponse(
		context,
		"Service is not available",
		fiber.StatusInternalServerError,
		result,
		nil)
}

func (service HotelControllerImp) DeleteHotel(context *fiber.Ctx) error {
	hotel := models.Hotel{}

	if err := context.BodyParser(hotel); err != nil {
		helpers.JsonResponse(
			context,
			"Server is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	result, err := service.HotelService.Delete(context, hotel)

	if err != nil {
		helpers.JsonResponse(
			context,
			"Service is not available",
			fiber.StatusInternalServerError,
			hotel,
			err)
	}

	return helpers.JsonResponse(
		context,
		"Service is not available",
		fiber.StatusInternalServerError,
		result,
		nil)
}
