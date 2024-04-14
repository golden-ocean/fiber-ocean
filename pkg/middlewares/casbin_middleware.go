package middlewares

// import (
// 	casbinMiddleware "github.com/gofiber/contrib/casbin"
// 	"github.com/gofiber/fiber/v2"
// )

// func CasbinProtected() func(fiber.Ctx) error {
// 	config := casbinMiddleware.Config{
// 		ModelFilePath: "./pkg/configs/rbac_model.conf",
// 		PolicyAdapter: casbin.NewAdapter(),
// 		//Enforcer:      enforcer(),
// 		Lookup:       lookup,
// 		Unauthorized: unauthorized,
// 		Forbidden:    forbidden,
// 	}
// 	// 按路由规则匹配权限
// 	return casbinMiddleware.New(config).RoutePermission()
// }

// //func enforcer() *casbin.Enforcer {
// //	a := casbinAdapter()
// //	m, err := model.NewModelFromString(modelConf)
// //	if err != nil {
// //		panic(err)
// //	}
// //	newEnforcer, err := casbin.NewEnforcer(m, a)
// //	if err != nil {
// //		panic(err)
// //	}
// //	return newEnforcer
// //}

// func lookup(c *fiber.Ctx) string {
// 	return c.Locals("id").(string)
// }

// func unauthorized(c *fiber.Ctx) error {
// 	return fiber.NewError(fiber.StatusUnauthorized, "没有权限！")
// }

// func forbidden(c *fiber.Ctx) error {
// 	return fiber.NewError(fiber.StatusForbidden, "禁止访问！")
// }
