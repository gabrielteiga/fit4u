package routes

import (
	uc "github.com/gabrielteiga/fit4u/backend/internal/app/usecase"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) *Router {
	return &Router{
		app: app,
	}
}

func (r *Router) Init() {
	api := r.app.Group("/api/v1")
	setup(api)
}

func setup(api fiber.Router) {
	api.Get("/ping", uc.Ping)
}
