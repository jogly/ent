// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"
	"fmt"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entc/integration/migrate/entv1/customtype"
	"github.com/jogly/ent/entc/integration/migrate/entv1/predicate"
	"github.com/jogly/ent/schema/field"
)

// CustomTypeUpdate is the builder for updating CustomType entities.
type CustomTypeUpdate struct {
	config
	hooks    []Hook
	mutation *CustomTypeMutation
}

// Where appends a list predicates to the CustomTypeUpdate builder.
func (ctu *CustomTypeUpdate) Where(ps ...predicate.CustomType) *CustomTypeUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetCustom sets the "custom" field.
func (ctu *CustomTypeUpdate) SetCustom(s string) *CustomTypeUpdate {
	ctu.mutation.SetCustom(s)
	return ctu
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (ctu *CustomTypeUpdate) SetNillableCustom(s *string) *CustomTypeUpdate {
	if s != nil {
		ctu.SetCustom(*s)
	}
	return ctu
}

// ClearCustom clears the value of the "custom" field.
func (ctu *CustomTypeUpdate) ClearCustom() *CustomTypeUpdate {
	ctu.mutation.ClearCustom()
	return ctu
}

// Mutation returns the CustomTypeMutation object of the builder.
func (ctu *CustomTypeUpdate) Mutation() *CustomTypeMutation {
	return ctu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CustomTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, CustomTypeMutation](ctx, ctu.sqlSave, ctu.mutation, ctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CustomTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CustomTypeUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CustomTypeUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctu *CustomTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(customtype.Table, customtype.Columns, sqlgraph.NewFieldSpec(customtype.FieldID, field.TypeInt))
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.Custom(); ok {
		_spec.SetField(customtype.FieldCustom, field.TypeString, value)
	}
	if ctu.mutation.CustomCleared() {
		_spec.ClearField(customtype.FieldCustom, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ctu.mutation.done = true
	return n, nil
}

// CustomTypeUpdateOne is the builder for updating a single CustomType entity.
type CustomTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomTypeMutation
}

// SetCustom sets the "custom" field.
func (ctuo *CustomTypeUpdateOne) SetCustom(s string) *CustomTypeUpdateOne {
	ctuo.mutation.SetCustom(s)
	return ctuo
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (ctuo *CustomTypeUpdateOne) SetNillableCustom(s *string) *CustomTypeUpdateOne {
	if s != nil {
		ctuo.SetCustom(*s)
	}
	return ctuo
}

// ClearCustom clears the value of the "custom" field.
func (ctuo *CustomTypeUpdateOne) ClearCustom() *CustomTypeUpdateOne {
	ctuo.mutation.ClearCustom()
	return ctuo
}

// Mutation returns the CustomTypeMutation object of the builder.
func (ctuo *CustomTypeUpdateOne) Mutation() *CustomTypeMutation {
	return ctuo.mutation
}

// Where appends a list predicates to the CustomTypeUpdate builder.
func (ctuo *CustomTypeUpdateOne) Where(ps ...predicate.CustomType) *CustomTypeUpdateOne {
	ctuo.mutation.Where(ps...)
	return ctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CustomTypeUpdateOne) Select(field string, fields ...string) *CustomTypeUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CustomType entity.
func (ctuo *CustomTypeUpdateOne) Save(ctx context.Context) (*CustomType, error) {
	return withHooks[*CustomType, CustomTypeMutation](ctx, ctuo.sqlSave, ctuo.mutation, ctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CustomTypeUpdateOne) SaveX(ctx context.Context) *CustomType {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CustomTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CustomTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctuo *CustomTypeUpdateOne) sqlSave(ctx context.Context) (_node *CustomType, err error) {
	_spec := sqlgraph.NewUpdateSpec(customtype.Table, customtype.Columns, sqlgraph.NewFieldSpec(customtype.FieldID, field.TypeInt))
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv1: missing "CustomType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customtype.FieldID)
		for _, f := range fields {
			if !customtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv1: invalid field %q for query", f)}
			}
			if f != customtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.Custom(); ok {
		_spec.SetField(customtype.FieldCustom, field.TypeString, value)
	}
	if ctuo.mutation.CustomCleared() {
		_spec.ClearField(customtype.FieldCustom, field.TypeString)
	}
	_node = &CustomType{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ctuo.mutation.done = true
	return _node, nil
}
