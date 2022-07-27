package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/providers"
)

func V1Routes(a *fiber.App) {
	// Create routes group.
	api := a.Group("/api/v1")

	hotelController := providers.InitializedServer()

	//Example Api CRUD
	api.Get("/test", hotelController.GetHotel)
	api.Get("/test/find", hotelController.FindHotel)
	api.Post("/test", hotelController.CreateHotel)
	api.Put("/test", hotelController.UpdateHotel)
	api.Delete("/test", hotelController.DeleteHotel)
}
