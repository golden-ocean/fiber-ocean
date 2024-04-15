package dictionary_item

import (
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(route fiber.Router) {
	h := NewHandler()
	dictionaryRoutes := route.Group("dictionary")
	r := dictionaryRoutes.Group("items")

	r.Get("/", h.QueryPage)
	r.Post("/", h.Create)
	r.Put("/", h.Update)
	r.Delete("/", h.Delete)
}
