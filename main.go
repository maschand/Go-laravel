package main

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")

	//Example Api CRUD
	api.Get("/test", controllers.GetRequest)
	api.Post("/test", controllers.PostRequest)
	api.Put("/test", controllers.PutRequest)
	api.Delete("/test", controllers.DeleteRequest)

	app.Listen(":3000")
}
