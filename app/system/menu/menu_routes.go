package menu

import (
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(route fiber.Router) {
	handler := NewHandler()
	r := route.Group("menus")

	r.Post("/", handler.Create)
	r.Put("/", handler.Update)
	r.Delete("/", handler.Delete)
	r.Get("/tree", handler.QueryTree)
}