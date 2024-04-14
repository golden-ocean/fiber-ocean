package global

import (
	"context"

	"github.com/golden-ocean/fiber-ocean/ent"
)

var (
	Client *ent.Client
	Ctx    context.Context
)
