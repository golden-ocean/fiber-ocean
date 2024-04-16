package middlewares

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gofiber/fiber/v3"
	"github.com/golden-ocean/fiber-ocean/app/security/casbinx"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
	"github.com/golden-ocean/fiber-ocean/pkg/middlewares/casbinware"
	"github.com/golden-ocean/fiber-ocean/pkg/utils"
)

func CasbinProtected() func(fiber.Ctx) error {
	config := casbinware.Config{
		Enforcer:     Enforcer(),
		Lookup:       lookup,
		Unauthorized: unauthorized,
		Forbidden:    forbidden,
	}
	// 按路由规则匹配权限
	return casbinware.New(config).RoutePermission()
}

var modelConf = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act || r.sub == "cochjmcvg7l2ka6gsqmg"
`

func Enforcer() *casbin.Enforcer {
	a := casbinx.NewAdapter()
	m, err := model.NewModelFromString(modelConf)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}
	global.Enforcer = e
	return e
}

func lookup(c fiber.Ctx) string {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return err.Error()
	}
	return claims.ID
}

func unauthorized(c fiber.Ctx) error {
	return fiber.NewError(fiber.StatusUnauthorized, "没有权限！")
}

func forbidden(c fiber.Ctx) error {
	return fiber.NewError(fiber.StatusForbidden, "禁止访问！")
}
