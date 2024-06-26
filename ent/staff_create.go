// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/organization"
	"github.com/golden-ocean/fiber-ocean/ent/staff"
	"github.com/golden-ocean/fiber-ocean/ent/staff_position"
	"github.com/golden-ocean/fiber-ocean/ent/staff_role"
)

// StaffCreate is the builder for creating a Staff entity.
type StaffCreate struct {
	config
	mutation *StaffMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *StaffCreate) SetCreatedAt(i int64) *StaffCreate {
	sc.mutation.SetCreatedAt(i)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StaffCreate) SetNillableCreatedAt(i *int64) *StaffCreate {
	if i != nil {
		sc.SetCreatedAt(*i)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StaffCreate) SetUpdatedAt(i int64) *StaffCreate {
	sc.mutation.SetUpdatedAt(i)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StaffCreate) SetNillableUpdatedAt(i *int64) *StaffCreate {
	if i != nil {
		sc.SetUpdatedAt(*i)
	}
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *StaffCreate) SetCreatedBy(s string) *StaffCreate {
	sc.mutation.SetCreatedBy(s)
	return sc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sc *StaffCreate) SetNillableCreatedBy(s *string) *StaffCreate {
	if s != nil {
		sc.SetCreatedBy(*s)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *StaffCreate) SetUpdatedBy(s string) *StaffCreate {
	sc.mutation.SetUpdatedBy(s)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *StaffCreate) SetNillableUpdatedBy(s *string) *StaffCreate {
	if s != nil {
		sc.SetUpdatedBy(*s)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *StaffCreate) SetStatus(s string) *StaffCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *StaffCreate) SetNillableStatus(s *string) *StaffCreate {
	if s != nil {
		sc.SetStatus(*s)
	}
	return sc
}

// SetSort sets the "sort" field.
func (sc *StaffCreate) SetSort(i int32) *StaffCreate {
	sc.mutation.SetSort(i)
	return sc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (sc *StaffCreate) SetNillableSort(i *int32) *StaffCreate {
	if i != nil {
		sc.SetSort(*i)
	}
	return sc
}

// SetRemark sets the "remark" field.
func (sc *StaffCreate) SetRemark(s string) *StaffCreate {
	sc.mutation.SetRemark(s)
	return sc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sc *StaffCreate) SetNillableRemark(s *string) *StaffCreate {
	if s != nil {
		sc.SetRemark(*s)
	}
	return sc
}

// SetUsername sets the "username" field.
func (sc *StaffCreate) SetUsername(s string) *StaffCreate {
	sc.mutation.SetUsername(s)
	return sc
}

// SetPassword sets the "password" field.
func (sc *StaffCreate) SetPassword(s string) *StaffCreate {
	sc.mutation.SetPassword(s)
	return sc
}

// SetName sets the "name" field.
func (sc *StaffCreate) SetName(s string) *StaffCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetGender sets the "gender" field.
func (sc *StaffCreate) SetGender(s string) *StaffCreate {
	sc.mutation.SetGender(s)
	return sc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (sc *StaffCreate) SetNillableGender(s *string) *StaffCreate {
	if s != nil {
		sc.SetGender(*s)
	}
	return sc
}

// SetWorkStatus sets the "work_status" field.
func (sc *StaffCreate) SetWorkStatus(s string) *StaffCreate {
	sc.mutation.SetWorkStatus(s)
	return sc
}

// SetNillableWorkStatus sets the "work_status" field if the given value is not nil.
func (sc *StaffCreate) SetNillableWorkStatus(s *string) *StaffCreate {
	if s != nil {
		sc.SetWorkStatus(*s)
	}
	return sc
}

// SetMobile sets the "mobile" field.
func (sc *StaffCreate) SetMobile(s string) *StaffCreate {
	sc.mutation.SetMobile(s)
	return sc
}

// SetEmail sets the "email" field.
func (sc *StaffCreate) SetEmail(s string) *StaffCreate {
	sc.mutation.SetEmail(s)
	return sc
}

// SetAvatar sets the "avatar" field.
func (sc *StaffCreate) SetAvatar(s string) *StaffCreate {
	sc.mutation.SetAvatar(s)
	return sc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (sc *StaffCreate) SetNillableAvatar(s *string) *StaffCreate {
	if s != nil {
		sc.SetAvatar(*s)
	}
	return sc
}

// SetOrganizationID sets the "organization_id" field.
func (sc *StaffCreate) SetOrganizationID(s string) *StaffCreate {
	sc.mutation.SetOrganizationID(s)
	return sc
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (sc *StaffCreate) SetNillableOrganizationID(s *string) *StaffCreate {
	if s != nil {
		sc.SetOrganizationID(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StaffCreate) SetID(s string) *StaffCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (sc *StaffCreate) SetOrganization(o *Organization) *StaffCreate {
	return sc.SetOrganizationID(o.ID)
}

// AddStaffsRoleIDs adds the "staffs_roles" edge to the Staff_Role entity by IDs.
func (sc *StaffCreate) AddStaffsRoleIDs(ids ...string) *StaffCreate {
	sc.mutation.AddStaffsRoleIDs(ids...)
	return sc
}

// AddStaffsRoles adds the "staffs_roles" edges to the Staff_Role entity.
func (sc *StaffCreate) AddStaffsRoles(s ...*Staff_Role) *StaffCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddStaffsRoleIDs(ids...)
}

// AddStaffsPositionIDs adds the "staffs_positions" edge to the Staff_Position entity by IDs.
func (sc *StaffCreate) AddStaffsPositionIDs(ids ...string) *StaffCreate {
	sc.mutation.AddStaffsPositionIDs(ids...)
	return sc
}

// AddStaffsPositions adds the "staffs_positions" edges to the Staff_Position entity.
func (sc *StaffCreate) AddStaffsPositions(s ...*Staff_Position) *StaffCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddStaffsPositionIDs(ids...)
}

// Mutation returns the StaffMutation object of the builder.
func (sc *StaffCreate) Mutation() *StaffMutation {
	return sc.mutation
}

// Save creates the Staff in the database.
func (sc *StaffCreate) Save(ctx context.Context) (*Staff, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StaffCreate) SaveX(ctx context.Context) *Staff {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StaffCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StaffCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StaffCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if staff.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized staff.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := staff.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if staff.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized staff.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := staff.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := staff.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.Sort(); !ok {
		v := staff.DefaultSort
		sc.mutation.SetSort(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *StaffCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Staff.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Staff.updated_at"`)}
	}
	if _, ok := sc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "Staff.username"`)}
	}
	if v, ok := sc.mutation.Username(); ok {
		if err := staff.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Staff.username": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Staff.password"`)}
	}
	if v, ok := sc.mutation.Password(); ok {
		if err := staff.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Staff.password": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Staff.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := staff.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Staff.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "Staff.mobile"`)}
	}
	if v, ok := sc.mutation.Mobile(); ok {
		if err := staff.MobileValidator(v); err != nil {
			return &ValidationError{Name: "mobile", err: fmt.Errorf(`ent: validator failed for field "Staff.mobile": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Staff.email"`)}
	}
	if v, ok := sc.mutation.Email(); ok {
		if err := staff.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Staff.email": %w`, err)}
		}
	}
	if v, ok := sc.mutation.ID(); ok {
		if err := staff.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Staff.id": %w`, err)}
		}
	}
	return nil
}

func (sc *StaffCreate) sqlSave(ctx context.Context) (*Staff, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Staff.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StaffCreate) createSpec() (*Staff, *sqlgraph.CreateSpec) {
	var (
		_node = &Staff{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(staff.Table, sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(staff.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(staff.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.SetField(staff.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.SetField(staff.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(staff.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Sort(); ok {
		_spec.SetField(staff.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := sc.mutation.Remark(); ok {
		_spec.SetField(staff.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := sc.mutation.Username(); ok {
		_spec.SetField(staff.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := sc.mutation.Password(); ok {
		_spec.SetField(staff.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(staff.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Gender(); ok {
		_spec.SetField(staff.FieldGender, field.TypeString, value)
		_node.Gender = value
	}
	if value, ok := sc.mutation.WorkStatus(); ok {
		_spec.SetField(staff.FieldWorkStatus, field.TypeString, value)
		_node.WorkStatus = value
	}
	if value, ok := sc.mutation.Mobile(); ok {
		_spec.SetField(staff.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := sc.mutation.Email(); ok {
		_spec.SetField(staff.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := sc.mutation.Avatar(); ok {
		_spec.SetField(staff.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if nodes := sc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   staff.OrganizationTable,
			Columns: []string{staff.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StaffsRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.StaffsRolesTable,
			Columns: []string{staff.StaffsRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StaffsPositionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   staff.StaffsPositionsTable,
			Columns: []string{staff.StaffsPositionsColumn},
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

// StaffCreateBulk is the builder for creating many Staff entities in bulk.
type StaffCreateBulk struct {
	config
	err      error
	builders []*StaffCreate
}

// Save creates the Staff entities in the database.
func (scb *StaffCreateBulk) Save(ctx context.Context) ([]*Staff, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Staff, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StaffMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StaffCreateBulk) SaveX(ctx context.Context) []*Staff {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StaffCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StaffCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
