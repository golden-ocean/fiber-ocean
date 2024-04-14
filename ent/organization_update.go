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
	"github.com/golden-ocean/fiber-ocean/ent/role_organization"
	"github.com/golden-ocean/fiber-ocean/ent/staff"
)

// OrganizationUpdate is the builder for updating Organization entities.
type OrganizationUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationMutation
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ou *OrganizationUpdate) Where(ps ...predicate.Organization) *OrganizationUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrganizationUpdate) SetUpdatedAt(i int64) *OrganizationUpdate {
	ou.mutation.ResetUpdatedAt()
	ou.mutation.SetUpdatedAt(i)
	return ou
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableUpdatedAt(i *int64) *OrganizationUpdate {
	if i != nil {
		ou.SetUpdatedAt(*i)
	}
	return ou
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ou *OrganizationUpdate) AddUpdatedAt(i int64) *OrganizationUpdate {
	ou.mutation.AddUpdatedAt(i)
	return ou
}

// SetUpdatedBy sets the "updated_by" field.
func (ou *OrganizationUpdate) SetUpdatedBy(s string) *OrganizationUpdate {
	ou.mutation.SetUpdatedBy(s)
	return ou
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableUpdatedBy(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetUpdatedBy(*s)
	}
	return ou
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ou *OrganizationUpdate) ClearUpdatedBy() *OrganizationUpdate {
	ou.mutation.ClearUpdatedBy()
	return ou
}

// SetStatus sets the "status" field.
func (ou *OrganizationUpdate) SetStatus(s string) *OrganizationUpdate {
	ou.mutation.SetStatus(s)
	return ou
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableStatus(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetStatus(*s)
	}
	return ou
}

// ClearStatus clears the value of the "status" field.
func (ou *OrganizationUpdate) ClearStatus() *OrganizationUpdate {
	ou.mutation.ClearStatus()
	return ou
}

// SetSort sets the "sort" field.
func (ou *OrganizationUpdate) SetSort(i int32) *OrganizationUpdate {
	ou.mutation.ResetSort()
	ou.mutation.SetSort(i)
	return ou
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableSort(i *int32) *OrganizationUpdate {
	if i != nil {
		ou.SetSort(*i)
	}
	return ou
}

// AddSort adds i to the "sort" field.
func (ou *OrganizationUpdate) AddSort(i int32) *OrganizationUpdate {
	ou.mutation.AddSort(i)
	return ou
}

// ClearSort clears the value of the "sort" field.
func (ou *OrganizationUpdate) ClearSort() *OrganizationUpdate {
	ou.mutation.ClearSort()
	return ou
}

// SetRemark sets the "remark" field.
func (ou *OrganizationUpdate) SetRemark(s string) *OrganizationUpdate {
	ou.mutation.SetRemark(s)
	return ou
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableRemark(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetRemark(*s)
	}
	return ou
}

// ClearRemark clears the value of the "remark" field.
func (ou *OrganizationUpdate) ClearRemark() *OrganizationUpdate {
	ou.mutation.ClearRemark()
	return ou
}

// SetName sets the "name" field.
func (ou *OrganizationUpdate) SetName(s string) *OrganizationUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableName(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetName(*s)
	}
	return ou
}

// SetCode sets the "code" field.
func (ou *OrganizationUpdate) SetCode(s string) *OrganizationUpdate {
	ou.mutation.SetCode(s)
	return ou
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableCode(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetCode(*s)
	}
	return ou
}

// SetParentID sets the "parent_id" field.
func (ou *OrganizationUpdate) SetParentID(s string) *OrganizationUpdate {
	ou.mutation.SetParentID(s)
	return ou
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableParentID(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetParentID(*s)
	}
	return ou
}

// ClearParentID clears the value of the "parent_id" field.
func (ou *OrganizationUpdate) ClearParentID() *OrganizationUpdate {
	ou.mutation.ClearParentID()
	return ou
}

// SetParent sets the "parent" edge to the Organization entity.
func (ou *OrganizationUpdate) SetParent(o *Organization) *OrganizationUpdate {
	return ou.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the Organization entity by IDs.
func (ou *OrganizationUpdate) AddChildIDs(ids ...string) *OrganizationUpdate {
	ou.mutation.AddChildIDs(ids...)
	return ou
}

// AddChildren adds the "children" edges to the Organization entity.
func (ou *OrganizationUpdate) AddChildren(o ...*Organization) *OrganizationUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.AddChildIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role_Organization entity by IDs.
func (ou *OrganizationUpdate) AddRoleIDs(ids ...string) *OrganizationUpdate {
	ou.mutation.AddRoleIDs(ids...)
	return ou
}

// AddRoles adds the "roles" edges to the Role_Organization entity.
func (ou *OrganizationUpdate) AddRoles(r ...*Role_Organization) *OrganizationUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ou.AddRoleIDs(ids...)
}

// AddStaffIDs adds the "staffs" edge to the Staff entity by IDs.
func (ou *OrganizationUpdate) AddStaffIDs(ids ...string) *OrganizationUpdate {
	ou.mutation.AddStaffIDs(ids...)
	return ou
}

// AddStaffs adds the "staffs" edges to the Staff entity.
func (ou *OrganizationUpdate) AddStaffs(s ...*Staff) *OrganizationUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ou.AddStaffIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ou *OrganizationUpdate) Mutation() *OrganizationMutation {
	return ou.mutation
}

// ClearParent clears the "parent" edge to the Organization entity.
func (ou *OrganizationUpdate) ClearParent() *OrganizationUpdate {
	ou.mutation.ClearParent()
	return ou
}

// ClearChildren clears all "children" edges to the Organization entity.
func (ou *OrganizationUpdate) ClearChildren() *OrganizationUpdate {
	ou.mutation.ClearChildren()
	return ou
}

// RemoveChildIDs removes the "children" edge to Organization entities by IDs.
func (ou *OrganizationUpdate) RemoveChildIDs(ids ...string) *OrganizationUpdate {
	ou.mutation.RemoveChildIDs(ids...)
	return ou
}

// RemoveChildren removes "children" edges to Organization entities.
func (ou *OrganizationUpdate) RemoveChildren(o ...*Organization) *OrganizationUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ou.RemoveChildIDs(ids...)
}

// ClearRoles clears all "roles" edges to the Role_Organization entity.
func (ou *OrganizationUpdate) ClearRoles() *OrganizationUpdate {
	ou.mutation.ClearRoles()
	return ou
}

// RemoveRoleIDs removes the "roles" edge to Role_Organization entities by IDs.
func (ou *OrganizationUpdate) RemoveRoleIDs(ids ...string) *OrganizationUpdate {
	ou.mutation.RemoveRoleIDs(ids...)
	return ou
}

// RemoveRoles removes "roles" edges to Role_Organization entities.
func (ou *OrganizationUpdate) RemoveRoles(r ...*Role_Organization) *OrganizationUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ou.RemoveRoleIDs(ids...)
}

// ClearStaffs clears all "staffs" edges to the Staff entity.
func (ou *OrganizationUpdate) ClearStaffs() *OrganizationUpdate {
	ou.mutation.ClearStaffs()
	return ou
}

// RemoveStaffIDs removes the "staffs" edge to Staff entities by IDs.
func (ou *OrganizationUpdate) RemoveStaffIDs(ids ...string) *OrganizationUpdate {
	ou.mutation.RemoveStaffIDs(ids...)
	return ou
}

// RemoveStaffs removes "staffs" edges to Staff entities.
func (ou *OrganizationUpdate) RemoveStaffs(s ...*Staff) *OrganizationUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ou.RemoveStaffIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrganizationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrganizationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrganizationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrganizationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrganizationUpdate) check() error {
	if v, ok := ou.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Code(); ok {
		if err := organization.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Organization.code": %w`, err)}
		}
	}
	return nil
}

func (ou *OrganizationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ou.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(organization.FieldUpdatedAt, field.TypeInt64, value)
	}
	if ou.mutation.CreatedByCleared() {
		_spec.ClearField(organization.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ou.mutation.UpdatedBy(); ok {
		_spec.SetField(organization.FieldUpdatedBy, field.TypeString, value)
	}
	if ou.mutation.UpdatedByCleared() {
		_spec.ClearField(organization.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ou.mutation.Status(); ok {
		_spec.SetField(organization.FieldStatus, field.TypeString, value)
	}
	if ou.mutation.StatusCleared() {
		_spec.ClearField(organization.FieldStatus, field.TypeString)
	}
	if value, ok := ou.mutation.Sort(); ok {
		_spec.SetField(organization.FieldSort, field.TypeInt32, value)
	}
	if value, ok := ou.mutation.AddedSort(); ok {
		_spec.AddField(organization.FieldSort, field.TypeInt32, value)
	}
	if ou.mutation.SortCleared() {
		_spec.ClearField(organization.FieldSort, field.TypeInt32)
	}
	if value, ok := ou.mutation.Remark(); ok {
		_spec.SetField(organization.FieldRemark, field.TypeString, value)
	}
	if ou.mutation.RemarkCleared() {
		_spec.ClearField(organization.FieldRemark, field.TypeString)
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.Code(); ok {
		_spec.SetField(organization.FieldCode, field.TypeString, value)
	}
	if ou.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
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
	if ou.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ou.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
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
	if ou.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.RolesTable,
			Columns: []string{organization.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedRolesIDs(); len(nodes) > 0 && !ou.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.RolesTable,
			Columns: []string{organization.RolesColumn},
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
	if nodes := ou.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.RolesTable,
			Columns: []string{organization.RolesColumn},
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
	if ou.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.StaffsTable,
			Columns: []string{organization.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedStaffsIDs(); len(nodes) > 0 && !ou.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.StaffsTable,
			Columns: []string{organization.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.StaffsTable,
			Columns: []string{organization.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrganizationUpdateOne is the builder for updating a single Organization entity.
type OrganizationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrganizationUpdateOne) SetUpdatedAt(i int64) *OrganizationUpdateOne {
	ouo.mutation.ResetUpdatedAt()
	ouo.mutation.SetUpdatedAt(i)
	return ouo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableUpdatedAt(i *int64) *OrganizationUpdateOne {
	if i != nil {
		ouo.SetUpdatedAt(*i)
	}
	return ouo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ouo *OrganizationUpdateOne) AddUpdatedAt(i int64) *OrganizationUpdateOne {
	ouo.mutation.AddUpdatedAt(i)
	return ouo
}

// SetUpdatedBy sets the "updated_by" field.
func (ouo *OrganizationUpdateOne) SetUpdatedBy(s string) *OrganizationUpdateOne {
	ouo.mutation.SetUpdatedBy(s)
	return ouo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableUpdatedBy(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetUpdatedBy(*s)
	}
	return ouo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ouo *OrganizationUpdateOne) ClearUpdatedBy() *OrganizationUpdateOne {
	ouo.mutation.ClearUpdatedBy()
	return ouo
}

// SetStatus sets the "status" field.
func (ouo *OrganizationUpdateOne) SetStatus(s string) *OrganizationUpdateOne {
	ouo.mutation.SetStatus(s)
	return ouo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableStatus(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetStatus(*s)
	}
	return ouo
}

// ClearStatus clears the value of the "status" field.
func (ouo *OrganizationUpdateOne) ClearStatus() *OrganizationUpdateOne {
	ouo.mutation.ClearStatus()
	return ouo
}

// SetSort sets the "sort" field.
func (ouo *OrganizationUpdateOne) SetSort(i int32) *OrganizationUpdateOne {
	ouo.mutation.ResetSort()
	ouo.mutation.SetSort(i)
	return ouo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableSort(i *int32) *OrganizationUpdateOne {
	if i != nil {
		ouo.SetSort(*i)
	}
	return ouo
}

// AddSort adds i to the "sort" field.
func (ouo *OrganizationUpdateOne) AddSort(i int32) *OrganizationUpdateOne {
	ouo.mutation.AddSort(i)
	return ouo
}

// ClearSort clears the value of the "sort" field.
func (ouo *OrganizationUpdateOne) ClearSort() *OrganizationUpdateOne {
	ouo.mutation.ClearSort()
	return ouo
}

// SetRemark sets the "remark" field.
func (ouo *OrganizationUpdateOne) SetRemark(s string) *OrganizationUpdateOne {
	ouo.mutation.SetRemark(s)
	return ouo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableRemark(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetRemark(*s)
	}
	return ouo
}

// ClearRemark clears the value of the "remark" field.
func (ouo *OrganizationUpdateOne) ClearRemark() *OrganizationUpdateOne {
	ouo.mutation.ClearRemark()
	return ouo
}

// SetName sets the "name" field.
func (ouo *OrganizationUpdateOne) SetName(s string) *OrganizationUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableName(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetName(*s)
	}
	return ouo
}

// SetCode sets the "code" field.
func (ouo *OrganizationUpdateOne) SetCode(s string) *OrganizationUpdateOne {
	ouo.mutation.SetCode(s)
	return ouo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableCode(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetCode(*s)
	}
	return ouo
}

// SetParentID sets the "parent_id" field.
func (ouo *OrganizationUpdateOne) SetParentID(s string) *OrganizationUpdateOne {
	ouo.mutation.SetParentID(s)
	return ouo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableParentID(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetParentID(*s)
	}
	return ouo
}

// ClearParentID clears the value of the "parent_id" field.
func (ouo *OrganizationUpdateOne) ClearParentID() *OrganizationUpdateOne {
	ouo.mutation.ClearParentID()
	return ouo
}

// SetParent sets the "parent" edge to the Organization entity.
func (ouo *OrganizationUpdateOne) SetParent(o *Organization) *OrganizationUpdateOne {
	return ouo.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the Organization entity by IDs.
func (ouo *OrganizationUpdateOne) AddChildIDs(ids ...string) *OrganizationUpdateOne {
	ouo.mutation.AddChildIDs(ids...)
	return ouo
}

// AddChildren adds the "children" edges to the Organization entity.
func (ouo *OrganizationUpdateOne) AddChildren(o ...*Organization) *OrganizationUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.AddChildIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role_Organization entity by IDs.
func (ouo *OrganizationUpdateOne) AddRoleIDs(ids ...string) *OrganizationUpdateOne {
	ouo.mutation.AddRoleIDs(ids...)
	return ouo
}

// AddRoles adds the "roles" edges to the Role_Organization entity.
func (ouo *OrganizationUpdateOne) AddRoles(r ...*Role_Organization) *OrganizationUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ouo.AddRoleIDs(ids...)
}

// AddStaffIDs adds the "staffs" edge to the Staff entity by IDs.
func (ouo *OrganizationUpdateOne) AddStaffIDs(ids ...string) *OrganizationUpdateOne {
	ouo.mutation.AddStaffIDs(ids...)
	return ouo
}

// AddStaffs adds the "staffs" edges to the Staff entity.
func (ouo *OrganizationUpdateOne) AddStaffs(s ...*Staff) *OrganizationUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ouo.AddStaffIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ouo *OrganizationUpdateOne) Mutation() *OrganizationMutation {
	return ouo.mutation
}

// ClearParent clears the "parent" edge to the Organization entity.
func (ouo *OrganizationUpdateOne) ClearParent() *OrganizationUpdateOne {
	ouo.mutation.ClearParent()
	return ouo
}

// ClearChildren clears all "children" edges to the Organization entity.
func (ouo *OrganizationUpdateOne) ClearChildren() *OrganizationUpdateOne {
	ouo.mutation.ClearChildren()
	return ouo
}

// RemoveChildIDs removes the "children" edge to Organization entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveChildIDs(ids ...string) *OrganizationUpdateOne {
	ouo.mutation.RemoveChildIDs(ids...)
	return ouo
}

// RemoveChildren removes "children" edges to Organization entities.
func (ouo *OrganizationUpdateOne) RemoveChildren(o ...*Organization) *OrganizationUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ouo.RemoveChildIDs(ids...)
}

// ClearRoles clears all "roles" edges to the Role_Organization entity.
func (ouo *OrganizationUpdateOne) ClearRoles() *OrganizationUpdateOne {
	ouo.mutation.ClearRoles()
	return ouo
}

// RemoveRoleIDs removes the "roles" edge to Role_Organization entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveRoleIDs(ids ...string) *OrganizationUpdateOne {
	ouo.mutation.RemoveRoleIDs(ids...)
	return ouo
}

// RemoveRoles removes "roles" edges to Role_Organization entities.
func (ouo *OrganizationUpdateOne) RemoveRoles(r ...*Role_Organization) *OrganizationUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ouo.RemoveRoleIDs(ids...)
}

// ClearStaffs clears all "staffs" edges to the Staff entity.
func (ouo *OrganizationUpdateOne) ClearStaffs() *OrganizationUpdateOne {
	ouo.mutation.ClearStaffs()
	return ouo
}

// RemoveStaffIDs removes the "staffs" edge to Staff entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveStaffIDs(ids ...string) *OrganizationUpdateOne {
	ouo.mutation.RemoveStaffIDs(ids...)
	return ouo
}

// RemoveStaffs removes "staffs" edges to Staff entities.
func (ouo *OrganizationUpdateOne) RemoveStaffs(s ...*Staff) *OrganizationUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ouo.RemoveStaffIDs(ids...)
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ouo *OrganizationUpdateOne) Where(ps ...predicate.Organization) *OrganizationUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrganizationUpdateOne) Select(field string, fields ...string) *OrganizationUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Organization entity.
func (ouo *OrganizationUpdateOne) Save(ctx context.Context) (*Organization, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) SaveX(ctx context.Context) *Organization {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrganizationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrganizationUpdateOne) check() error {
	if v, ok := ouo.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Code(); ok {
		if err := organization.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Organization.code": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrganizationUpdateOne) sqlSave(ctx context.Context) (_node *Organization, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Organization.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for _, f := range fields {
			if !organization.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ouo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(organization.FieldUpdatedAt, field.TypeInt64, value)
	}
	if ouo.mutation.CreatedByCleared() {
		_spec.ClearField(organization.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ouo.mutation.UpdatedBy(); ok {
		_spec.SetField(organization.FieldUpdatedBy, field.TypeString, value)
	}
	if ouo.mutation.UpdatedByCleared() {
		_spec.ClearField(organization.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ouo.mutation.Status(); ok {
		_spec.SetField(organization.FieldStatus, field.TypeString, value)
	}
	if ouo.mutation.StatusCleared() {
		_spec.ClearField(organization.FieldStatus, field.TypeString)
	}
	if value, ok := ouo.mutation.Sort(); ok {
		_spec.SetField(organization.FieldSort, field.TypeInt32, value)
	}
	if value, ok := ouo.mutation.AddedSort(); ok {
		_spec.AddField(organization.FieldSort, field.TypeInt32, value)
	}
	if ouo.mutation.SortCleared() {
		_spec.ClearField(organization.FieldSort, field.TypeInt32)
	}
	if value, ok := ouo.mutation.Remark(); ok {
		_spec.SetField(organization.FieldRemark, field.TypeString, value)
	}
	if ouo.mutation.RemarkCleared() {
		_spec.ClearField(organization.FieldRemark, field.TypeString)
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Code(); ok {
		_spec.SetField(organization.FieldCode, field.TypeString, value)
	}
	if ouo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
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
	if ouo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !ouo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
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
	if ouo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.RolesTable,
			Columns: []string{organization.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_organization.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !ouo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.RolesTable,
			Columns: []string{organization.RolesColumn},
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
	if nodes := ouo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.RolesTable,
			Columns: []string{organization.RolesColumn},
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
	if ouo.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.StaffsTable,
			Columns: []string{organization.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedStaffsIDs(); len(nodes) > 0 && !ouo.mutation.StaffsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.StaffsTable,
			Columns: []string{organization.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.StaffsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.StaffsTable,
			Columns: []string{organization.StaffsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(staff.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Organization{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}