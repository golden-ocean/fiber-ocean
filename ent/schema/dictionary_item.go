package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Dictionary_Item struct {
	ent.Schema
}

func (Dictionary_Item) Annotations() []schema.Annotation {
	return []schema.Annotation{

		entsql.Annotation{Table: "system_dictionary_items"},
		entsql.WithComments(true),
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}

func (Dictionary_Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").NotEmpty().MinLen(2).MaxLen(64).Comment("字典选项标签"),
		field.String("value").NotEmpty().MinLen(2).MaxLen(64).Comment("字典选项值"),
		field.String("color").MinLen(2).MaxLen(64).Comment("颜色"),
		field.String("dictionary_id").StructTag(`json:"-"`).NotEmpty().Comment("字典ID"),
	}
}

func (Dictionary_Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("dictionary", Dictionary.Type).
			Ref("items").
			Field("dictionary_id").
			Required().
			Unique(),
	}
}

func (Dictionary_Item) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
		BaseMixin{},
	}
}

func (Dictionary_Item) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("dictionary_id", "label").Unique().StorageKey("dictionary_id_label"),
		index.Fields("dictionary_id", "value").Unique().StorageKey("dictionary_id_value"),
	}
}
