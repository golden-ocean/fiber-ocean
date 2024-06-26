// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/golden-ocean/fiber-ocean/ent/position"
	"github.com/golden-ocean/fiber-ocean/ent/predicate"
	"github.com/golden-ocean/fiber-ocean/ent/staff"
	"github.com/golden-ocean/fiber-ocean/ent/staff_position"
)

// StaffPositionQuery is the builder for querying Staff_Position entities.
type StaffPositionQuery struct {
	config
	ctx           *QueryContext
	order         []staff_position.OrderOption
	inters        []Interceptor
	predicates    []predicate.Staff_Position
	withStaffs    *StaffQuery
	withPositions *PositionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StaffPositionQuery builder.
func (spq *StaffPositionQuery) Where(ps ...predicate.Staff_Position) *StaffPositionQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *StaffPositionQuery) Limit(limit int) *StaffPositionQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *StaffPositionQuery) Offset(offset int) *StaffPositionQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *StaffPositionQuery) Unique(unique bool) *StaffPositionQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *StaffPositionQuery) Order(o ...staff_position.OrderOption) *StaffPositionQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryStaffs chains the current query on the "staffs" edge.
func (spq *StaffPositionQuery) QueryStaffs() *StaffQuery {
	query := (&StaffClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(staff_position.Table, staff_position.FieldID, selector),
			sqlgraph.To(staff.Table, staff.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, staff_position.StaffsTable, staff_position.StaffsColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPositions chains the current query on the "positions" edge.
func (spq *StaffPositionQuery) QueryPositions() *PositionQuery {
	query := (&PositionClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(staff_position.Table, staff_position.FieldID, selector),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, staff_position.PositionsTable, staff_position.PositionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Staff_Position entity from the query.
// Returns a *NotFoundError when no Staff_Position was found.
func (spq *StaffPositionQuery) First(ctx context.Context) (*Staff_Position, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{staff_position.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *StaffPositionQuery) FirstX(ctx context.Context) *Staff_Position {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Staff_Position ID from the query.
// Returns a *NotFoundError when no Staff_Position ID was found.
func (spq *StaffPositionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{staff_position.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *StaffPositionQuery) FirstIDX(ctx context.Context) string {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Staff_Position entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Staff_Position entity is found.
// Returns a *NotFoundError when no Staff_Position entities are found.
func (spq *StaffPositionQuery) Only(ctx context.Context) (*Staff_Position, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{staff_position.Label}
	default:
		return nil, &NotSingularError{staff_position.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *StaffPositionQuery) OnlyX(ctx context.Context) *Staff_Position {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Staff_Position ID in the query.
// Returns a *NotSingularError when more than one Staff_Position ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *StaffPositionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{staff_position.Label}
	default:
		err = &NotSingularError{staff_position.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *StaffPositionQuery) OnlyIDX(ctx context.Context) string {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Staff_Positions.
func (spq *StaffPositionQuery) All(ctx context.Context) ([]*Staff_Position, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Staff_Position, *StaffPositionQuery]()
	return withInterceptors[[]*Staff_Position](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *StaffPositionQuery) AllX(ctx context.Context) []*Staff_Position {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Staff_Position IDs.
func (spq *StaffPositionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(staff_position.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *StaffPositionQuery) IDsX(ctx context.Context) []string {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *StaffPositionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*StaffPositionQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *StaffPositionQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *StaffPositionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *StaffPositionQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StaffPositionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *StaffPositionQuery) Clone() *StaffPositionQuery {
	if spq == nil {
		return nil
	}
	return &StaffPositionQuery{
		config:        spq.config,
		ctx:           spq.ctx.Clone(),
		order:         append([]staff_position.OrderOption{}, spq.order...),
		inters:        append([]Interceptor{}, spq.inters...),
		predicates:    append([]predicate.Staff_Position{}, spq.predicates...),
		withStaffs:    spq.withStaffs.Clone(),
		withPositions: spq.withPositions.Clone(),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// WithStaffs tells the query-builder to eager-load the nodes that are connected to
// the "staffs" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StaffPositionQuery) WithStaffs(opts ...func(*StaffQuery)) *StaffPositionQuery {
	query := (&StaffClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withStaffs = query
	return spq
}

// WithPositions tells the query-builder to eager-load the nodes that are connected to
// the "positions" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StaffPositionQuery) WithPositions(opts ...func(*PositionQuery)) *StaffPositionQuery {
	query := (&PositionClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withPositions = query
	return spq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.StaffPosition.Query().
//		GroupBy(staff_position.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (spq *StaffPositionQuery) GroupBy(field string, fields ...string) *StaffPositionGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StaffPositionGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = staff_position.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//	}
//
//	client.StaffPosition.Query().
//		Select(staff_position.FieldCreatedAt).
//		Scan(ctx, &v)
func (spq *StaffPositionQuery) Select(fields ...string) *StaffPositionSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &StaffPositionSelect{StaffPositionQuery: spq}
	sbuild.label = staff_position.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StaffPositionSelect configured with the given aggregations.
func (spq *StaffPositionQuery) Aggregate(fns ...AggregateFunc) *StaffPositionSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *StaffPositionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	for _, f := range spq.ctx.Fields {
		if !staff_position.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *StaffPositionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Staff_Position, error) {
	var (
		nodes       = []*Staff_Position{}
		_spec       = spq.querySpec()
		loadedTypes = [2]bool{
			spq.withStaffs != nil,
			spq.withPositions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Staff_Position).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Staff_Position{config: spq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := spq.withStaffs; query != nil {
		if err := spq.loadStaffs(ctx, query, nodes, nil,
			func(n *Staff_Position, e *Staff) { n.Edges.Staffs = e }); err != nil {
			return nil, err
		}
	}
	if query := spq.withPositions; query != nil {
		if err := spq.loadPositions(ctx, query, nodes, nil,
			func(n *Staff_Position, e *Position) { n.Edges.Positions = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (spq *StaffPositionQuery) loadStaffs(ctx context.Context, query *StaffQuery, nodes []*Staff_Position, init func(*Staff_Position), assign func(*Staff_Position, *Staff)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Staff_Position)
	for i := range nodes {
		fk := nodes[i].StaffID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(staff.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "staff_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (spq *StaffPositionQuery) loadPositions(ctx context.Context, query *PositionQuery, nodes []*Staff_Position, init func(*Staff_Position), assign func(*Staff_Position, *Position)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Staff_Position)
	for i := range nodes {
		fk := nodes[i].PositionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(position.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "position_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (spq *StaffPositionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.ctx.Fields
	if len(spq.ctx.Fields) > 0 {
		_spec.Unique = spq.ctx.Unique != nil && *spq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *StaffPositionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(staff_position.Table, staff_position.Columns, sqlgraph.NewFieldSpec(staff_position.FieldID, field.TypeString))
	_spec.From = spq.sql
	if unique := spq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if spq.path != nil {
		_spec.Unique = true
	}
	if fields := spq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, staff_position.FieldID)
		for i := range fields {
			if fields[i] != staff_position.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if spq.withStaffs != nil {
			_spec.Node.AddColumnOnce(staff_position.FieldStaffID)
		}
		if spq.withPositions != nil {
			_spec.Node.AddColumnOnce(staff_position.FieldPositionID)
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *StaffPositionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(staff_position.Table)
	columns := spq.ctx.Fields
	if len(columns) == 0 {
		columns = staff_position.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.ctx.Unique != nil && *spq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StaffPositionGroupBy is the group-by builder for Staff_Position entities.
type StaffPositionGroupBy struct {
	selector
	build *StaffPositionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *StaffPositionGroupBy) Aggregate(fns ...AggregateFunc) *StaffPositionGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *StaffPositionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StaffPositionQuery, *StaffPositionGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *StaffPositionGroupBy) sqlScan(ctx context.Context, root *StaffPositionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*spgb.flds)+len(spgb.fns))
		for _, f := range *spgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*spgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StaffPositionSelect is the builder for selecting fields of StaffPosition entities.
type StaffPositionSelect struct {
	*StaffPositionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *StaffPositionSelect) Aggregate(fns ...AggregateFunc) *StaffPositionSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *StaffPositionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StaffPositionQuery, *StaffPositionSelect](ctx, sps.StaffPositionQuery, sps, sps.inters, v)
}

func (sps *StaffPositionSelect) sqlScan(ctx context.Context, root *StaffPositionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sps.fns))
	for _, fn := range sps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
