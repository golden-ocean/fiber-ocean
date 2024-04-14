// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/position"
	"github.com/golden-ocean/fiber-ocean/ent/staff_position"
)

// PositionCreate is the builder for creating a Position entity.
type PositionCreate struct {
	config
	mutation *PositionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *PositionCreate) SetCreatedAt(i int64) *PositionCreate {
	pc.mutation.SetCreatedAt(i)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PositionCreate) SetNillableCreatedAt(i *int64) *PositionCreate {
	if i != nil {
		pc.SetCreatedAt(*i)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PositionCreate) SetUpdatedAt(i int64) *PositionCreate {
	pc.mutation.SetUpdatedAt(i)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PositionCreate) SetNillableUpdatedAt(i *int64) *PositionCreate {
	if i != nil {
		pc.SetUpdatedAt(*i)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *PositionCreate) SetCreatedBy(s string) *PositionCreate {
	pc.mutation.SetCreatedBy(s)
	return pc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pc *PositionCreate) SetNillableCreatedBy(s *string) *PositionCreate {
	if s != nil {
		pc.SetCreatedBy(*s)
	}
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *PositionCreate) SetUpdatedBy(s string) *PositionCreate {
	pc.mutation.SetUpdatedBy(s)
	return pc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pc *PositionCreate) SetNillableUpdatedBy(s *string) *PositionCreate {
	if s != nil {
		pc.SetUpdatedBy(*s)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *PositionCreate) SetStatus(s string) *PositionCreate {
	pc.mutation.SetStatus(s)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PositionCreate) SetNillableStatus(s *string) *PositionCreate {
	if s != nil {
		pc.SetStatus(*s)
	}
	return pc
}

// SetSort sets the "sort" field.
func (pc *PositionCreate) SetSort(i int32) *PositionCreate {
	pc.mutation.SetSort(i)
	return pc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (pc *PositionCreate) SetNillableSort(i *int32) *PositionCreate {
	if i != nil {
		pc.SetSort(*i)
	}
	return pc
}

// SetRemark sets the "remark" field.
func (pc *PositionCreate) SetRemark(s string) *PositionCreate {
	pc.mutation.SetRemark(s)
	return pc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pc *PositionCreate) SetNillableRemark(s *string) *PositionCreate {
	if s != nil {
		pc.SetRemark(*s)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PositionCreate) SetName(s string) *PositionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetCode sets the "code" field.
func (pc *PositionCreate) SetCode(s string) *PositionCreate {
	pc.mutation.SetCode(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PositionCreate) SetID(s string) *PositionCreate {
	pc.mutation.SetID(s)
	return pc
}

// AddStaffIDs adds the "staffs" edge to the Staff_Position entity by IDs.
func (pc *PositionCreate) AddStaffIDs(ids ...string) *PositionCreate {
	pc.mutation.AddStaffIDs(ids...)
	return pc
}

// AddStaffs adds the "staffs" edges to the Staff_Position entity.
func (pc *PositionCreate) AddStaffs(s ...*Staff_Position) *PositionCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddStaffIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pc *PositionCreate) Mutation() *PositionMutation {
	return pc.mutation
}

// Save creates the Position in the database.
func (pc *PositionCreate) Save(ctx context.Context) (*Position, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PositionCreate) SaveX(ctx context.Context) *Position {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PositionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PositionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PositionCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if position.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized position.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := position.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if position.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized position.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := position.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := position.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.Sort(); !ok {
		v := position.DefaultSort
		pc.mutation.SetSort(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PositionCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Position.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Position.updated_at"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Position.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Position.code"`)}
	}
	if v, ok := pc.mutation.Code(); ok {
		if err := position.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Position.code": %w`, err)}
		}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := position.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Position.id": %w`, err)}
		}
	}
	return nil
}

func (pc *PositionCreate) sqlSave(ctx context.Context) (*Position, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Position.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PositionCreate) createSpec() (*Position, *sqlgraph.CreateSpec) {
	var (
		_node = &Position{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(position.Table, sqlgraph.NewFieldSpec(position.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(position.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(position.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(position.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(position.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.Sort(); ok {
		_spec.SetField(position.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := pc.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if nodes := pc.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.StaffsTable,
			Columns: []string{position.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_position.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PositionCreateBulk is the builder for creating many Position entities in bulk.
type PositionCreateBulk struct {
	config
	err      error
	builders []*PositionCreate
}

// Save creates the Position entities in the database.
func (pcb *PositionCreateBulk) Save(ctx context.Context) ([]*Position, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Position, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PositionMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PositionCreateBulk) SaveX(ctx context.Context) []*Position {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PositionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PositionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
