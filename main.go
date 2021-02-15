package main

import (
	"fmt"
	"log"
	"os"

	"github.com/connorstevens/meta-labs-middleman-go/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"
	metalabs "github.com/meta-labs/meta-labs-go/metalabs_sdk"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
		os.Exit(3)
    }
}

func main() {
	//Get ENV Variables
	port := os.Getenv("PORT")
	metaAPIKey := os.Getenv("META_API_KEY")
	

	//Create Server
	app := fiber.New()
	client := metalabs.New(metaAPIKey)

	//Setup Middleware
	app.Use(logger.New())
	app.Use(helmet.New())

	//Setup routes
	app.Post("/auth/login", auth.Login(client))
	app.Post("/auth/reset", auth.Reset(client))

	if metaAPIKey == "" {
		log.Print("A Meta Labs API Key is required to start the server.")
		os.Exit(3)
	}

	//Start Server
	app.Listen(fmt.Sprintf(":%s", port))
	log.Printf("Server Listening at port %s", port)
}