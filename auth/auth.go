package auth

import (
	"github.com/gofiber/fiber/v2"
	metalabs "github.com/meta-labs/meta-labs-go/metalabs_sdk"
)

type validBody struct {
	License string `json:"license"`
	Machine string `json:"machine,omitempty"`
}

//Login  - POST /auth/login
func Login(client metalabs.Client) fiber.Handler {
	return func (c *fiber.Ctx) error {

		b := new(validBody)
		if err := c.BodyParser(b); err != nil {
			return c.SendStatus(400)
		}

		if b.License == "" || b.Machine == "" {
			return c.SendStatus(400)
		}

		data := map[string]interface{}{"machine": b.Machine}

		updatedLicense, err := client.UpdateKey(b.License, data)

		if err != nil {
			return c.SendStatus(400)
		}

		//&updatedLicense is the entire Meta response, I recommend filtering it down, or making your own response
		return c.Status(200).JSON(&updatedLicense)
	}
}

//Reset  - POST /auth/reset
func Reset(client metalabs.Client) fiber.Handler {
	return func (c *fiber.Ctx) error {

		b := new(validBody)
		if err := c.BodyParser(b); err != nil {
			return c.SendStatus(400)
		}

		if b.License == "" || b.Machine == "" {
			return c.SendStatus(400)
		}

		data := map[string]interface{}{"machine": nil}

		updatedLicense, err := client.UpdateKey(b.License, data)

		if err != nil {
			return c.SendStatus(400)
		}

		//&updatedLicense is the entire Meta response, I recommend filtering it down, or making your own response
		return c.Status(200).JSON(&updatedLicense)
	}
}