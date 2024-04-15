package role

import (
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(route fiber.Router) {
	h := NewHandler()
	r := route.Group("roles")

	r.Post("/", h.Create)
	r.Put("/", h.Update)
	r.Delete("/", h.Delete)
	r.Get("/", h.QueryPage)
	r.Get("/all", h.QueryAll)
	r.Get("/menus/:role_id", h.QueryMenus)
	r.Put("/menus", h.GrantMenus)
}
