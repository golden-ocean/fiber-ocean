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

// DictionaryCreate is the builder for creating a Dictionary entity.
type DictionaryCreate struct {
	config
	mutation *DictionaryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DictionaryCreate) SetCreatedAt(i int64) *DictionaryCreate {
	dc.mutation.SetCreatedAt(i)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableCreatedAt(i *int64) *DictionaryCreate {
	if i != nil {
		dc.SetCreatedAt(*i)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DictionaryCreate) SetUpdatedAt(i int64) *DictionaryCreate {
	dc.mutation.SetUpdatedAt(i)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableUpdatedAt(i *int64) *DictionaryCreate {
	if i != nil {
		dc.SetUpdatedAt(*i)
	}
	return dc
}

// SetCreatedBy sets the "created_by" field.
func (dc *DictionaryCreate) SetCreatedBy(s string) *DictionaryCreate {
	dc.mutation.SetCreatedBy(s)
	return dc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableCreatedBy(s *string) *DictionaryCreate {
	if s != nil {
		dc.SetCreatedBy(*s)
	}
	return dc
}

// SetUpdatedBy sets the "updated_by" field.
func (dc *DictionaryCreate) SetUpdatedBy(s string) *DictionaryCreate {
	dc.mutation.SetUpdatedBy(s)
	return dc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableUpdatedBy(s *string) *DictionaryCreate {
	if s != nil {
		dc.SetUpdatedBy(*s)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DictionaryCreate) SetStatus(s string) *DictionaryCreate {
	dc.mutation.SetStatus(s)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableStatus(s *string) *DictionaryCreate {
	if s != nil {
		dc.SetStatus(*s)
	}
	return dc
}

// SetSort sets the "sort" field.
func (dc *DictionaryCreate) SetSort(i int32) *DictionaryCreate {
	dc.mutation.SetSort(i)
	return dc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableSort(i *int32) *DictionaryCreate {
	if i != nil {
		dc.SetSort(*i)
	}
	return dc
}

// SetRemark sets the "remark" field.
func (dc *DictionaryCreate) SetRemark(s string) *DictionaryCreate {
	dc.mutation.SetRemark(s)
	return dc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableRemark(s *string) *DictionaryCreate {
	if s != nil {
		dc.SetRemark(*s)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DictionaryCreate) SetName(s string) *DictionaryCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetCode sets the "code" field.
func (dc *DictionaryCreate) SetCode(s string) *DictionaryCreate {
	dc.mutation.SetCode(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DictionaryCreate) SetID(s string) *DictionaryCreate {
	dc.mutation.SetID(s)
	return dc
}

// AddItemIDs adds the "items" edge to the Dictionary_Item entity by IDs.
func (dc *DictionaryCreate) AddItemIDs(ids ...string) *DictionaryCreate {
	dc.mutation.AddItemIDs(ids...)
	return dc
}

// AddItems adds the "items" edges to the Dictionary_Item entity.
func (dc *DictionaryCreate) AddItems(d ...*Dictionary_Item) *DictionaryCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddItemIDs(ids...)
}

// Mutation returns the DictionaryMutation object of the builder.
func (dc *DictionaryCreate) Mutation() *DictionaryMutation {
	return dc.mutation
}

// Save creates the Dictionary in the database.
func (dc *DictionaryCreate) Save(ctx context.Context) (*Dictionary, error) {
	if err := dc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DictionaryCreate) SaveX(ctx context.Context) *Dictionary {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DictionaryCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DictionaryCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DictionaryCreate) defaults() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		if dictionary.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized dictionary.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := dictionary.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		if dictionary.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized dictionary.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := dictionary.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := dictionary.DefaultStatus
		dc.mutation.SetStatus(v)
	}
	if _, ok := dc.mutation.Sort(); !ok {
		v := dictionary.DefaultSort
		dc.mutation.SetSort(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dc *DictionaryCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Dictionary.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Dictionary.updated_at"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Dictionary.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := dictionary.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Dictionary.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Dictionary.code"`)}
	}
	if v, ok := dc.mutation.Code(); ok {
		if err := dictionary.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Dictionary.code": %w`, err)}
		}
	}
	if v, ok := dc.mutation.ID(); ok {
		if err := dictionary.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Dictionary.id": %w`, err)}
		}
	}
	return nil
}

func (dc *DictionaryCreate) sqlSave(ctx context.Context) (*Dictionary, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Dictionary.ID type: %T", _spec.ID.Value)
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DictionaryCreate) createSpec() (*Dictionary, *sqlgraph.CreateSpec) {
	var (
		_node = &Dictionary{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(dictionary.Table, sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeString))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(dictionary.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionary.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.CreatedBy(); ok {
		_spec.SetField(dictionary.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := dc.mutation.UpdatedBy(); ok {
		_spec.SetField(dictionary.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(dictionary.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := dc.mutation.Sort(); ok {
		_spec.SetField(dictionary.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := dc.mutation.Remark(); ok {
		_spec.SetField(dictionary.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(dictionary.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Code(); ok {
		_spec.SetField(dictionary.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if nodes := dc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dictionary.ItemsTable,
			Columns: []string{dictionary.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary_item.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DictionaryCreateBulk is the builder for creating many Dictionary entities in bulk.
type DictionaryCreateBulk struct {
	config
	err      error
	builders []*DictionaryCreate
}

// Save creates the Dictionary entities in the database.
func (dcb *DictionaryCreateBulk) Save(ctx context.Context) ([]*Dictionary, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dictionary, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DictionaryMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DictionaryCreateBulk) SaveX(ctx context.Context) []*Dictionary {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DictionaryCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DictionaryCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}