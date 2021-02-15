package routes

import (
	"fmt"
	"log"

	"github.com/connorstevens/meta-labs-middleman-go/common"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes (Linter)
func SetupRoutes(app *fiber.App) {
	app.Post("/auth/login", Login())
	app.Post("/auth/reset", Reset())
}

// Body (Linter)
type Body struct {
	License string `json:"license"`
	Machine string `json:"machine,omitempty"`
}

// Login (Linter)
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body Body
		
		if err := c.BodyParser(&body); err != nil || body.License == "" || body.Machine == "" {
			log.Println(err)
			return c.Status(400).JSON(fiber.Map{
				"message": "Body Malformed",
			})
		}

		updatedLicense, err := common.Client.UpdateKey(body.License, map[string]interface{}{"machine": body.Machine})

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": fmt.Sprint(err),
			})
		}

		return c.Status(200).JSON(updatedLicense)
	}
}

// Reset (Linter)
func Reset() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body Body

		if err := c.BodyParser(&body); err != nil || body.License == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Body Malformed",
			})
		}

		_, err := common.Client.UpdateKey(body.License, map[string]interface{}{"machine": nil})

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": fmt.Sprint(err),
			})
		}

		return c.SendStatus(200)
	}
}