// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/dictionary"
	"github.com/golden-ocean/fiber-ocean/ent/dictionary_item"
)

// DictionaryItemCreate is the builder for creating a Dictionary_Item entity.
type DictionaryItemCreate struct {
	config
	mutation *DictionaryItemMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dic *DictionaryItemCreate) SetCreatedAt(i int64) *DictionaryItemCreate {
	dic.mutation.SetCreatedAt(i)
	return dic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dic *DictionaryItemCreate) SetNillableCreatedAt(i *int64) *DictionaryItemCreate {
	if i != nil {
		dic.SetCreatedAt(*i)
	}
	return dic
}

// SetUpdatedAt sets the "updated_at" field.
func (dic *DictionaryItemCreate) SetUpdatedAt(i int64) *DictionaryItemCreate {
	dic.mutation.SetUpdatedAt(i)
	return dic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dic *DictionaryItemCreate) SetNillableUpdatedAt(i *int64) *DictionaryItemCreate {
	if i != nil {
		dic.SetUpdatedAt(*i)
	}
	return dic
}

// SetCreatedBy sets the "created_by" field.
func (dic *DictionaryItemCreate) SetCreatedBy(s string) *DictionaryItemCreate {
	dic.mutation.SetCreatedBy(s)
	return dic
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dic *DictionaryItemCreate) SetNillableCreatedBy(s *string) *DictionaryItemCreate {
	if s != nil {
		dic.SetCreatedBy(*s)
	}
	return dic
}

// SetUpdatedBy sets the "updated_by" field.
func (dic *DictionaryItemCreate) SetUpdatedBy(s string) *DictionaryItemCreate {
	dic.mutation.SetUpdatedBy(s)
	return dic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dic *DictionaryItemCreate) SetNillableUpdatedBy(s *string) *DictionaryItemCreate {
	if s != nil {
		dic.SetUpdatedBy(*s)
	}
	return dic
}

// SetStatus sets the "status" field.
func (dic *DictionaryItemCreate) SetStatus(s string) *DictionaryItemCreate {
	dic.mutation.SetStatus(s)
	return dic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dic *DictionaryItemCreate) SetNillableStatus(s *string) *DictionaryItemCreate {
	if s != nil {
		dic.SetStatus(*s)
	}
	return dic
}

// SetSort sets the "sort" field.
func (dic *DictionaryItemCreate) SetSort(i int32) *DictionaryItemCreate {
	dic.mutation.SetSort(i)
	return dic
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (dic *DictionaryItemCreate) SetNillableSort(i *int32) *DictionaryItemCreate {
	if i != nil {
		dic.SetSort(*i)
	}
	return dic
}

// SetRemark sets the "remark" field.
func (dic *DictionaryItemCreate) SetRemark(s string) *DictionaryItemCreate {
	dic.mutation.SetRemark(s)
	return dic
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dic *DictionaryItemCreate) SetNillableRemark(s *string) *DictionaryItemCreate {
	if s != nil {
		dic.SetRemark(*s)
	}
	return dic
}

// SetLabel sets the "label" field.
func (dic *DictionaryItemCreate) SetLabel(s string) *DictionaryItemCreate {
	dic.mutation.SetLabel(s)
	return dic
}

// SetValue sets the "value" field.
func (dic *DictionaryItemCreate) SetValue(s string) *DictionaryItemCreate {
	dic.mutation.SetValue(s)
	return dic
}

// SetColor sets the "color" field.
func (dic *DictionaryItemCreate) SetColor(s string) *DictionaryItemCreate {
	dic.mutation.SetColor(s)
	return dic
}

// SetDictionaryID sets the "dictionary_id" field.
func (dic *DictionaryItemCreate) SetDictionaryID(s string) *DictionaryItemCreate {
	dic.mutation.SetDictionaryID(s)
	return dic
}

// SetID sets the "id" field.
func (dic *DictionaryItemCreate) SetID(s string) *DictionaryItemCreate {
	dic.mutation.SetID(s)
	return dic
}

// SetDictionary sets the "dictionary" edge to the Dictionary entity.
func (dic *DictionaryItemCreate) SetDictionary(d *Dictionary) *DictionaryItemCreate {
	return dic.SetDictionaryID(d.ID)
}

// Mutation returns the DictionaryItemMutation object of the builder.
func (dic *DictionaryItemCreate) Mutation() *DictionaryItemMutation {
	return dic.mutation
}

// Save creates the Dictionary_Item in the database.
func (dic *DictionaryItemCreate) Save(ctx context.Context) (*Dictionary_Item, error) {
	if err := dic.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dic.sqlSave, dic.mutation, dic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dic *DictionaryItemCreate) SaveX(ctx context.Context) *Dictionary_Item {
	v, err := dic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dic *DictionaryItemCreate) Exec(ctx context.Context) error {
	_, err := dic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dic *DictionaryItemCreate) ExecX(ctx context.Context) {
	if err := dic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dic *DictionaryItemCreate) defaults() error {
	if _, ok := dic.mutation.CreatedAt(); !ok {
		if dictionary_item.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized dictionary_item.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := dictionary_item.DefaultCreatedAt()
		dic.mutation.SetCreatedAt(v)
	}
	if _, ok := dic.mutation.UpdatedAt(); !ok {
		if dictionary_item.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized dictionary_item.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := dictionary_item.DefaultUpdatedAt()
		dic.mutation.SetUpdatedAt(v)
	}
	if _, ok := dic.mutation.Status(); !ok {
		v := dictionary_item.DefaultStatus
		dic.mutation.SetStatus(v)
	}
	if _, ok := dic.mutation.Sort(); !ok {
		v := dictionary_item.DefaultSort
		dic.mutation.SetSort(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dic *DictionaryItemCreate) check() error {
	if _, ok := dic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Dictionary_Item.created_at"`)}
	}
	if _, ok := dic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Dictionary_Item.updated_at"`)}
	}
	if _, ok := dic.mutation.Label(); !ok {
		return &ValidationError{Name: "label", err: errors.New(`ent: missing required field "Dictionary_Item.label"`)}
	}
	if v, ok := dic.mutation.Label(); ok {
		if err := dictionary_item.LabelValidator(v); err != nil {
			return &ValidationError{Name: "label", err: fmt.Errorf(`ent: validator failed for field "Dictionary_Item.label": %w`, err)}
		}
	}
	if _, ok := dic.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Dictionary_Item.value"`)}
	}
	if v, ok := dic.mutation.Value(); ok {
		if err := dictionary_item.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Dictionary_Item.value": %w`, err)}
		}
	}
	if _, ok := dic.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "Dictionary_Item.color"`)}
	}
	if v, ok := dic.mutation.Color(); ok {
		if err := dictionary_item.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "Dictionary_Item.color": %w`, err)}
		}
	}
	if _, ok := dic.mutation.DictionaryID(); !ok {
		return &ValidationError{Name: "dictionary_id", err: errors.New(`ent: missing required field "Dictionary_Item.dictionary_id"`)}
	}
	if v, ok := dic.mutation.DictionaryID(); ok {
		if err := dictionary_item.DictionaryIDValidator(v); err != nil {
			return &ValidationError{Name: "dictionary_id", err: fmt.Errorf(`ent: validator failed for field "Dictionary_Item.dictionary_id": %w`, err)}
		}
	}
	if v, ok := dic.mutation.ID(); ok {
		if err := dictionary_item.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Dictionary_Item.id": %w`, err)}
		}
	}
	if _, ok := dic.mutation.DictionaryID(); !ok {
		return &ValidationError{Name: "dictionary", err: errors.New(`ent: missing required edge "Dictionary_Item.dictionary"`)}
	}
	return nil
}

func (dic *DictionaryItemCreate) sqlSave(ctx context.Context) (*Dictionary_Item, error) {
	if err := dic.check(); err != nil {
		return nil, err
	}
	_node, _spec := dic.createSpec()
	if err := sqlgraph.CreateNode(ctx, dic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Dictionary_Item.ID type: %T", _spec.ID.Value)
		}
	}
	dic.mutation.id = &_node.ID
	dic.mutation.done = true
	return _node, nil
}

func (dic *DictionaryItemCreate) createSpec() (*Dictionary_Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Dictionary_Item{config: dic.config}
		_spec = sqlgraph.NewCreateSpec(dictionary_item.Table, sqlgraph.NewFieldSpec(dictionary_item.FieldID, field.TypeString))
	)
	if id, ok := dic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dic.mutation.CreatedAt(); ok {
		_spec.SetField(dictionary_item.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := dic.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionary_item.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := dic.mutation.CreatedBy(); ok {
		_spec.SetField(dictionary_item.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := dic.mutation.UpdatedBy(); ok {
		_spec.SetField(dictionary_item.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := dic.mutation.Status(); ok {
		_spec.SetField(dictionary_item.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := dic.mutation.Sort(); ok {
		_spec.SetField(dictionary_item.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := dic.mutation.Remark(); ok {
		_spec.SetField(dictionary_item.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := dic.mutation.Label(); ok {
		_spec.SetField(dictionary_item.FieldLabel, field.TypeString, value)
		_node.Label = value
	}
	if value, ok := dic.mutation.Value(); ok {
		_spec.SetField(dictionary_item.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := dic.mutation.Color(); ok {
		_spec.SetField(dictionary_item.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if nodes := dic.mutation.DictionaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionary_item.DictionaryTable,
			Columns: []string{dictionary_item.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DictionaryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DictionaryItemCreateBulk is the builder for creating many Dictionary_Item entities in bulk.
type DictionaryItemCreateBulk struct {
	config
	err      error
	builders []*DictionaryItemCreate
}

// Save creates the Dictionary_Item entities in the database.
func (dicb *DictionaryItemCreateBulk) Save(ctx context.Context) ([]*Dictionary_Item, error) {
	if dicb.err != nil {
		return nil, dicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dicb.builders))
	nodes := make([]*Dictionary_Item, len(dicb.builders))
	mutators := make([]Mutator, len(dicb.builders))
	for i := range dicb.builders {
		func(i int, root context.Context) {
			builder := dicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DictionaryItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dicb *DictionaryItemCreateBulk) SaveX(ctx context.Context) []*Dictionary_Item {
	v, err := dicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dicb *DictionaryItemCreateBulk) Exec(ctx context.Context) error {
	_, err := dicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dicb *DictionaryItemCreateBulk) ExecX(ctx context.Context) {
	if err := dicb.Exec(ctx); err != nil {
		panic(err)
	}
}
