package main

import (
	"fmt"
	"os"
	"github.com/connorstevens/meta-labs-middleman-go/routes"
	"github.com/connorstevens/meta-labs-middleman-go/common"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	metalabs "github.com/meta-labs/meta-labs-go/metalabs_sdk"
)

func main(){
	common.LoadEnv()
	common.Client = metalabs.New(os.Getenv("META_API_KEY"))

	app := fiber.New()
	app.Use(logger.New())
	app.Use(helmet.New())

	routes.SetupRoutes(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}