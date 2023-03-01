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
	"github.com/jogly/ent/entc/integration/gremlin/ent/node"
	"github.com/jogly/ent/entc/integration/gremlin/ent/predicate"
)

// NodeQuery is the builder for querying Node entities.
type NodeQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Node
	withPrev   *NodeQuery
	withNext   *NodeQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the NodeQuery builder.
func (nq *NodeQuery) Where(ps ...predicate.Node) *NodeQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NodeQuery) Limit(limit int) *NodeQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NodeQuery) Offset(offset int) *NodeQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NodeQuery) Unique(unique bool) *NodeQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NodeQuery) Order(o ...OrderFunc) *NodeQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryPrev chains the current query on the "prev" edge.
func (nq *NodeQuery) QueryPrev() *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := nq.gremlinQuery(ctx)
		fromU = gremlin.InE(node.NextLabel).OutV()
		return fromU, nil
	}
	return query
}

// QueryNext chains the current query on the "next" edge.
func (nq *NodeQuery) QueryNext() *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := nq.gremlinQuery(ctx)
		fromU = gremlin.OutE(node.NextLabel).InV()
		return fromU, nil
	}
	return query
}

// First returns the first Node entity from the query.
// Returns a *NotFoundError when no Node was found.
func (nq *NodeQuery) First(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{node.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NodeQuery) FirstX(ctx context.Context) *Node {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Node ID from the query.
// Returns a *NotFoundError when no Node ID was found.
func (nq *NodeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{node.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NodeQuery) FirstIDX(ctx context.Context) string {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Node entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Node entity is found.
// Returns a *NotFoundError when no Node entities are found.
func (nq *NodeQuery) Only(ctx context.Context) (*Node, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{node.Label}
	default:
		return nil, &NotSingularError{node.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NodeQuery) OnlyX(ctx context.Context) *Node {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Node ID in the query.
// Returns a *NotSingularError when more than one Node ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NodeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{node.Label}
	default:
		err = &NotSingularError{node.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NodeQuery) OnlyIDX(ctx context.Context) string {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Nodes.
func (nq *NodeQuery) All(ctx context.Context) ([]*Node, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Node, *NodeQuery]()
	return withInterceptors[[]*Node](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NodeQuery) AllX(ctx context.Context) []*Node {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Node IDs.
func (nq *NodeQuery) IDs(ctx context.Context) (ids []string, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(node.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NodeQuery) IDsX(ctx context.Context) []string {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NodeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NodeQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NodeQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NodeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NodeQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NodeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NodeQuery) Clone() *NodeQuery {
	if nq == nil {
		return nil
	}
	return &NodeQuery{
		config:     nq.config,
		ctx:        nq.ctx.Clone(),
		order:      append([]OrderFunc{}, nq.order...),
		inters:     append([]Interceptor{}, nq.inters...),
		predicates: append([]predicate.Node{}, nq.predicates...),
		withPrev:   nq.withPrev.Clone(),
		withNext:   nq.withNext.Clone(),
		// clone intermediate query.
		gremlin: nq.gremlin.Clone(),
		path:    nq.path,
	}
}

// WithPrev tells the query-builder to eager-load the nodes that are connected to
// the "prev" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithPrev(opts ...func(*NodeQuery)) *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withPrev = query
	return nq
}

// WithNext tells the query-builder to eager-load the nodes that are connected to
// the "next" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NodeQuery) WithNext(opts ...func(*NodeQuery)) *NodeQuery {
	query := (&NodeClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withNext = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Value int `json:"value,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Node.Query().
//		GroupBy(node.FieldValue).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NodeQuery) GroupBy(field string, fields ...string) *NodeGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NodeGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = node.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Value int `json:"value,omitempty"`
//	}
//
//	client.Node.Query().
//		Select(node.FieldValue).
//		Scan(ctx, &v)
func (nq *NodeQuery) Select(fields ...string) *NodeSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NodeSelect{NodeQuery: nq}
	sbuild.label = node.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NodeSelect configured with the given aggregations.
func (nq *NodeQuery) Aggregate(fns ...AggregateFunc) *NodeSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NodeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.gremlin = prev
	}
	return nil
}

func (nq *NodeQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Node, error) {
	res := &gremlin.Response{}
	traversal := nq.gremlinQuery(ctx)
	if len(nq.ctx.Fields) > 0 {
		fields := make([]any, len(nq.ctx.Fields))
		for i, f := range nq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := nq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var ns Nodes
	if err := ns.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range ns {
		ns[i].config = nq.config
	}
	return ns, nil
}

func (nq *NodeQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := nq.gremlinQuery(ctx).Count().Query()
	if err := nq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (nq *NodeQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(node.Label)
	if nq.gremlin != nil {
		v = nq.gremlin.Clone()
	}
	for _, p := range nq.predicates {
		p(v)
	}
	if len(nq.order) > 0 {
		v.Order()
		for _, p := range nq.order {
			p(v)
		}
	}
	switch limit, offset := nq.ctx.Limit, nq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := nq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// NodeGroupBy is the group-by builder for Node entities.
type NodeGroupBy struct {
	selector
	build *NodeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NodeGroupBy) Aggregate(fns ...AggregateFunc) *NodeGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NodeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeQuery, *NodeGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NodeGroupBy) gremlinScan(ctx context.Context, root *NodeQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range ngb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *ngb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*ngb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := ngb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*ngb.flds)+len(ngb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// NodeSelect is the builder for selecting fields of Node entities.
type NodeSelect struct {
	*NodeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NodeSelect) Aggregate(fns ...AggregateFunc) *NodeSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NodeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NodeQuery, *NodeSelect](ctx, ns.NodeQuery, ns, ns.inters, v)
}

func (ns *NodeSelect) gremlinScan(ctx context.Context, root *NodeQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := ns.ctx.Fields; len(fields) == 1 {
		if fields[0] != node.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(ns.ctx.Fields))
		for i, f := range ns.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ns.driver.Exec(ctx, query, bindings, res); err != nil {
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
