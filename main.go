package main

import (
	"github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/create-go-app/fiber-go-template/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//Initialization new app
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Midtrans",
		AppName:       "Midtrans v1",
	})

	//Declare versioning or grouping
	api := app.Group("/api")

	//Declare Middleware
	middleware.FiberMiddleware(app)

	//Example Api CRUD
	api.Get("/test", controllers.GetRequest)
	api.Post("/test", controllers.PostRequest)
	api.Put("/test", controllers.PutRequest)
	api.Delete("/test", controllers.DeleteRequest)

	app.Listen(":3000")
}
