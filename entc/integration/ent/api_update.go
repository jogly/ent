// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entc/integration/ent/api"
	"github.com/jogly/ent/entc/integration/ent/predicate"
	"github.com/jogly/ent/schema/field"
)

// APIUpdate is the builder for updating Api entities.
type APIUpdate struct {
	config
	hooks     []Hook
	mutation  *APIMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the APIUpdate builder.
func (au *APIUpdate) Where(ps ...predicate.Api) *APIUpdate {
	au.mutation.Where(ps...)
	return au
}

// Mutation returns the APIMutation object of the builder.
func (au *APIUpdate) Mutation() *APIMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *APIUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, APIMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *APIUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *APIUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *APIUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *APIUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *APIUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *APIUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(api.Table, api.Columns, sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{api.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// APIUpdateOne is the builder for updating a single Api entity.
type APIUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *APIMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Mutation returns the APIMutation object of the builder.
func (auo *APIUpdateOne) Mutation() *APIMutation {
	return auo.mutation
}

// Where appends a list predicates to the APIUpdate builder.
func (auo *APIUpdateOne) Where(ps ...predicate.Api) *APIUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *APIUpdateOne) Select(field string, fields ...string) *APIUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Api entity.
func (auo *APIUpdateOne) Save(ctx context.Context) (*Api, error) {
	return withHooks[*Api, APIMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *APIUpdateOne) SaveX(ctx context.Context) *Api {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *APIUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *APIUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *APIUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *APIUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *APIUpdateOne) sqlSave(ctx context.Context) (_node *Api, err error) {
	_spec := sqlgraph.NewUpdateSpec(api.Table, api.Columns, sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Api.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, api.FieldID)
		for _, f := range fields {
			if !api.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != api.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Api{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{api.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
