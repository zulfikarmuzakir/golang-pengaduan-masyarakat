package main

import (
	"fmt"
	"github.com/zulfikarmuzakir/golang-pengaduan-masyarakat/config"
)

var port = fmt.Sprintf(":%s", config.EnvConfig("PORT"))

func main() {
	app := fiber.New()

	route.AuthRoute(app)
	app.Listen(port)
}
