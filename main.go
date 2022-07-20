package main

import (
	"github.com/create-go-app/fiber-go-template/pkg/middleware"
	"github.com/create-go-app/fiber-go-template/pkg/routes"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
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
	routes.V1Routes(app)

	//Declare Middleware
	middleware.FiberMiddleware(app)
	utils.LoadEnvironment()

	//Start Server
	utils.StartServerWithGracefulShutdown(app)
}
