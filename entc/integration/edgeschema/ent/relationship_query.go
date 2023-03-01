// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/predicate"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/relationship"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/user"
)

// RelationshipQuery is the builder for querying Relationship entities.
type RelationshipQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.Relationship
	withUser     *UserQuery
	withRelative *UserQuery
	withInfo     *RelationshipInfoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RelationshipQuery builder.
func (rq *RelationshipQuery) Where(ps ...predicate.Relationship) *RelationshipQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RelationshipQuery) Limit(limit int) *RelationshipQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RelationshipQuery) Offset(offset int) *RelationshipQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RelationshipQuery) Unique(unique bool) *RelationshipQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RelationshipQuery) Order(o ...OrderFunc) *RelationshipQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryUser chains the current query on the "user" edge.
func (rq *RelationshipQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relationship.Table, relationship.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, relationship.UserTable, relationship.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRelative chains the current query on the "relative" edge.
func (rq *RelationshipQuery) QueryRelative() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relationship.Table, relationship.RelativeColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, relationship.RelativeTable, relationship.RelativeColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInfo chains the current query on the "info" edge.
func (rq *RelationshipQuery) QueryInfo() *RelationshipInfoQuery {
	query := (&RelationshipInfoClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(relationship.Table, relationship.InfoColumn, selector),
			sqlgraph.To(relationshipinfo.Table, relationshipinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, relationship.InfoTable, relationship.InfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Relationship entity from the query.
// Returns a *NotFoundError when no Relationship was found.
func (rq *RelationshipQuery) First(ctx context.Context) (*Relationship, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{relationship.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RelationshipQuery) FirstX(ctx context.Context) *Relationship {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single Relationship entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Relationship entity is found.
// Returns a *NotFoundError when no Relationship entities are found.
func (rq *RelationshipQuery) Only(ctx context.Context) (*Relationship, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{relationship.Label}
	default:
		return nil, &NotSingularError{relationship.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RelationshipQuery) OnlyX(ctx context.Context) *Relationship {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of Relationships.
func (rq *RelationshipQuery) All(ctx context.Context) ([]*Relationship, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Relationship, *RelationshipQuery]()
	return withInterceptors[[]*Relationship](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RelationshipQuery) AllX(ctx context.Context) []*Relationship {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (rq *RelationshipQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RelationshipQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RelationshipQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RelationshipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RelationshipQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RelationshipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RelationshipQuery) Clone() *RelationshipQuery {
	if rq == nil {
		return nil
	}
	return &RelationshipQuery{
		config:       rq.config,
		ctx:          rq.ctx.Clone(),
		order:        append([]OrderFunc{}, rq.order...),
		inters:       append([]Interceptor{}, rq.inters...),
		predicates:   append([]predicate.Relationship{}, rq.predicates...),
		withUser:     rq.withUser.Clone(),
		withRelative: rq.withRelative.Clone(),
		withInfo:     rq.withInfo.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RelationshipQuery) WithUser(opts ...func(*UserQuery)) *RelationshipQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUser = query
	return rq
}

// WithRelative tells the query-builder to eager-load the nodes that are connected to
// the "relative" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RelationshipQuery) WithRelative(opts ...func(*UserQuery)) *RelationshipQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withRelative = query
	return rq
}

// WithInfo tells the query-builder to eager-load the nodes that are connected to
// the "info" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RelationshipQuery) WithInfo(opts ...func(*RelationshipInfoQuery)) *RelationshipQuery {
	query := (&RelationshipInfoClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withInfo = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Weight int `json:"weight,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Relationship.Query().
//		GroupBy(relationship.FieldWeight).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RelationshipQuery) GroupBy(field string, fields ...string) *RelationshipGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RelationshipGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = relationship.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Weight int `json:"weight,omitempty"`
//	}
//
//	client.Relationship.Query().
//		Select(relationship.FieldWeight).
//		Scan(ctx, &v)
func (rq *RelationshipQuery) Select(fields ...string) *RelationshipSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RelationshipSelect{RelationshipQuery: rq}
	sbuild.label = relationship.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RelationshipSelect configured with the given aggregations.
func (rq *RelationshipQuery) Aggregate(fns ...AggregateFunc) *RelationshipSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RelationshipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !relationship.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	if relationship.Policy == nil {
		return errors.New("ent: uninitialized relationship.Policy (forgotten import ent/runtime?)")
	}
	if err := relationship.Policy.EvalQuery(ctx, rq); err != nil {
		return err
	}
	return nil
}

func (rq *RelationshipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Relationship, error) {
	var (
		nodes       = []*Relationship{}
		_spec       = rq.querySpec()
		loadedTypes = [3]bool{
			rq.withUser != nil,
			rq.withRelative != nil,
			rq.withInfo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Relationship).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Relationship{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withUser; query != nil {
		if err := rq.loadUser(ctx, query, nodes, nil,
			func(n *Relationship, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withRelative; query != nil {
		if err := rq.loadRelative(ctx, query, nodes, nil,
			func(n *Relationship, e *User) { n.Edges.Relative = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withInfo; query != nil {
		if err := rq.loadInfo(ctx, query, nodes, nil,
			func(n *Relationship, e *RelationshipInfo) { n.Edges.Info = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RelationshipQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Relationship, init func(*Relationship), assign func(*Relationship, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Relationship)
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
func (rq *RelationshipQuery) loadRelative(ctx context.Context, query *UserQuery, nodes []*Relationship, init func(*Relationship), assign func(*Relationship, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Relationship)
	for i := range nodes {
		fk := nodes[i].RelativeID
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
			return fmt.Errorf(`unexpected foreign-key "relative_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *RelationshipQuery) loadInfo(ctx context.Context, query *RelationshipInfoQuery, nodes []*Relationship, init func(*Relationship), assign func(*Relationship, *RelationshipInfo)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Relationship)
	for i := range nodes {
		fk := nodes[i].InfoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(relationshipinfo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "info_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *RelationshipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RelationshipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(relationship.Table, relationship.Columns, nil)
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RelationshipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(relationship.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = relationship.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RelationshipGroupBy is the group-by builder for Relationship entities.
type RelationshipGroupBy struct {
	selector
	build *RelationshipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RelationshipGroupBy) Aggregate(fns ...AggregateFunc) *RelationshipGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RelationshipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RelationshipQuery, *RelationshipGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RelationshipGroupBy) sqlScan(ctx context.Context, root *RelationshipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RelationshipSelect is the builder for selecting fields of Relationship entities.
type RelationshipSelect struct {
	*RelationshipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RelationshipSelect) Aggregate(fns ...AggregateFunc) *RelationshipSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RelationshipSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RelationshipQuery, *RelationshipSelect](ctx, rs.RelationshipQuery, rs, rs.inters, v)
}

func (rs *RelationshipSelect) sqlScan(ctx context.Context, root *RelationshipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
