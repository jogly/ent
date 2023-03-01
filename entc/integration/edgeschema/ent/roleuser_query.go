// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/predicate"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/role"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/roleuser"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/user"
)

// RoleUserQuery is the builder for querying RoleUser entities.
type RoleUserQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.RoleUser
	withRole   *RoleQuery
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RoleUserQuery builder.
func (ruq *RoleUserQuery) Where(ps ...predicate.RoleUser) *RoleUserQuery {
	ruq.predicates = append(ruq.predicates, ps...)
	return ruq
}

// Limit the number of records to be returned by this query.
func (ruq *RoleUserQuery) Limit(limit int) *RoleUserQuery {
	ruq.ctx.Limit = &limit
	return ruq
}

// Offset to start from.
func (ruq *RoleUserQuery) Offset(offset int) *RoleUserQuery {
	ruq.ctx.Offset = &offset
	return ruq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ruq *RoleUserQuery) Unique(unique bool) *RoleUserQuery {
	ruq.ctx.Unique = &unique
	return ruq
}

// Order specifies how the records should be ordered.
func (ruq *RoleUserQuery) Order(o ...OrderFunc) *RoleUserQuery {
	ruq.order = append(ruq.order, o...)
	return ruq
}

// QueryRole chains the current query on the "role" edge.
func (ruq *RoleUserQuery) QueryRole() *RoleQuery {
	query := (&RoleClient{config: ruq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ruq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ruq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roleuser.Table, roleuser.RoleColumn, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, roleuser.RoleTable, roleuser.RoleColumn),
		)
		fromU = sqlgraph.SetNeighbors(ruq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (ruq *RoleUserQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ruq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ruq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ruq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(roleuser.Table, roleuser.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, roleuser.UserTable, roleuser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ruq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RoleUser entity from the query.
// Returns a *NotFoundError when no RoleUser was found.
func (ruq *RoleUserQuery) First(ctx context.Context) (*RoleUser, error) {
	nodes, err := ruq.Limit(1).All(setContextOp(ctx, ruq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{roleuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ruq *RoleUserQuery) FirstX(ctx context.Context) *RoleUser {
	node, err := ruq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single RoleUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RoleUser entity is found.
// Returns a *NotFoundError when no RoleUser entities are found.
func (ruq *RoleUserQuery) Only(ctx context.Context) (*RoleUser, error) {
	nodes, err := ruq.Limit(2).All(setContextOp(ctx, ruq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{roleuser.Label}
	default:
		return nil, &NotSingularError{roleuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ruq *RoleUserQuery) OnlyX(ctx context.Context) *RoleUser {
	node, err := ruq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of RoleUsers.
func (ruq *RoleUserQuery) All(ctx context.Context) ([]*RoleUser, error) {
	ctx = setContextOp(ctx, ruq.ctx, "All")
	if err := ruq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RoleUser, *RoleUserQuery]()
	return withInterceptors[[]*RoleUser](ctx, ruq, qr, ruq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ruq *RoleUserQuery) AllX(ctx context.Context) []*RoleUser {
	nodes, err := ruq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (ruq *RoleUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ruq.ctx, "Count")
	if err := ruq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ruq, querierCount[*RoleUserQuery](), ruq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ruq *RoleUserQuery) CountX(ctx context.Context) int {
	count, err := ruq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ruq *RoleUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ruq.ctx, "Exist")
	switch _, err := ruq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ruq *RoleUserQuery) ExistX(ctx context.Context) bool {
	exist, err := ruq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RoleUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ruq *RoleUserQuery) Clone() *RoleUserQuery {
	if ruq == nil {
		return nil
	}
	return &RoleUserQuery{
		config:     ruq.config,
		ctx:        ruq.ctx.Clone(),
		order:      append([]OrderFunc{}, ruq.order...),
		inters:     append([]Interceptor{}, ruq.inters...),
		predicates: append([]predicate.RoleUser{}, ruq.predicates...),
		withRole:   ruq.withRole.Clone(),
		withUser:   ruq.withUser.Clone(),
		// clone intermediate query.
		sql:  ruq.sql.Clone(),
		path: ruq.path,
	}
}

// WithRole tells the query-builder to eager-load the nodes that are connected to
// the "role" edge. The optional arguments are used to configure the query builder of the edge.
func (ruq *RoleUserQuery) WithRole(opts ...func(*RoleQuery)) *RoleUserQuery {
	query := (&RoleClient{config: ruq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ruq.withRole = query
	return ruq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ruq *RoleUserQuery) WithUser(opts ...func(*UserQuery)) *RoleUserQuery {
	query := (&UserClient{config: ruq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ruq.withUser = query
	return ruq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RoleUser.Query().
//		GroupBy(roleuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ruq *RoleUserQuery) GroupBy(field string, fields ...string) *RoleUserGroupBy {
	ruq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RoleUserGroupBy{build: ruq}
	grbuild.flds = &ruq.ctx.Fields
	grbuild.label = roleuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.RoleUser.Query().
//		Select(roleuser.FieldCreatedAt).
//		Scan(ctx, &v)
func (ruq *RoleUserQuery) Select(fields ...string) *RoleUserSelect {
	ruq.ctx.Fields = append(ruq.ctx.Fields, fields...)
	sbuild := &RoleUserSelect{RoleUserQuery: ruq}
	sbuild.label = roleuser.Label
	sbuild.flds, sbuild.scan = &ruq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RoleUserSelect configured with the given aggregations.
func (ruq *RoleUserQuery) Aggregate(fns ...AggregateFunc) *RoleUserSelect {
	return ruq.Select().Aggregate(fns...)
}

func (ruq *RoleUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ruq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ruq); err != nil {
				return err
			}
		}
	}
	for _, f := range ruq.ctx.Fields {
		if !roleuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ruq.path != nil {
		prev, err := ruq.path(ctx)
		if err != nil {
			return err
		}
		ruq.sql = prev
	}
	return nil
}

func (ruq *RoleUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RoleUser, error) {
	var (
		nodes       = []*RoleUser{}
		_spec       = ruq.querySpec()
		loadedTypes = [2]bool{
			ruq.withRole != nil,
			ruq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RoleUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RoleUser{config: ruq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ruq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ruq.withRole; query != nil {
		if err := ruq.loadRole(ctx, query, nodes, nil,
			func(n *RoleUser, e *Role) { n.Edges.Role = e }); err != nil {
			return nil, err
		}
	}
	if query := ruq.withUser; query != nil {
		if err := ruq.loadUser(ctx, query, nodes, nil,
			func(n *RoleUser, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ruq *RoleUserQuery) loadRole(ctx context.Context, query *RoleQuery, nodes []*RoleUser, init func(*RoleUser), assign func(*RoleUser, *Role)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*RoleUser)
	for i := range nodes {
		fk := nodes[i].RoleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(role.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "role_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ruq *RoleUserQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*RoleUser, init func(*RoleUser), assign func(*RoleUser, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*RoleUser)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ruq *RoleUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ruq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, ruq.driver, _spec)
}

func (ruq *RoleUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(roleuser.Table, roleuser.Columns, nil)
	_spec.From = ruq.sql
	if unique := ruq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ruq.path != nil {
		_spec.Unique = true
	}
	if fields := ruq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := ruq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ruq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ruq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ruq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ruq *RoleUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ruq.driver.Dialect())
	t1 := builder.Table(roleuser.Table)
	columns := ruq.ctx.Fields
	if len(columns) == 0 {
		columns = roleuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ruq.sql != nil {
		selector = ruq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ruq.ctx.Unique != nil && *ruq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ruq.predicates {
		p(selector)
	}
	for _, p := range ruq.order {
		p(selector)
	}
	if offset := ruq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ruq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RoleUserGroupBy is the group-by builder for RoleUser entities.
type RoleUserGroupBy struct {
	selector
	build *RoleUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rugb *RoleUserGroupBy) Aggregate(fns ...AggregateFunc) *RoleUserGroupBy {
	rugb.fns = append(rugb.fns, fns...)
	return rugb
}

// Scan applies the selector query and scans the result into the given value.
func (rugb *RoleUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rugb.build.ctx, "GroupBy")
	if err := rugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleUserQuery, *RoleUserGroupBy](ctx, rugb.build, rugb, rugb.build.inters, v)
}

func (rugb *RoleUserGroupBy) sqlScan(ctx context.Context, root *RoleUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rugb.fns))
	for _, fn := range rugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rugb.flds)+len(rugb.fns))
		for _, f := range *rugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RoleUserSelect is the builder for selecting fields of RoleUser entities.
type RoleUserSelect struct {
	*RoleUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rus *RoleUserSelect) Aggregate(fns ...AggregateFunc) *RoleUserSelect {
	rus.fns = append(rus.fns, fns...)
	return rus
}

// Scan applies the selector query and scans the result into the given value.
func (rus *RoleUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rus.ctx, "Select")
	if err := rus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RoleUserQuery, *RoleUserSelect](ctx, rus.RoleUserQuery, rus, rus.inters, v)
}

func (rus *RoleUserSelect) sqlScan(ctx context.Context, root *RoleUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rus.fns))
	for _, fn := range rus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
