package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/golden-ocean/fiber-ocean/pkg/common/constants"
)

// -------------------------------------------------
// Mixin definition

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type BaseMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("status").Optional().Default(constants.ENABLE).Comment("状态: Enable/Disable"),
		field.Int32("sort").Optional().Default(1000).Comment("排序"),
		field.String("remark").Optional().Comment("备注"),
	}
}
