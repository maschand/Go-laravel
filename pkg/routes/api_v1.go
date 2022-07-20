package routes

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/gofiber/fiber/v2"
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
