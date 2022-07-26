package main

import (
	"github.com/chand19-af/digitels-template/pkg/middleware"
	"github.com/chand19-af/digitels-template/pkg/routes"
	"github.com/chand19-af/digitels-template/pkg/utils"
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
