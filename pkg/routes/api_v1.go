package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/app/controllers"
)

func V1Routes(a *fiber.App) {
	// Create routes group.
	api := a.Group("/api/v1")

	//Example Api CRUD
	api.Get("/test", controllers.GetRequest)
	api.Post("/test", controllers.PostRequest)
	api.Put("/test", controllers.PutRequest)
	api.Delete("/test", controllers.DeleteRequest)
}
