package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Dictionary struct {
	ent.Schema
}

func (Dictionary) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_dictionaries"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}
func (Dictionary) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique().MinLen(2).MaxLen(64).Comment("字典名称"),
		field.String("code").NotEmpty().Unique().MinLen(2).MaxLen(64).Comment("字典编码"),
	}
}

func (Dictionary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", Dictionary_Item.Type),
		// edge.To("items", DictionaryItem.Type).Annotations(entsql.OnDelete(entsql.Restrict)),
	}
}

func (Dictionary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
		BaseMixin{},
	}
}
