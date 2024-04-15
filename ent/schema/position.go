package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Position struct {
	ent.Schema
}

func (Position) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_positions"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().MinLen(2).MaxLen(64).Comment("职位名称"),
		field.String("code").NotEmpty().Unique().MinLen(2).MaxLen(64).Comment("职位编码"),
	}
}

func (Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("staffs_positions", Staff_Position.Type),
	}
}

func (Position) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
		BaseMixin{},
	}
}
