package staff

import (
	"github.com/gofiber/fiber/v3"
)

func InitRoutes(route fiber.Router) {
	h := NewHandler()
	r := route.Group("staffs")

	r.Post("/", h.Create)
	r.Put("/", h.Update)
	r.Delete("/", h.Delete)
	r.Get("/", h.QueryPage)
	// r.Post("/roles", h.AssignRole)
	//r.Get("/info", controller.FindInfo)
	//staff.Get("/info", h.Info)
	//staff.Get("/role/:id", h.FindRole)
	//staff.Post("/assign", h.Assign)
}
