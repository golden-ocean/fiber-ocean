// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/menu"
	"github.com/golden-ocean/fiber-ocean/ent/role_menu"
)

// MenuCreate is the builder for creating a Menu entity.
type MenuCreate struct {
	config
	mutation *MenuMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MenuCreate) SetCreatedAt(i int64) *MenuCreate {
	mc.mutation.SetCreatedAt(i)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCreatedAt(i *int64) *MenuCreate {
	if i != nil {
		mc.SetCreatedAt(*i)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MenuCreate) SetUpdatedAt(i int64) *MenuCreate {
	mc.mutation.SetUpdatedAt(i)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableUpdatedAt(i *int64) *MenuCreate {
	if i != nil {
		mc.SetUpdatedAt(*i)
	}
	return mc
}

// SetCreatedBy sets the "created_by" field.
func (mc *MenuCreate) SetCreatedBy(s string) *MenuCreate {
	mc.mutation.SetCreatedBy(s)
	return mc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCreatedBy(s *string) *MenuCreate {
	if s != nil {
		mc.SetCreatedBy(*s)
	}
	return mc
}

// SetUpdatedBy sets the "updated_by" field.
func (mc *MenuCreate) SetUpdatedBy(s string) *MenuCreate {
	mc.mutation.SetUpdatedBy(s)
	return mc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mc *MenuCreate) SetNillableUpdatedBy(s *string) *MenuCreate {
	if s != nil {
		mc.SetUpdatedBy(*s)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MenuCreate) SetStatus(s string) *MenuCreate {
	mc.mutation.SetStatus(s)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *MenuCreate) SetNillableStatus(s *string) *MenuCreate {
	if s != nil {
		mc.SetStatus(*s)
	}
	return mc
}

// SetSort sets the "sort" field.
func (mc *MenuCreate) SetSort(i int32) *MenuCreate {
	mc.mutation.SetSort(i)
	return mc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mc *MenuCreate) SetNillableSort(i *int32) *MenuCreate {
	if i != nil {
		mc.SetSort(*i)
	}
	return mc
}

// SetRemark sets the "remark" field.
func (mc *MenuCreate) SetRemark(s string) *MenuCreate {
	mc.mutation.SetRemark(s)
	return mc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (mc *MenuCreate) SetNillableRemark(s *string) *MenuCreate {
	if s != nil {
		mc.SetRemark(*s)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *MenuCreate) SetName(s string) *MenuCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetParentID sets the "parent_id" field.
func (mc *MenuCreate) SetParentID(s string) *MenuCreate {
	mc.mutation.SetParentID(s)
	return mc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mc *MenuCreate) SetNillableParentID(s *string) *MenuCreate {
	if s != nil {
		mc.SetParentID(*s)
	}
	return mc
}

// SetIcon sets the "icon" field.
func (mc *MenuCreate) SetIcon(s string) *MenuCreate {
	mc.mutation.SetIcon(s)
	return mc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (mc *MenuCreate) SetNillableIcon(s *string) *MenuCreate {
	if s != nil {
		mc.SetIcon(*s)
	}
	return mc
}

// SetPath sets the "path" field.
func (mc *MenuCreate) SetPath(s string) *MenuCreate {
	mc.mutation.SetPath(s)
	return mc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mc *MenuCreate) SetNillablePath(s *string) *MenuCreate {
	if s != nil {
		mc.SetPath(*s)
	}
	return mc
}

// SetPermission sets the "permission" field.
func (mc *MenuCreate) SetPermission(s string) *MenuCreate {
	mc.mutation.SetPermission(s)
	return mc
}

// SetNillablePermission sets the "permission" field if the given value is not nil.
func (mc *MenuCreate) SetNillablePermission(s *string) *MenuCreate {
	if s != nil {
		mc.SetPermission(*s)
	}
	return mc
}

// SetComponent sets the "component" field.
func (mc *MenuCreate) SetComponent(s string) *MenuCreate {
	mc.mutation.SetComponent(s)
	return mc
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (mc *MenuCreate) SetNillableComponent(s *string) *MenuCreate {
	if s != nil {
		mc.SetComponent(*s)
	}
	return mc
}

// SetType sets the "type" field.
func (mc *MenuCreate) SetType(s string) *MenuCreate {
	mc.mutation.SetType(s)
	return mc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mc *MenuCreate) SetNillableType(s *string) *MenuCreate {
	if s != nil {
		mc.SetType(*s)
	}
	return mc
}

// SetMethod sets the "method" field.
func (mc *MenuCreate) SetMethod(s string) *MenuCreate {
	mc.mutation.SetMethod(s)
	return mc
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (mc *MenuCreate) SetNillableMethod(s *string) *MenuCreate {
	if s != nil {
		mc.SetMethod(*s)
	}
	return mc
}

// SetVisible sets the "visible" field.
func (mc *MenuCreate) SetVisible(b bool) *MenuCreate {
	mc.mutation.SetVisible(b)
	return mc
}

// SetNillableVisible sets the "visible" field if the given value is not nil.
func (mc *MenuCreate) SetNillableVisible(b *bool) *MenuCreate {
	if b != nil {
		mc.SetVisible(*b)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MenuCreate) SetID(s string) *MenuCreate {
	mc.mutation.SetID(s)
	return mc
}

// SetParent sets the "parent" edge to the Menu entity.
func (mc *MenuCreate) SetParent(m *Menu) *MenuCreate {
	return mc.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (mc *MenuCreate) AddChildIDs(ids ...string) *MenuCreate {
	mc.mutation.AddChildIDs(ids...)
	return mc
}

// AddChildren adds the "children" edges to the Menu entity.
func (mc *MenuCreate) AddChildren(m ...*Menu) *MenuCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddChildIDs(ids...)
}

// AddRoleIDs adds the "roles" edge to the Role_Menu entity by IDs.
func (mc *MenuCreate) AddRoleIDs(ids ...string) *MenuCreate {
	mc.mutation.AddRoleIDs(ids...)
	return mc
}

// AddRoles adds the "roles" edges to the Role_Menu entity.
func (mc *MenuCreate) AddRoles(r ...*Role_Menu) *MenuCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRoleIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mc *MenuCreate) Mutation() *MenuMutation {
	return mc.mutation
}

// Save creates the Menu in the database.
func (mc *MenuCreate) Save(ctx context.Context) (*Menu, error) {
	if err := mc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MenuCreate) SaveX(ctx context.Context) *Menu {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MenuCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MenuCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MenuCreate) defaults() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		if menu.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized menu.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := menu.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		if menu.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized menu.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := menu.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := menu.DefaultStatus
		mc.mutation.SetStatus(v)
	}
	if _, ok := mc.mutation.Sort(); !ok {
		v := menu.DefaultSort
		mc.mutation.SetSort(v)
	}
	if _, ok := mc.mutation.Visible(); !ok {
		v := menu.DefaultVisible
		mc.mutation.SetVisible(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mc *MenuCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Menu.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Menu.updated_at"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Menu.name"`)}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Menu.name": %w`, err)}
		}
	}
	if v, ok := mc.mutation.ID(); ok {
		if err := menu.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Menu.id": %w`, err)}
		}
	}
	return nil
}

func (mc *MenuCreate) sqlSave(ctx context.Context) (*Menu, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Menu.ID type: %T", _spec.ID.Value)
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MenuCreate) createSpec() (*Menu, *sqlgraph.CreateSpec) {
	var (
		_node = &Menu{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(menu.Table, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(menu.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.CreatedBy(); ok {
		_spec.SetField(menu.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := mc.mutation.UpdatedBy(); ok {
		_spec.SetField(menu.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := mc.mutation.Sort(); ok {
		_spec.SetField(menu.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := mc.mutation.Remark(); ok {
		_spec.SetField(menu.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Icon(); ok {
		_spec.SetField(menu.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := mc.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := mc.mutation.Permission(); ok {
		_spec.SetField(menu.FieldPermission, field.TypeString, value)
		_node.Permission = value
	}
	if value, ok := mc.mutation.Component(); ok {
		_spec.SetField(menu.FieldComponent, field.TypeString, value)
		_node.Component = value
	}
	if value, ok := mc.mutation.GetType(); ok {
		_spec.SetField(menu.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mc.mutation.Method(); ok {
		_spec.SetField(menu.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	if value, ok := mc.mutation.Visible(); ok {
		_spec.SetField(menu.FieldVisible, field.TypeBool, value)
		_node.Visible = value
	}
	if nodes := mc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.RolesTable,
			Columns: []string{menu.RolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role_menu.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MenuCreateBulk is the builder for creating many Menu entities in bulk.
type MenuCreateBulk struct {
	config
	err      error
	builders []*MenuCreate
}

// Save creates the Menu entities in the database.
func (mcb *MenuCreateBulk) Save(ctx context.Context) ([]*Menu, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Menu, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MenuCreateBulk) SaveX(ctx context.Context) []*Menu {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MenuCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MenuCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
