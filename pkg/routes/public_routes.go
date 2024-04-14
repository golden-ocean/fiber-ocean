package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/app/security/auth"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/")
	auth.InitRoutes(route)
	// staff.InitRoutes(route)
	// auth.InitRoutes(route)
}
