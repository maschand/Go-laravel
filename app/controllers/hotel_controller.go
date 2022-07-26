package controllers

import "github.com/gofiber/fiber/v2"

type HotelControlller interface {
	GetHotel(context *fiber.Ctx) error
	FindHotel(context *fiber.Ctx) error
	CreateHotel(context *fiber.Ctx) error
	UpdateHotel(context *fiber.Ctx) error
	DeleteHotel(context *fiber.Ctx) error
}
