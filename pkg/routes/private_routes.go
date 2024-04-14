package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/app/system/dictionary"
	"github.com/golden-ocean/fiber-ocean/app/system/dictionary_item"
	"github.com/golden-ocean/fiber-ocean/app/system/menu"
	"github.com/golden-ocean/fiber-ocean/app/system/organization"
	"github.com/golden-ocean/fiber-ocean/app/system/position"
	"github.com/golden-ocean/fiber-ocean/app/system/role"
	"github.com/golden-ocean/fiber-ocean/app/system/staff"
)

func PrivateRoutes(a *fiber.App) {
	appRoute := a.Group("/")
	system := appRoute.Group("/system")
	// system := appRoute.Group("/system", middlewares.JWTProtected(), middlewares.CasbinProtected())
	// system := appRoute.Group("/system", middlewares.JWTProtected())

	dictionary.InitRoutes(system)
	dictionary_item.InitRoutes(system)
	position.InitRoutes(system)
	organization.InitRoutes(system)
	role.InitRoutes(system)
	menu.InitRoutes(system)
	staff.InitRoutes(system)
}
