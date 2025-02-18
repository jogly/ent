// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"github.com/jogly/ent/dialect/gremlin"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl/__"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl/g"
	"github.com/jogly/ent/entc/integration/gremlin/ent/license"
	"github.com/jogly/ent/entc/integration/gremlin/ent/predicate"
)

// LicenseQuery is the builder for querying License entities.
type LicenseQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.License
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the LicenseQuery builder.
func (lq *LicenseQuery) Where(ps ...predicate.License) *LicenseQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LicenseQuery) Limit(limit int) *LicenseQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LicenseQuery) Offset(offset int) *LicenseQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LicenseQuery) Unique(unique bool) *LicenseQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LicenseQuery) Order(o ...OrderFunc) *LicenseQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// First returns the first License entity from the query.
// Returns a *NotFoundError when no License was found.
func (lq *LicenseQuery) First(ctx context.Context) (*License, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{license.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LicenseQuery) FirstX(ctx context.Context) *License {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first License ID from the query.
// Returns a *NotFoundError when no License ID was found.
func (lq *LicenseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{license.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LicenseQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single License entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one License entity is found.
// Returns a *NotFoundError when no License entities are found.
func (lq *LicenseQuery) Only(ctx context.Context) (*License, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{license.Label}
	default:
		return nil, &NotSingularError{license.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LicenseQuery) OnlyX(ctx context.Context) *License {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only License ID in the query.
// Returns a *NotSingularError when more than one License ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LicenseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{license.Label}
	default:
		err = &NotSingularError{license.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LicenseQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Licenses.
func (lq *LicenseQuery) All(ctx context.Context) ([]*License, error) {
	ctx = setContextOp(ctx, lq.ctx, "All")
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*License, *LicenseQuery]()
	return withInterceptors[[]*License](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LicenseQuery) AllX(ctx context.Context) []*License {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of License IDs.
func (lq *LicenseQuery) IDs(ctx context.Context) (ids []int, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, "IDs")
	if err = lq.Select(license.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LicenseQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LicenseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, "Count")
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LicenseQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LicenseQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LicenseQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, "Exist")
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LicenseQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LicenseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LicenseQuery) Clone() *LicenseQuery {
	if lq == nil {
		return nil
	}
	return &LicenseQuery{
		config:     lq.config,
		ctx:        lq.ctx.Clone(),
		order:      append([]OrderFunc{}, lq.order...),
		inters:     append([]Interceptor{}, lq.inters...),
		predicates: append([]predicate.License{}, lq.predicates...),
		// clone intermediate query.
		gremlin: lq.gremlin.Clone(),
		path:    lq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.License.Query().
//		GroupBy(license.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LicenseQuery) GroupBy(field string, fields ...string) *LicenseGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LicenseGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = license.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.License.Query().
//		Select(license.FieldCreateTime).
//		Scan(ctx, &v)
func (lq *LicenseQuery) Select(fields ...string) *LicenseSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LicenseSelect{LicenseQuery: lq}
	sbuild.label = license.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LicenseSelect configured with the given aggregations.
func (lq *LicenseQuery) Aggregate(fns ...AggregateFunc) *LicenseSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LicenseQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.gremlin = prev
	}
	return nil
}

func (lq *LicenseQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*License, error) {
	res := &gremlin.Response{}
	traversal := lq.gremlinQuery(ctx)
	if len(lq.ctx.Fields) > 0 {
		fields := make([]any, len(lq.ctx.Fields))
		for i, f := range lq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := lq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var ls Licenses
	if err := ls.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range ls {
		ls[i].config = lq.config
	}
	return ls, nil
}

func (lq *LicenseQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := lq.gremlinQuery(ctx).Count().Query()
	if err := lq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (lq *LicenseQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(license.Label)
	if lq.gremlin != nil {
		v = lq.gremlin.Clone()
	}
	for _, p := range lq.predicates {
		p(v)
	}
	if len(lq.order) > 0 {
		v.Order()
		for _, p := range lq.order {
			p(v)
		}
	}
	switch limit, offset := lq.ctx.Limit, lq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := lq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// LicenseGroupBy is the group-by builder for License entities.
type LicenseGroupBy struct {
	selector
	build *LicenseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LicenseGroupBy) Aggregate(fns ...AggregateFunc) *LicenseGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LicenseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, "GroupBy")
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LicenseQuery, *LicenseGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LicenseGroupBy) gremlinScan(ctx context.Context, root *LicenseQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range lgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *lgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*lgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := lgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*lgb.flds)+len(lgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// LicenseSelect is the builder for selecting fields of License entities.
type LicenseSelect struct {
	*LicenseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LicenseSelect) Aggregate(fns ...AggregateFunc) *LicenseSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LicenseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, "Select")
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LicenseQuery, *LicenseSelect](ctx, ls.LicenseQuery, ls, ls.inters, v)
}

func (ls *LicenseSelect) gremlinScan(ctx context.Context, root *LicenseQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := ls.ctx.Fields; len(fields) == 1 {
		if fields[0] != license.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(ls.ctx.Fields))
		for i, f := range ls.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ls.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
