package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Menu struct {
	ent.Schema
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_menus"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().MinLen(2).MaxLen(64).Comment("权限名称"),
		field.String("parent_id").Optional().Comment("父级ID").StructTag(`json:"parent_id,omitempty"`),
		field.String("icon").Optional().Comment("图标"),
		field.String("path").Optional().Comment("路径"),
		field.String("permission").Optional().Comment("权限"),
		field.String("component").Optional().Comment("组件"),
		field.String("type").Optional().Comment("菜单类型"),
		field.String("method").Optional().Comment("请求方法"),
		field.String("visible").Optional().Default("true").Comment("是否显示"),
	}
}

func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Menu.Type).
			From("parent").Unique().Field("parent_id"),
		edge.To("roles_menus", Role_Menu.Type),
	}
}

func (Menu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
		BaseMixin{},
	}
}
