package main

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/d6825/golang_template/pkg/middleware"
	"gitlab.com/d6825/golang_template/pkg/routes"
	"gitlab.com/d6825/golang_template/pkg/utils"
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
