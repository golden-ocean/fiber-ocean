package global

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/golden-ocean/fiber-ocean/ent"
)

var (
	Client   *ent.Client
	Enforcer *casbin.Enforcer
	Ctx      context.Context
)
