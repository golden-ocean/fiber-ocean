package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Role_Organization struct {
	ent.Schema
}

func (Role_Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_roles_organizations"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Role_Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_id").NotEmpty().Comment("角色ID"),
		field.String("organization_id").NotEmpty().Comment("组织ID"),
	}
}

func (Role_Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Field("role_id").Required().Ref("roles_organizations").Unique(),
		edge.From("organizations", Organization.Type).Field("organization_id").Required().Ref("roles_organizations").Unique(),
	}
}

func (Role_Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
func (Role_Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "organization_id").Unique().StorageKey("role_id_organization_id"),
	}
}
