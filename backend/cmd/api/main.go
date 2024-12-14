package main

import (
	"github.com/gabrielteiga/fit4u/backend/cmd/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	initRouter(app)
	app.Listen(":8080")
}

func initRouter(app *fiber.App) {
	router := routes.NewRouter(app)
	router.Init()
}
