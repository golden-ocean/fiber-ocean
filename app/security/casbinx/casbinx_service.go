package casbinx

import (
	"github.com/golden-ocean/fiber-ocean/ent"
	"github.com/golden-ocean/fiber-ocean/pkg/common/global"
)

// var modelConf = `
// [request_definition]
// r = sub, obj, act

// [policy_definition]
// p = sub, obj, act

// [role_definition]
// g = _, _

// [policy_effect]
// e = some(where (p.eft == allow))

// [matchers]
// m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act || r.sub == "cochjmcvg7l2ka6gsqmg"
// `

type Service struct {
	client *ent.Client
}

func NewService() *Service {
	return &Service{
		client: global.Client,
	}
}

// func (s *Service) Enforcer() *casbin.Enforcer {
// 	a := NewAdapter()
// 	// m, err := model.NewModelFromFile("./pkg/configs/rbac_model.conf")
// 	m, err := model.NewModelFromString(modelConf)
// 	if err != nil {
// 		panic(err)
// 	}
// 	e, err := casbin.NewEnforcer(m, a)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return e
// }
