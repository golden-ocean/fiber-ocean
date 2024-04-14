package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Staff struct {
	ent.Schema
}

func (Staff) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "system_staffs"},
		entsql.WithComments(true),
	}
}

func (Staff) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique().Comment("用户名"),
		field.String("password").NotEmpty().Sensitive().Comment("密码"),
		field.String("name").NotEmpty().Comment("姓名"),
		field.String("gender").Optional().Comment("性别"),
		field.String("work_status").Optional().Comment("工作状态"),
		field.String("mobile").NotEmpty().Unique().Comment("手机号码"),
		field.String("email").NotEmpty().Unique().Comment("邮箱"),
		field.String("avatar").Optional().Comment("头像"),
		field.String("organization_id").Optional().Comment("组织ID"),
	}
}

func (Staff) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organization", Organization.Type).
			Field("organization_id").
			Unique(),
		edge.To("roles", Staff_Role.Type),
		edge.To("positions", Staff_Position.Type),
	}
}

func (Staff) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
		BaseMixin{},
	}
}
