package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/config"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/routes"
)

var port = fmt.Sprintf(":%s", config.EnvConfig("PORT"))

func main() {
	app := fiber.New()

	routes.UserRoute(app)
	app.Listen(port)
}
