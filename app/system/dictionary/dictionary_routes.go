package dictionary

import (
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(route fiber.Router) {
	h := NewHandler()
	r := route.Group("dictionaries")

	r.Post("/", h.Create)
	r.Put("/", h.Update)
	r.Delete("/", h.Delete)
	r.Get("/", h.QueryPage)
	r.Get("/code/:code", h.QueryByCode)

}
