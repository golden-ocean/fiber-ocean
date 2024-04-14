package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Staff_Role struct {
	ent.Schema
}

func (Staff_Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_staffs_roles"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Staff_Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("staff_id").NotEmpty().Comment("员工ID"),
		field.String("role_id").NotEmpty().Comment("角色ID"),
	}
}

func (Staff_Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("staffs", Staff.Type).Field("staff_id").Required().Ref("roles").Unique(),
		edge.From("roles", Role.Type).Field("role_id").Required().Ref("staffs").Unique(),
	}
}

func (Staff_Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
func (Staff_Role) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("staff_id", "role_id").Unique().StorageKey("staff_id_role_id"),
	}
}
