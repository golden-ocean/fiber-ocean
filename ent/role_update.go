// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/predicate"
	"github.com/golden-ocean/fiber-ocean/ent/role"
	"github.com/golden-ocean/fiber-ocean/ent/role_menu"
	"github.com/golden-ocean/fiber-ocean/ent/role_organization"
	"github.com/golden-ocean/fiber-ocean/ent/staff_role"
)

// RoleUpdate is the builder for updating Role entities.
type RoleUpdate struct {
	config
	hooks    []Hook
	mutation *RoleMutation
}

// Where appends a list predicates to the RoleUpdate builder.
func (ru *RoleUpdate) Where(ps ...predicate.Role) *RoleUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RoleUpdate) SetUpdatedAt(i int64) *RoleUpdate {
	ru.mutation.ResetUpdatedAt()
	ru.mutation.SetUpdatedAt(i)
	return ru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableUpdatedAt(i *int64) *RoleUpdate {
	if i != nil {
		ru.SetUpdatedAt(*i)
	}
	return ru
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ru *RoleUpdate) AddUpdatedAt(i int64) *RoleUpdate {
	ru.mutation.AddUpdatedAt(i)
	return ru
}

// SetUpdatedBy sets the "updated_by" field.
func (ru *RoleUpdate) SetUpdatedBy(s string) *RoleUpdate {
	ru.mutation.SetUpdatedBy(s)
	return ru
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableUpdatedBy(s *string) *RoleUpdate {
	if s != nil {
		ru.SetUpdatedBy(*s)
	}
	return ru
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ru *RoleUpdate) ClearUpdatedBy() *RoleUpdate {
	ru.mutation.ClearUpdatedBy()
	return ru
}

// SetStatus sets the "status" field.
func (ru *RoleUpdate) SetStatus(s string) *RoleUpdate {
	ru.mutation.SetStatus(s)
	return ru
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableStatus(s *string) *RoleUpdate {
	if s != nil {
		ru.SetStatus(*s)
	}
	return ru
}

// ClearStatus clears the value of the "status" field.
func (ru *RoleUpdate) ClearStatus() *RoleUpdate {
	ru.mutation.ClearStatus()
	return ru
}

// SetSort sets the "sort" field.
func (ru *RoleUpdate) SetSort(i int32) *RoleUpdate {
	ru.mutation.ResetSort()
	ru.mutation.SetSort(i)
	return ru
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableSort(i *int32) *RoleUpdate {
	if i != nil {
		ru.SetSort(*i)
	}
	return ru
}

// AddSort adds i to the "sort" field.
func (ru *RoleUpdate) AddSort(i int32) *RoleUpdate {
	ru.mutation.AddSort(i)
	return ru
}

// ClearSort clears the value of the "sort" field.
func (ru *RoleUpdate) ClearSort() *RoleUpdate {
	ru.mutation.ClearSort()
	return ru
}

// SetRemark sets the "remark" field.
func (ru *RoleUpdate) SetRemark(s string) *RoleUpdate {
	ru.mutation.SetRemark(s)
	return ru
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableRemark(s *string) *RoleUpdate {
	if s != nil {
		ru.SetRemark(*s)
	}
	return ru
}

// ClearRemark clears the value of the "remark" field.
func (ru *RoleUpdate) ClearRemark() *RoleUpdate {
	ru.mutation.ClearRemark()
	return ru
}

// SetName sets the "name" field.
func (ru *RoleUpdate) SetName(s string) *RoleUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableName(s *string) *RoleUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetCode sets the "code" field.
func (ru *RoleUpdate) SetCode(s string) *RoleUpdate {
	ru.mutation.SetCode(s)
	return ru
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ru *RoleUpdate) SetNillableCode(s *string) *RoleUpdate {
	if s != nil {
		ru.SetCode(*s)
	}
	return ru
}

// AddMenuIDs adds the "menus" edge to the Role_Menu entity by IDs.
func (ru *RoleUpdate) AddMenuIDs(ids ...string) *RoleUpdate {
	ru.mutation.AddMenuIDs(ids...)
	return ru
}

// AddMenus adds the "menus" edges to the Role_Menu entity.
func (ru *RoleUpdate) AddMenus(r ...*Role_Menu) *RoleUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddMenuIDs(ids...)
}

// AddOrganizationIDs adds the "organizations" edge to the Role_Organization entity by IDs.
func (ru *RoleUpdate) AddOrganizationIDs(ids ...string) *RoleUpdate {
	ru.mutation.AddOrganizationIDs(ids...)
	return ru
}

// AddOrganizations adds the "organizations" edges to the Role_Organization entity.
func (ru *RoleUpdate) AddOrganizations(r ...*Role_Organization) *RoleUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddOrganizationIDs(ids...)
}

// AddStaffIDs adds the "staffs" edge to the Staff_Role entity by IDs.
func (ru *RoleUpdate) AddStaffIDs(ids ...string) *RoleUpdate {
	ru.mutation.AddStaffIDs(ids...)
	return ru
}

// AddStaffs adds the "staffs" edges to the Staff_Role entity.
func (ru *RoleUpdate) AddStaffs(s ...*Staff_Role) *RoleUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddStaffIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ru *RoleUpdate) Mutation() *RoleMutation {
	return ru.mutation
}

// ClearMenus clears all "menus" edges to the Role_Menu entity.
func (ru *RoleUpdate) ClearMenus() *RoleUpdate {
	ru.mutation.ClearMenus()
	return ru
}

// RemoveMenuIDs removes the "menus" edge to Role_Menu entities by IDs.
func (ru *RoleUpdate) RemoveMenuIDs(ids ...string) *RoleUpdate {
	ru.mutation.RemoveMenuIDs(ids...)
	return ru
}

// RemoveMenus removes "menus" edges to Role_Menu entities.
func (ru *RoleUpdate) RemoveMenus(r ...*Role_Menu) *RoleUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveMenuIDs(ids...)
}

// ClearOrganizations clears all "organizations" edges to the Role_Organization entity.
func (ru *RoleUpdate) ClearOrganizations() *RoleUpdate {
	ru.mutation.ClearOrganizations()
	return ru
}

// RemoveOrganizationIDs removes the "organizations" edge to Role_Organization entities by IDs.
func (ru *RoleUpdate) RemoveOrganizationIDs(ids ...string) *RoleUpdate {
	ru.mutation.RemoveOrganizationIDs(ids...)
	return ru
}

// RemoveOrganizations removes "organizations" edges to Role_Organization entities.
func (ru *RoleUpdate) RemoveOrganizations(r ...*Role_Organization) *RoleUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveOrganizationIDs(ids...)
}

// ClearStaffs clears all "staffs" edges to the Staff_Role entity.
func (ru *RoleUpdate) ClearStaffs() *RoleUpdate {
	ru.mutation.ClearStaffs()
	return ru
}

// RemoveStaffIDs removes the "staffs" edge to Staff_Role entities by IDs.
func (ru *RoleUpdate) RemoveStaffIDs(ids ...string) *RoleUpdate {
	ru.mutation.RemoveStaffIDs(ids...)
	return ru
}

// RemoveStaffs removes "staffs" edges to Staff_Role entities.
func (ru *RoleUpdate) RemoveStaffs(s ...*Staff_Role) *RoleUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveStaffIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoleUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoleUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoleUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RoleUpdate) check() error {
	if v, ok := ru.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Role.name": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Role.code": %w`, err)}
		}
	}
	return nil
}

func (ru *RoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeString))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ru.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(role.FieldUpdatedAt, field.TypeInt64, value)
	}
	if ru.mutation.CreatedByCleared() {
		_spec.ClearField(role.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ru.mutation.UpdatedBy(); ok {
		_spec.SetField(role.FieldUpdatedBy, field.TypeString, value)
	}
	if ru.mutation.UpdatedByCleared() {
		_spec.ClearField(role.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeString, value)
	}
	if ru.mutation.StatusCleared() {
		_spec.ClearField(role.FieldStatus, field.TypeString)
	}
	if value, ok := ru.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeInt32, value)
	}
	if value, ok := ru.mutation.AddedSort(); ok {
		_spec.AddField(role.FieldSort, field.TypeInt32, value)
	}
	if ru.mutation.SortCleared() {
		_spec.ClearField(role.FieldSort, field.TypeInt32)
	}
	if value, ok := ru.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
	}
	if ru.mutation.RemarkCleared() {
		_spec.ClearField(role.FieldRemark, field.TypeString)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if ru.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: []string{role.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_menu.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedMenusIDs(); len(nodes) > 0 && !ru.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: []string{role.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: []string{role.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.OrganizationsTable,
			Columns: []string{role.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedOrganizationsIDs(); len(nodes) > 0 && !ru.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.OrganizationsTable,
			Columns: []string{role.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.OrganizationsTable,
			Columns: []string{role.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StaffsTable,
			Columns: []string{role.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_role.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedStaffsIDs(); len(nodes) > 0 && !ru.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StaffsTable,
			Columns: []string{role.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StaffsTable,
			Columns: []string{role.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoleUpdateOne is the builder for updating a single Role entity.
type RoleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RoleUpdateOne) SetUpdatedAt(i int64) *RoleUpdateOne {
	ruo.mutation.ResetUpdatedAt()
	ruo.mutation.SetUpdatedAt(i)
	return ruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableUpdatedAt(i *int64) *RoleUpdateOne {
	if i != nil {
		ruo.SetUpdatedAt(*i)
	}
	return ruo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ruo *RoleUpdateOne) AddUpdatedAt(i int64) *RoleUpdateOne {
	ruo.mutation.AddUpdatedAt(i)
	return ruo
}

// SetUpdatedBy sets the "updated_by" field.
func (ruo *RoleUpdateOne) SetUpdatedBy(s string) *RoleUpdateOne {
	ruo.mutation.SetUpdatedBy(s)
	return ruo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableUpdatedBy(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetUpdatedBy(*s)
	}
	return ruo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ruo *RoleUpdateOne) ClearUpdatedBy() *RoleUpdateOne {
	ruo.mutation.ClearUpdatedBy()
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RoleUpdateOne) SetStatus(s string) *RoleUpdateOne {
	ruo.mutation.SetStatus(s)
	return ruo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableStatus(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetStatus(*s)
	}
	return ruo
}

// ClearStatus clears the value of the "status" field.
func (ruo *RoleUpdateOne) ClearStatus() *RoleUpdateOne {
	ruo.mutation.ClearStatus()
	return ruo
}

// SetSort sets the "sort" field.
func (ruo *RoleUpdateOne) SetSort(i int32) *RoleUpdateOne {
	ruo.mutation.ResetSort()
	ruo.mutation.SetSort(i)
	return ruo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableSort(i *int32) *RoleUpdateOne {
	if i != nil {
		ruo.SetSort(*i)
	}
	return ruo
}

// AddSort adds i to the "sort" field.
func (ruo *RoleUpdateOne) AddSort(i int32) *RoleUpdateOne {
	ruo.mutation.AddSort(i)
	return ruo
}

// ClearSort clears the value of the "sort" field.
func (ruo *RoleUpdateOne) ClearSort() *RoleUpdateOne {
	ruo.mutation.ClearSort()
	return ruo
}

// SetRemark sets the "remark" field.
func (ruo *RoleUpdateOne) SetRemark(s string) *RoleUpdateOne {
	ruo.mutation.SetRemark(s)
	return ruo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableRemark(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetRemark(*s)
	}
	return ruo
}

// ClearRemark clears the value of the "remark" field.
func (ruo *RoleUpdateOne) ClearRemark() *RoleUpdateOne {
	ruo.mutation.ClearRemark()
	return ruo
}

// SetName sets the "name" field.
func (ruo *RoleUpdateOne) SetName(s string) *RoleUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableName(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetCode sets the "code" field.
func (ruo *RoleUpdateOne) SetCode(s string) *RoleUpdateOne {
	ruo.mutation.SetCode(s)
	return ruo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ruo *RoleUpdateOne) SetNillableCode(s *string) *RoleUpdateOne {
	if s != nil {
		ruo.SetCode(*s)
	}
	return ruo
}

// AddMenuIDs adds the "menus" edge to the Role_Menu entity by IDs.
func (ruo *RoleUpdateOne) AddMenuIDs(ids ...string) *RoleUpdateOne {
	ruo.mutation.AddMenuIDs(ids...)
	return ruo
}

// AddMenus adds the "menus" edges to the Role_Menu entity.
func (ruo *RoleUpdateOne) AddMenus(r ...*Role_Menu) *RoleUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddMenuIDs(ids...)
}

// AddOrganizationIDs adds the "organizations" edge to the Role_Organization entity by IDs.
func (ruo *RoleUpdateOne) AddOrganizationIDs(ids ...string) *RoleUpdateOne {
	ruo.mutation.AddOrganizationIDs(ids...)
	return ruo
}

// AddOrganizations adds the "organizations" edges to the Role_Organization entity.
func (ruo *RoleUpdateOne) AddOrganizations(r ...*Role_Organization) *RoleUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddOrganizationIDs(ids...)
}

// AddStaffIDs adds the "staffs" edge to the Staff_Role entity by IDs.
func (ruo *RoleUpdateOne) AddStaffIDs(ids ...string) *RoleUpdateOne {
	ruo.mutation.AddStaffIDs(ids...)
	return ruo
}

// AddStaffs adds the "staffs" edges to the Staff_Role entity.
func (ruo *RoleUpdateOne) AddStaffs(s ...*Staff_Role) *RoleUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddStaffIDs(ids...)
}

// Mutation returns the RoleMutation object of the builder.
func (ruo *RoleUpdateOne) Mutation() *RoleMutation {
	return ruo.mutation
}

// ClearMenus clears all "menus" edges to the Role_Menu entity.
func (ruo *RoleUpdateOne) ClearMenus() *RoleUpdateOne {
	ruo.mutation.ClearMenus()
	return ruo
}

// RemoveMenuIDs removes the "menus" edge to Role_Menu entities by IDs.
func (ruo *RoleUpdateOne) RemoveMenuIDs(ids ...string) *RoleUpdateOne {
	ruo.mutation.RemoveMenuIDs(ids...)
	return ruo
}

// RemoveMenus removes "menus" edges to Role_Menu entities.
func (ruo *RoleUpdateOne) RemoveMenus(r ...*Role_Menu) *RoleUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveMenuIDs(ids...)
}

// ClearOrganizations clears all "organizations" edges to the Role_Organization entity.
func (ruo *RoleUpdateOne) ClearOrganizations() *RoleUpdateOne {
	ruo.mutation.ClearOrganizations()
	return ruo
}

// RemoveOrganizationIDs removes the "organizations" edge to Role_Organization entities by IDs.
func (ruo *RoleUpdateOne) RemoveOrganizationIDs(ids ...string) *RoleUpdateOne {
	ruo.mutation.RemoveOrganizationIDs(ids...)
	return ruo
}

// RemoveOrganizations removes "organizations" edges to Role_Organization entities.
func (ruo *RoleUpdateOne) RemoveOrganizations(r ...*Role_Organization) *RoleUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveOrganizationIDs(ids...)
}

// ClearStaffs clears all "staffs" edges to the Staff_Role entity.
func (ruo *RoleUpdateOne) ClearStaffs() *RoleUpdateOne {
	ruo.mutation.ClearStaffs()
	return ruo
}

// RemoveStaffIDs removes the "staffs" edge to Staff_Role entities by IDs.
func (ruo *RoleUpdateOne) RemoveStaffIDs(ids ...string) *RoleUpdateOne {
	ruo.mutation.RemoveStaffIDs(ids...)
	return ruo
}

// RemoveStaffs removes "staffs" edges to Staff_Role entities.
func (ruo *RoleUpdateOne) RemoveStaffs(s ...*Staff_Role) *RoleUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveStaffIDs(ids...)
}

// Where appends a list predicates to the RoleUpdate builder.
func (ruo *RoleUpdateOne) Where(ps ...predicate.Role) *RoleUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoleUpdateOne) Select(field string, fields ...string) *RoleUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Role entity.
func (ruo *RoleUpdateOne) Save(ctx context.Context) (*Role, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoleUpdateOne) SaveX(ctx context.Context) *Role {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoleUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoleUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RoleUpdateOne) check() error {
	if v, ok := ruo.mutation.Name(); ok {
		if err := role.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Role.name": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Code(); ok {
		if err := role.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Role.code": %w`, err)}
		}
	}
	return nil
}

func (ruo *RoleUpdateOne) sqlSave(ctx context.Context) (_node *Role, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(role.Table, role.Columns, sqlgraph.NewFieldSpec(role.FieldID, field.TypeString))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Role.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, role.FieldID)
		for _, f := range fields {
			if !role.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != role.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(role.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ruo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(role.FieldUpdatedAt, field.TypeInt64, value)
	}
	if ruo.mutation.CreatedByCleared() {
		_spec.ClearField(role.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ruo.mutation.UpdatedBy(); ok {
		_spec.SetField(role.FieldUpdatedBy, field.TypeString, value)
	}
	if ruo.mutation.UpdatedByCleared() {
		_spec.ClearField(role.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(role.FieldStatus, field.TypeString, value)
	}
	if ruo.mutation.StatusCleared() {
		_spec.ClearField(role.FieldStatus, field.TypeString)
	}
	if value, ok := ruo.mutation.Sort(); ok {
		_spec.SetField(role.FieldSort, field.TypeInt32, value)
	}
	if value, ok := ruo.mutation.AddedSort(); ok {
		_spec.AddField(role.FieldSort, field.TypeInt32, value)
	}
	if ruo.mutation.SortCleared() {
		_spec.ClearField(role.FieldSort, field.TypeInt32)
	}
	if value, ok := ruo.mutation.Remark(); ok {
		_spec.SetField(role.FieldRemark, field.TypeString, value)
	}
	if ruo.mutation.RemarkCleared() {
		_spec.ClearField(role.FieldRemark, field.TypeString)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(role.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Code(); ok {
		_spec.SetField(role.FieldCode, field.TypeString, value)
	}
	if ruo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: []string{role.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_menu.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedMenusIDs(); len(nodes) > 0 && !ruo.mutation.MenusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: []string{role.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.MenusTable,
			Columns: []string{role.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.OrganizationsTable,
			Columns: []string{role.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedOrganizationsIDs(); len(nodes) > 0 && !ruo.mutation.OrganizationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.OrganizationsTable,
			Columns: []string{role.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.OrganizationsTable,
			Columns: []string{role.OrganizationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StaffsTable,
			Columns: []string{role.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_role.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedStaffsIDs(); len(nodes) > 0 && !ruo.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StaffsTable,
			Columns: []string{role.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.StaffsTable,
			Columns: []string{role.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff_role.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Role{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{role.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}