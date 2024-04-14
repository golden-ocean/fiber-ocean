// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/organization"
	"github.com/golden-ocean/fiber-ocean/ent/predicate"
	"github.com/golden-ocean/fiber-ocean/ent/role"
	"github.com/golden-ocean/fiber-ocean/ent/role_organization"
)

// RoleOrganizationUpdate is the builder for updating Role_Organization entities.
type RoleOrganizationUpdate struct {
	config
	hooks    []Hook
	mutation *RoleOrganizationMutation
}

// Where appends a list predicates to the RoleOrganizationUpdate builder.
func (rou *RoleOrganizationUpdate) Where(ps ...predicate.Role_Organization) *RoleOrganizationUpdate {
	rou.mutation.Where(ps...)
	return rou
}

// SetUpdatedAt sets the "updated_at" field.
func (rou *RoleOrganizationUpdate) SetUpdatedAt(i int64) *RoleOrganizationUpdate {
	rou.mutation.ResetUpdatedAt()
	rou.mutation.SetUpdatedAt(i)
	return rou
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rou *RoleOrganizationUpdate) SetNillableUpdatedAt(i *int64) *RoleOrganizationUpdate {
	if i != nil {
		rou.SetUpdatedAt(*i)
	}
	return rou
}

// AddUpdatedAt adds i to the "updated_at" field.
func (rou *RoleOrganizationUpdate) AddUpdatedAt(i int64) *RoleOrganizationUpdate {
	rou.mutation.AddUpdatedAt(i)
	return rou
}

// SetUpdatedBy sets the "updated_by" field.
func (rou *RoleOrganizationUpdate) SetUpdatedBy(s string) *RoleOrganizationUpdate {
	rou.mutation.SetUpdatedBy(s)
	return rou
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rou *RoleOrganizationUpdate) SetNillableUpdatedBy(s *string) *RoleOrganizationUpdate {
	if s != nil {
		rou.SetUpdatedBy(*s)
	}
	return rou
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (rou *RoleOrganizationUpdate) ClearUpdatedBy() *RoleOrganizationUpdate {
	rou.mutation.ClearUpdatedBy()
	return rou
}

// SetRoleID sets the "role_id" field.
func (rou *RoleOrganizationUpdate) SetRoleID(s string) *RoleOrganizationUpdate {
	rou.mutation.SetRoleID(s)
	return rou
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (rou *RoleOrganizationUpdate) SetNillableRoleID(s *string) *RoleOrganizationUpdate {
	if s != nil {
		rou.SetRoleID(*s)
	}
	return rou
}

// SetOrganizationID sets the "organization_id" field.
func (rou *RoleOrganizationUpdate) SetOrganizationID(s string) *RoleOrganizationUpdate {
	rou.mutation.SetOrganizationID(s)
	return rou
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (rou *RoleOrganizationUpdate) SetNillableOrganizationID(s *string) *RoleOrganizationUpdate {
	if s != nil {
		rou.SetOrganizationID(*s)
	}
	return rou
}

// SetRolesID sets the "roles" edge to the Role entity by ID.
func (rou *RoleOrganizationUpdate) SetRolesID(id string) *RoleOrganizationUpdate {
	rou.mutation.SetRolesID(id)
	return rou
}

// SetRoles sets the "roles" edge to the Role entity.
func (rou *RoleOrganizationUpdate) SetRoles(r *Role) *RoleOrganizationUpdate {
	return rou.SetRolesID(r.ID)
}

// SetOrganizationsID sets the "organizations" edge to the Organization entity by ID.
func (rou *RoleOrganizationUpdate) SetOrganizationsID(id string) *RoleOrganizationUpdate {
	rou.mutation.SetOrganizationsID(id)
	return rou
}

// SetOrganizations sets the "organizations" edge to the Organization entity.
func (rou *RoleOrganizationUpdate) SetOrganizations(o *Organization) *RoleOrganizationUpdate {
	return rou.SetOrganizationsID(o.ID)
}

// Mutation returns the RoleOrganizationMutation object of the builder.
func (rou *RoleOrganizationUpdate) Mutation() *RoleOrganizationMutation {
	return rou.mutation
}

// ClearRoles clears the "roles" edge to the Role entity.
func (rou *RoleOrganizationUpdate) ClearRoles() *RoleOrganizationUpdate {
	rou.mutation.ClearRoles()
	return rou
}

// ClearOrganizations clears the "organizations" edge to the Organization entity.
func (rou *RoleOrganizationUpdate) ClearOrganizations() *RoleOrganizationUpdate {
	rou.mutation.ClearOrganizations()
	return rou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rou *RoleOrganizationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rou.sqlSave, rou.mutation, rou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rou *RoleOrganizationUpdate) SaveX(ctx context.Context) int {
	affected, err := rou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rou *RoleOrganizationUpdate) Exec(ctx context.Context) error {
	_, err := rou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rou *RoleOrganizationUpdate) ExecX(ctx context.Context) {
	if err := rou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rou *RoleOrganizationUpdate) check() error {
	if v, ok := rou.mutation.RoleID(); ok {
		if err := role_organization.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "Role_Organization.role_id": %w`, err)}
		}
	}
	if v, ok := rou.mutation.OrganizationID(); ok {
		if err := role_organization.OrganizationIDValidator(v); err != nil {
			return &ValidationError{Name: "organization_id", err: fmt.Errorf(`ent: validator failed for field "Role_Organization.organization_id": %w`, err)}
		}
	}
	if _, ok := rou.mutation.RolesID(); rou.mutation.RolesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Role_Organization.roles"`)
	}
	if _, ok := rou.mutation.OrganizationsID(); rou.mutation.OrganizationsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Role_Organization.organizations"`)
	}
	return nil
}

func (rou *RoleOrganizationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(role_organization.Table, role_organization.Columns, sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString))
	if ps := rou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rou.mutation.UpdatedAt(); ok {
		_spec.SetField(role_organization.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := rou.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(role_organization.FieldUpdatedAt, field.TypeInt64, value)
	}
	if rou.mutation.CreatedByCleared() {
		_spec.ClearField(role_organization.FieldCreatedBy, field.TypeString)
	}
	if value, ok := rou.mutation.UpdatedBy(); ok {
		_spec.SetField(role_organization.FieldUpdatedBy, field.TypeString, value)
	}
	if rou.mutation.UpdatedByCleared() {
		_spec.ClearField(role_organization.FieldUpdatedBy, field.TypeString)
	}
	if rou.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role_organization.RolesTable,
			Columns: []string{role_organization.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rou.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role_organization.RolesTable,
			Columns: []string{role_organization.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rou.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role_organization.OrganizationsTable,
			Columns: []string{role_organization.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rou.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role_organization.OrganizationsTable,
			Columns: []string{role_organization.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role_organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rou.mutation.done = true
	return n, nil
}

// RoleOrganizationUpdateOne is the builder for updating a single Role_Organization entity.
type RoleOrganizationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleOrganizationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (rouo *RoleOrganizationUpdateOne) SetUpdatedAt(i int64) *RoleOrganizationUpdateOne {
	rouo.mutation.ResetUpdatedAt()
	rouo.mutation.SetUpdatedAt(i)
	return rouo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rouo *RoleOrganizationUpdateOne) SetNillableUpdatedAt(i *int64) *RoleOrganizationUpdateOne {
	if i != nil {
		rouo.SetUpdatedAt(*i)
	}
	return rouo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (rouo *RoleOrganizationUpdateOne) AddUpdatedAt(i int64) *RoleOrganizationUpdateOne {
	rouo.mutation.AddUpdatedAt(i)
	return rouo
}

// SetUpdatedBy sets the "updated_by" field.
func (rouo *RoleOrganizationUpdateOne) SetUpdatedBy(s string) *RoleOrganizationUpdateOne {
	rouo.mutation.SetUpdatedBy(s)
	return rouo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (rouo *RoleOrganizationUpdateOne) SetNillableUpdatedBy(s *string) *RoleOrganizationUpdateOne {
	if s != nil {
		rouo.SetUpdatedBy(*s)
	}
	return rouo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (rouo *RoleOrganizationUpdateOne) ClearUpdatedBy() *RoleOrganizationUpdateOne {
	rouo.mutation.ClearUpdatedBy()
	return rouo
}

// SetRoleID sets the "role_id" field.
func (rouo *RoleOrganizationUpdateOne) SetRoleID(s string) *RoleOrganizationUpdateOne {
	rouo.mutation.SetRoleID(s)
	return rouo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (rouo *RoleOrganizationUpdateOne) SetNillableRoleID(s *string) *RoleOrganizationUpdateOne {
	if s != nil {
		rouo.SetRoleID(*s)
	}
	return rouo
}

// SetOrganizationID sets the "organization_id" field.
func (rouo *RoleOrganizationUpdateOne) SetOrganizationID(s string) *RoleOrganizationUpdateOne {
	rouo.mutation.SetOrganizationID(s)
	return rouo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (rouo *RoleOrganizationUpdateOne) SetNillableOrganizationID(s *string) *RoleOrganizationUpdateOne {
	if s != nil {
		rouo.SetOrganizationID(*s)
	}
	return rouo
}

// SetRolesID sets the "roles" edge to the Role entity by ID.
func (rouo *RoleOrganizationUpdateOne) SetRolesID(id string) *RoleOrganizationUpdateOne {
	rouo.mutation.SetRolesID(id)
	return rouo
}

// SetRoles sets the "roles" edge to the Role entity.
func (rouo *RoleOrganizationUpdateOne) SetRoles(r *Role) *RoleOrganizationUpdateOne {
	return rouo.SetRolesID(r.ID)
}

// SetOrganizationsID sets the "organizations" edge to the Organization entity by ID.
func (rouo *RoleOrganizationUpdateOne) SetOrganizationsID(id string) *RoleOrganizationUpdateOne {
	rouo.mutation.SetOrganizationsID(id)
	return rouo
}

// SetOrganizations sets the "organizations" edge to the Organization entity.
func (rouo *RoleOrganizationUpdateOne) SetOrganizations(o *Organization) *RoleOrganizationUpdateOne {
	return rouo.SetOrganizationsID(o.ID)
}

// Mutation returns the RoleOrganizationMutation object of the builder.
func (rouo *RoleOrganizationUpdateOne) Mutation() *RoleOrganizationMutation {
	return rouo.mutation
}

// ClearRoles clears the "roles" edge to the Role entity.
func (rouo *RoleOrganizationUpdateOne) ClearRoles() *RoleOrganizationUpdateOne {
	rouo.mutation.ClearRoles()
	return rouo
}

// ClearOrganizations clears the "organizations" edge to the Organization entity.
func (rouo *RoleOrganizationUpdateOne) ClearOrganizations() *RoleOrganizationUpdateOne {
	rouo.mutation.ClearOrganizations()
	return rouo
}

// Where appends a list predicates to the RoleOrganizationUpdate builder.
func (rouo *RoleOrganizationUpdateOne) Where(ps ...predicate.Role_Organization) *RoleOrganizationUpdateOne {
	rouo.mutation.Where(ps...)
	return rouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rouo *RoleOrganizationUpdateOne) Select(field string, fields ...string) *RoleOrganizationUpdateOne {
	rouo.fields = append([]string{field}, fields...)
	return rouo
}

// Save executes the query and returns the updated Role_Organization entity.
func (rouo *RoleOrganizationUpdateOne) Save(ctx context.Context) (*Role_Organization, error) {
	return withHooks(ctx, rouo.sqlSave, rouo.mutation, rouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rouo *RoleOrganizationUpdateOne) SaveX(ctx context.Context) *Role_Organization {
	node, err := rouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rouo *RoleOrganizationUpdateOne) Exec(ctx context.Context) error {
	_, err := rouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rouo *RoleOrganizationUpdateOne) ExecX(ctx context.Context) {
	if err := rouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rouo *RoleOrganizationUpdateOne) check() error {
	if v, ok := rouo.mutation.RoleID(); ok {
		if err := role_organization.RoleIDValidator(v); err != nil {
			return &ValidationError{Name: "role_id", err: fmt.Errorf(`ent: validator failed for field "Role_Organization.role_id": %w`, err)}
		}
	}
	if v, ok := rouo.mutation.OrganizationID(); ok {
		if err := role_organization.OrganizationIDValidator(v); err != nil {
			return &ValidationError{Name: "organization_id", err: fmt.Errorf(`ent: validator failed for field "Role_Organization.organization_id": %w`, err)}
		}
	}
	if _, ok := rouo.mutation.RolesID(); rouo.mutation.RolesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Role_Organization.roles"`)
	}
	if _, ok := rouo.mutation.OrganizationsID(); rouo.mutation.OrganizationsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Role_Organization.organizations"`)
	}
	return nil
}

func (rouo *RoleOrganizationUpdateOne) sqlSave(ctx context.Context) (_node *Role_Organization, err error) {
	if err := rouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(role_organization.Table, role_organization.Columns, sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString))
	id, ok := rouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role_Organization.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role_organization.FieldID)
		for _, f := range fields {
			if !role_organization.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role_organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rouo.mutation.UpdatedAt(); ok {
		_spec.SetField(role_organization.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := rouo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(role_organization.FieldUpdatedAt, field.TypeInt64, value)
	}
	if rouo.mutation.CreatedByCleared() {
		_spec.ClearField(role_organization.FieldCreatedBy, field.TypeString)
	}
	if value, ok := rouo.mutation.UpdatedBy(); ok {
		_spec.SetField(role_organization.FieldUpdatedBy, field.TypeString, value)
	}
	if rouo.mutation.UpdatedByCleared() {
		_spec.ClearField(role_organization.FieldUpdatedBy, field.TypeString)
	}
	if rouo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role_organization.RolesTable,
			Columns: []string{role_organization.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rouo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role_organization.RolesTable,
			Columns: []string{role_organization.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rouo.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role_organization.OrganizationsTable,
			Columns: []string{role_organization.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rouo.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role_organization.OrganizationsTable,
			Columns: []string{role_organization.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Role_Organization{config: rouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role_organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rouo.mutation.done = true
	return _node, nil
}