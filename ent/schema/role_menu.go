package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Role_Menu struct {
	ent.Schema
}

func (Role_Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_roles_menus"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Role_Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_id").NotEmpty().Comment("角色ID"),
		field.String("menu_id").NotEmpty().Comment("菜单ID"),
	}
}

func (Role_Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Field("role_id").Required().Ref("menus").Unique(),
		edge.From("menus", Menu.Type).Field("menu_id").Required().Ref("roles").Unique(),
	}
}

func (Role_Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
func (Role_Menu) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "menu_id").Unique().StorageKey("role_id_menu_id"),
	}
}
