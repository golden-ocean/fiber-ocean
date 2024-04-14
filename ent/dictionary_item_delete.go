// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/dictionary_item"
	"github.com/golden-ocean/fiber-ocean/ent/predicate"
)

// DictionaryItemDelete is the builder for deleting a Dictionary_Item entity.
type DictionaryItemDelete struct {
	config
	hooks    []Hook
	mutation *DictionaryItemMutation
}

// Where appends a list predicates to the DictionaryItemDelete builder.
func (did *DictionaryItemDelete) Where(ps ...predicate.Dictionary_Item) *DictionaryItemDelete {
	did.mutation.Where(ps...)
	return did
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (did *DictionaryItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, did.sqlExec, did.mutation, did.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (did *DictionaryItemDelete) ExecX(ctx context.Context) int {
	n, err := did.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (did *DictionaryItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(dictionary_item.Table, sqlgraph.NewFieldSpec(dictionary_item.FieldID, field.TypeString))
	if ps := did.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, did.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	did.mutation.done = true
	return affected, err
}

// DictionaryItemDeleteOne is the builder for deleting a single Dictionary_Item entity.
type DictionaryItemDeleteOne struct {
	did *DictionaryItemDelete
}

// Where appends a list predicates to the DictionaryItemDelete builder.
func (dido *DictionaryItemDeleteOne) Where(ps ...predicate.Dictionary_Item) *DictionaryItemDeleteOne {
	dido.did.mutation.Where(ps...)
	return dido
}

// Exec executes the deletion query.
func (dido *DictionaryItemDeleteOne) Exec(ctx context.Context) error {
	n, err := dido.did.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dictionary_item.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dido *DictionaryItemDeleteOne) ExecX(ctx context.Context) {
	if err := dido.Exec(ctx); err != nil {
		panic(err)
	}
}
