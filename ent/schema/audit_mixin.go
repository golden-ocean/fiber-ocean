package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/rs/xid"
)

const (
	FieldCreatedBy = "created_by"
	FieldCreatedAt = "created_at"
	FieldUpdatedBy = "updated_by"
	FieldUpdatedAt = "updated_at"
)

// -------------------------------------------------
// Mixin definition

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type AuditMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Immutable().Comment("主键").Unique().MaxLen(20).Comment("主键"),
		field.Int64(FieldCreatedAt).Immutable().DefaultFunc(time.Now().Unix).Comment("创建时间"),
		field.Int64(FieldUpdatedAt).DefaultFunc(time.Now().Unix).Comment("更新时间"),
		field.String(FieldCreatedBy).Immutable().Optional().Comment("创建人员"),
		field.String(FieldUpdatedBy).Optional().Comment("更新人员"),
	}
}

func (AuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		IDHook,
		AuditHook,
	}
}

func IDHook(next ent.Mutator) ent.Mutator {
	type IDSetter interface {
		SetID(string)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(IDSetter)
		if !ok {
			return nil, fmt.Errorf("unexpected mutation %T", m)
		}
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			ml.SetID(xid.New().String())
		}
		return next.Mutate(ctx, m)
	})
}

func AuditHook(next ent.Mutator) ent.Mutator {
	type AuditLogger interface {
		SetCreatedBy(string)
		CreatedBy() (id string, exists bool)
		SetUpdatedBy(string)
		UpdatedBy() (id string, exists bool)
		// SetDeletedBy(string)
		// DeletedBy() (id string, exists bool)
		// SetDeletedAt(time.Time)
		// DeletedAt() (value time.Time, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}
		ID := "123"
		// usr, err := viewer.UserFromContext(ctx)
		// if err != nil {
		// 	return nil, err
		// }
		switch op := m.Op(); {
		case op.Is(ent.OpCreate):
			if _, exists := ml.CreatedBy(); !exists {
				ml.SetCreatedBy(ID)
			}
		case op.Is(ent.OpUpdateOne | ent.OpUpdate):
			if _, exists := ml.UpdatedBy(); !exists {
				ml.SetUpdatedBy(ID)
			}
		}
		return next.Mutate(ctx, m)
	})
}
