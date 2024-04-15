package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_roles"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().MinLen(2).MaxLen(64).Comment("角色名称"),
		field.String("code").NotEmpty().Unique().MinLen(2).MaxLen(64).Comment("角色编码"),
	}
}

func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles_menus", Role_Menu.Type),
		edge.To("roles_organizations", Role_Organization.Type),
		edge.To("staffs_roles", Staff_Role.Type),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
		BaseMixin{},
	}
}
