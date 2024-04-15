package organization

import (
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(route fiber.Router) {
	h := NewHandler()
	r := route.Group("organizations")

	r.Post("/", h.Create)
	r.Put("/", h.Update)
	r.Delete("/", h.Delete)
	r.Get("/tree", h.QueryTree)
}
