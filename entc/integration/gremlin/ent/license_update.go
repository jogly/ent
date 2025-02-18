// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/jogly/ent/dialect/gremlin"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl/g"
	"github.com/jogly/ent/entc/integration/gremlin/ent/license"
	"github.com/jogly/ent/entc/integration/gremlin/ent/predicate"
)

// LicenseUpdate is the builder for updating License entities.
type LicenseUpdate struct {
	config
	hooks    []Hook
	mutation *LicenseMutation
}

// Where appends a list predicates to the LicenseUpdate builder.
func (lu *LicenseUpdate) Where(ps ...predicate.License) *LicenseUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdateTime sets the "update_time" field.
func (lu *LicenseUpdate) SetUpdateTime(t time.Time) *LicenseUpdate {
	lu.mutation.SetUpdateTime(t)
	return lu
}

// Mutation returns the LicenseMutation object of the builder.
func (lu *LicenseUpdate) Mutation() *LicenseMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LicenseUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks[int, LicenseMutation](ctx, lu.gremlinSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LicenseUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LicenseUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LicenseUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LicenseUpdate) defaults() {
	if _, ok := lu.mutation.UpdateTime(); !ok {
		v := license.UpdateDefaultUpdateTime()
		lu.mutation.SetUpdateTime(v)
	}
}

func (lu *LicenseUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := lu.gremlin().Query()
	if err := lu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	lu.mutation.done = true
	return res.ReadInt()
}

func (lu *LicenseUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(license.Label)
	for _, p := range lu.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if value, ok := lu.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, license.FieldUpdateTime, value)
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// LicenseUpdateOne is the builder for updating a single License entity.
type LicenseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LicenseMutation
}

// SetUpdateTime sets the "update_time" field.
func (luo *LicenseUpdateOne) SetUpdateTime(t time.Time) *LicenseUpdateOne {
	luo.mutation.SetUpdateTime(t)
	return luo
}

// Mutation returns the LicenseMutation object of the builder.
func (luo *LicenseUpdateOne) Mutation() *LicenseMutation {
	return luo.mutation
}

// Where appends a list predicates to the LicenseUpdate builder.
func (luo *LicenseUpdateOne) Where(ps ...predicate.License) *LicenseUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LicenseUpdateOne) Select(field string, fields ...string) *LicenseUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated License entity.
func (luo *LicenseUpdateOne) Save(ctx context.Context) (*License, error) {
	luo.defaults()
	return withHooks[*License, LicenseMutation](ctx, luo.gremlinSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LicenseUpdateOne) SaveX(ctx context.Context) *License {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LicenseUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LicenseUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LicenseUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdateTime(); !ok {
		v := license.UpdateDefaultUpdateTime()
		luo.mutation.SetUpdateTime(v)
	}
}

func (luo *LicenseUpdateOne) gremlinSave(ctx context.Context) (*License, error) {
	res := &gremlin.Response{}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "License.id" for update`)}
	}
	query, bindings := luo.gremlin(id).Query()
	if err := luo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	luo.mutation.done = true
	l := &License{config: luo.config}
	if err := l.FromResponse(res); err != nil {
		return nil, err
	}
	return l, nil
}

func (luo *LicenseUpdateOne) gremlin(id int) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if value, ok := luo.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, license.FieldUpdateTime, value)
	}
	if len(luo.fields) > 0 {
		fields := make([]any, 0, len(luo.fields)+1)
		fields = append(fields, true)
		for _, f := range luo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
