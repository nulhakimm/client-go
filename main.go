package main

import (
	"github.com/faridlan/client-go/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./public/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	router.SetupRoute(app)

	app.Listen(":3000")
}
