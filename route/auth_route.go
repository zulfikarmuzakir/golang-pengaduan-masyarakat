package route

import "github.com/gofiber/fiber/v2"

func AuthRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
}
