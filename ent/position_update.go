// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/position"
	"github.com/golden-ocean/fiber-ocean/ent/predicate"
	"github.com/golden-ocean/fiber-ocean/ent/staff_position"
)

// PositionUpdate is the builder for updating Position entities.
type PositionUpdate struct {
	config
	hooks    []Hook
	mutation *PositionMutation
}

// Where appends a list predicates to the PositionUpdate builder.
func (pu *PositionUpdate) Where(ps ...predicate.Position) *PositionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PositionUpdate) SetUpdatedAt(i int64) *PositionUpdate {
	pu.mutation.ResetUpdatedAt()
	pu.mutation.SetUpdatedAt(i)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableUpdatedAt(i *int64) *PositionUpdate {
	if i != nil {
		pu.SetUpdatedAt(*i)
	}
	return pu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (pu *PositionUpdate) AddUpdatedAt(i int64) *PositionUpdate {
	pu.mutation.AddUpdatedAt(i)
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *PositionUpdate) SetUpdatedBy(s string) *PositionUpdate {
	pu.mutation.SetUpdatedBy(s)
	return pu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableUpdatedBy(s *string) *PositionUpdate {
	if s != nil {
		pu.SetUpdatedBy(*s)
	}
	return pu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pu *PositionUpdate) ClearUpdatedBy() *PositionUpdate {
	pu.mutation.ClearUpdatedBy()
	return pu
}

// SetStatus sets the "status" field.
func (pu *PositionUpdate) SetStatus(s string) *PositionUpdate {
	pu.mutation.SetStatus(s)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableStatus(s *string) *PositionUpdate {
	if s != nil {
		pu.SetStatus(*s)
	}
	return pu
}

// ClearStatus clears the value of the "status" field.
func (pu *PositionUpdate) ClearStatus() *PositionUpdate {
	pu.mutation.ClearStatus()
	return pu
}

// SetSort sets the "sort" field.
func (pu *PositionUpdate) SetSort(i int32) *PositionUpdate {
	pu.mutation.ResetSort()
	pu.mutation.SetSort(i)
	return pu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableSort(i *int32) *PositionUpdate {
	if i != nil {
		pu.SetSort(*i)
	}
	return pu
}

// AddSort adds i to the "sort" field.
func (pu *PositionUpdate) AddSort(i int32) *PositionUpdate {
	pu.mutation.AddSort(i)
	return pu
}

// ClearSort clears the value of the "sort" field.
func (pu *PositionUpdate) ClearSort() *PositionUpdate {
	pu.mutation.ClearSort()
	return pu
}

// SetRemark sets the "remark" field.
func (pu *PositionUpdate) SetRemark(s string) *PositionUpdate {
	pu.mutation.SetRemark(s)
	return pu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableRemark(s *string) *PositionUpdate {
	if s != nil {
		pu.SetRemark(*s)
	}
	return pu
}

// ClearRemark clears the value of the "remark" field.
func (pu *PositionUpdate) ClearRemark() *PositionUpdate {
	pu.mutation.ClearRemark()
	return pu
}

// SetName sets the "name" field.
func (pu *PositionUpdate) SetName(s string) *PositionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableName(s *string) *PositionUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetCode sets the "code" field.
func (pu *PositionUpdate) SetCode(s string) *PositionUpdate {
	pu.mutation.SetCode(s)
	return pu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (pu *PositionUpdate) SetNillableCode(s *string) *PositionUpdate {
	if s != nil {
		pu.SetCode(*s)
	}
	return pu
}

// AddStaffsPositionIDs adds the "staffs_positions" edge to the Staff_Position entity by IDs.
func (pu *PositionUpdate) AddStaffsPositionIDs(ids ...string) *PositionUpdate {
	pu.mutation.AddStaffsPositionIDs(ids...)
	return pu
}

// AddStaffsPositions adds the "staffs_positions" edges to the Staff_Position entity.
func (pu *PositionUpdate) AddStaffsPositions(s ...*Staff_Position) *PositionUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddStaffsPositionIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pu *PositionUpdate) Mutation() *PositionMutation {
	return pu.mutation
}

// ClearStaffsPositions clears all "staffs_positions" edges to the Staff_Position entity.
func (pu *PositionUpdate) ClearStaffsPositions() *PositionUpdate {
	pu.mutation.ClearStaffsPositions()
	return pu
}

// RemoveStaffsPositionIDs removes the "staffs_positions" edge to Staff_Position entities by IDs.
func (pu *PositionUpdate) RemoveStaffsPositionIDs(ids ...string) *PositionUpdate {
	pu.mutation.RemoveStaffsPositionIDs(ids...)
	return pu
}

// RemoveStaffsPositions removes "staffs_positions" edges to Staff_Position entities.
func (pu *PositionUpdate) RemoveStaffsPositions(s ...*Staff_Position) *PositionUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveStaffsPositionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PositionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PositionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PositionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PositionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PositionUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Code(); ok {
		if err := position.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Position.code": %w`, err)}
		}
	}
	return nil
}

func (pu *PositionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(position.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(position.FieldUpdatedAt, field.TypeInt64, value)
	}
	if pu.mutation.CreatedByCleared() {
		_spec.ClearField(position.FieldCreatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.SetField(position.FieldUpdatedBy, field.TypeString, value)
	}
	if pu.mutation.UpdatedByCleared() {
		_spec.ClearField(position.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeString, value)
	}
	if pu.mutation.StatusCleared() {
		_spec.ClearField(position.FieldStatus, field.TypeString)
	}
	if value, ok := pu.mutation.Sort(); ok {
		_spec.SetField(position.FieldSort, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedSort(); ok {
		_spec.AddField(position.FieldSort, field.TypeInt32, value)
	}
	if pu.mutation.SortCleared() {
		_spec.ClearField(position.FieldSort, field.TypeInt32)
	}
	if value, ok := pu.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
	}
	if pu.mutation.RemarkCleared() {
		_spec.ClearField(position.FieldRemark, field.TypeString)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
	}
	if pu.mutation.StaffsPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.StaffsPositionsTable,
			Columns: []string{position.StaffsPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_position.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedStaffsPositionsIDs(); len(nodes) > 0 && !pu.mutation.StaffsPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.StaffsPositionsTable,
			Columns: []string{position.StaffsPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_position.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StaffsPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.StaffsPositionsTable,
			Columns: []string{position.StaffsPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_position.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PositionUpdateOne is the builder for updating a single Position entity.
type PositionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PositionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PositionUpdateOne) SetUpdatedAt(i int64) *PositionUpdateOne {
	puo.mutation.ResetUpdatedAt()
	puo.mutation.SetUpdatedAt(i)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableUpdatedAt(i *int64) *PositionUpdateOne {
	if i != nil {
		puo.SetUpdatedAt(*i)
	}
	return puo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (puo *PositionUpdateOne) AddUpdatedAt(i int64) *PositionUpdateOne {
	puo.mutation.AddUpdatedAt(i)
	return puo
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *PositionUpdateOne) SetUpdatedBy(s string) *PositionUpdateOne {
	puo.mutation.SetUpdatedBy(s)
	return puo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableUpdatedBy(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetUpdatedBy(*s)
	}
	return puo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (puo *PositionUpdateOne) ClearUpdatedBy() *PositionUpdateOne {
	puo.mutation.ClearUpdatedBy()
	return puo
}

// SetStatus sets the "status" field.
func (puo *PositionUpdateOne) SetStatus(s string) *PositionUpdateOne {
	puo.mutation.SetStatus(s)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableStatus(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetStatus(*s)
	}
	return puo
}

// ClearStatus clears the value of the "status" field.
func (puo *PositionUpdateOne) ClearStatus() *PositionUpdateOne {
	puo.mutation.ClearStatus()
	return puo
}

// SetSort sets the "sort" field.
func (puo *PositionUpdateOne) SetSort(i int32) *PositionUpdateOne {
	puo.mutation.ResetSort()
	puo.mutation.SetSort(i)
	return puo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableSort(i *int32) *PositionUpdateOne {
	if i != nil {
		puo.SetSort(*i)
	}
	return puo
}

// AddSort adds i to the "sort" field.
func (puo *PositionUpdateOne) AddSort(i int32) *PositionUpdateOne {
	puo.mutation.AddSort(i)
	return puo
}

// ClearSort clears the value of the "sort" field.
func (puo *PositionUpdateOne) ClearSort() *PositionUpdateOne {
	puo.mutation.ClearSort()
	return puo
}

// SetRemark sets the "remark" field.
func (puo *PositionUpdateOne) SetRemark(s string) *PositionUpdateOne {
	puo.mutation.SetRemark(s)
	return puo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableRemark(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetRemark(*s)
	}
	return puo
}

// ClearRemark clears the value of the "remark" field.
func (puo *PositionUpdateOne) ClearRemark() *PositionUpdateOne {
	puo.mutation.ClearRemark()
	return puo
}

// SetName sets the "name" field.
func (puo *PositionUpdateOne) SetName(s string) *PositionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableName(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetCode sets the "code" field.
func (puo *PositionUpdateOne) SetCode(s string) *PositionUpdateOne {
	puo.mutation.SetCode(s)
	return puo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (puo *PositionUpdateOne) SetNillableCode(s *string) *PositionUpdateOne {
	if s != nil {
		puo.SetCode(*s)
	}
	return puo
}

// AddStaffsPositionIDs adds the "staffs_positions" edge to the Staff_Position entity by IDs.
func (puo *PositionUpdateOne) AddStaffsPositionIDs(ids ...string) *PositionUpdateOne {
	puo.mutation.AddStaffsPositionIDs(ids...)
	return puo
}

// AddStaffsPositions adds the "staffs_positions" edges to the Staff_Position entity.
func (puo *PositionUpdateOne) AddStaffsPositions(s ...*Staff_Position) *PositionUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddStaffsPositionIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (puo *PositionUpdateOne) Mutation() *PositionMutation {
	return puo.mutation
}

// ClearStaffsPositions clears all "staffs_positions" edges to the Staff_Position entity.
func (puo *PositionUpdateOne) ClearStaffsPositions() *PositionUpdateOne {
	puo.mutation.ClearStaffsPositions()
	return puo
}

// RemoveStaffsPositionIDs removes the "staffs_positions" edge to Staff_Position entities by IDs.
func (puo *PositionUpdateOne) RemoveStaffsPositionIDs(ids ...string) *PositionUpdateOne {
	puo.mutation.RemoveStaffsPositionIDs(ids...)
	return puo
}

// RemoveStaffsPositions removes "staffs_positions" edges to Staff_Position entities.
func (puo *PositionUpdateOne) RemoveStaffsPositions(s ...*Staff_Position) *PositionUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveStaffsPositionIDs(ids...)
}

// Where appends a list predicates to the PositionUpdate builder.
func (puo *PositionUpdateOne) Where(ps ...predicate.Position) *PositionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PositionUpdateOne) Select(field string, fields ...string) *PositionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Position entity.
func (puo *PositionUpdateOne) Save(ctx context.Context) (*Position, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PositionUpdateOne) SaveX(ctx context.Context) *Position {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PositionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PositionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PositionUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := position.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Position.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Code(); ok {
		if err := position.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Position.code": %w`, err)}
		}
	}
	return nil
}

func (puo *PositionUpdateOne) sqlSave(ctx context.Context) (_node *Position, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(position.Table, position.Columns, sqlgraph.NewFieldSpec(position.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Position.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, position.FieldID)
		for _, f := range fields {
			if !position.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != position.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(position.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(position.FieldUpdatedAt, field.TypeInt64, value)
	}
	if puo.mutation.CreatedByCleared() {
		_spec.ClearField(position.FieldCreatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.SetField(position.FieldUpdatedBy, field.TypeString, value)
	}
	if puo.mutation.UpdatedByCleared() {
		_spec.ClearField(position.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(position.FieldStatus, field.TypeString, value)
	}
	if puo.mutation.StatusCleared() {
		_spec.ClearField(position.FieldStatus, field.TypeString)
	}
	if value, ok := puo.mutation.Sort(); ok {
		_spec.SetField(position.FieldSort, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedSort(); ok {
		_spec.AddField(position.FieldSort, field.TypeInt32, value)
	}
	if puo.mutation.SortCleared() {
		_spec.ClearField(position.FieldSort, field.TypeInt32)
	}
	if value, ok := puo.mutation.Remark(); ok {
		_spec.SetField(position.FieldRemark, field.TypeString, value)
	}
	if puo.mutation.RemarkCleared() {
		_spec.ClearField(position.FieldRemark, field.TypeString)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(position.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Code(); ok {
		_spec.SetField(position.FieldCode, field.TypeString, value)
	}
	if puo.mutation.StaffsPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.StaffsPositionsTable,
			Columns: []string{position.StaffsPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_position.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedStaffsPositionsIDs(); len(nodes) > 0 && !puo.mutation.StaffsPositionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.StaffsPositionsTable,
			Columns: []string{position.StaffsPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_position.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StaffsPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.StaffsPositionsTable,
			Columns: []string{position.StaffsPositionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_position.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Position{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{position.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
