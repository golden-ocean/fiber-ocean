package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Staff_Position struct {
	ent.Schema
}

func (Staff_Position) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_staffs_positions"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Staff_Position) Fields() []ent.Field {
	return []ent.Field{
		field.String("staff_id").NotEmpty().Comment("员工ID"),
		field.String("position_id").NotEmpty().Comment("职位ID"),
	}
}

func (Staff_Position) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("staffs", Staff.Type).Field("staff_id").Required().Ref("staffs_positions").Unique(),
		edge.From("positions", Position.Type).Field("position_id").Required().Ref("staffs_positions").Unique(),
	}
}

func (Staff_Position) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
func (Staff_Position) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("staff_id", "position_id").Unique().StorageKey("staff_id_position_id"),
	}
}
