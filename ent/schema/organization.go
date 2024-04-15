package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Organization struct {
	ent.Schema
}

func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_organizations"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MinLen(2).MaxLen(64).Comment("组织名称"),
		field.String("code").NotEmpty().MinLen(2).MaxLen(64).Comment("组织编码"),
		field.String("parent_id").Optional().Comment("父级ID"),
	}
}

func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Organization.Type).
			From("parent").Unique().Field("parent_id"),
		edge.To("roles_organizations", Role_Organization.Type),
		edge.From("staffs", Staff.Type).Ref("organization"),
	}
}

func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
		BaseMixin{},
	}
}

func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("parent_id", "name").Unique().StorageKey("parent_id_name"),
		index.Fields("parent_id", "code").Unique().StorageKey("parent_id_code"),
	}
}
