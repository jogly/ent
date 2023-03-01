// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entc/integration/multischema/ent/friendship"
	"github.com/jogly/ent/entc/integration/multischema/ent/internal"
	"github.com/jogly/ent/entc/integration/multischema/ent/predicate"
	"github.com/jogly/ent/schema/field"
)

// FriendshipUpdate is the builder for updating Friendship entities.
type FriendshipUpdate struct {
	config
	hooks     []Hook
	mutation  *FriendshipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (fu *FriendshipUpdate) Where(ps ...predicate.Friendship) *FriendshipUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetWeight sets the "weight" field.
func (fu *FriendshipUpdate) SetWeight(i int) *FriendshipUpdate {
	fu.mutation.ResetWeight()
	fu.mutation.SetWeight(i)
	return fu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableWeight(i *int) *FriendshipUpdate {
	if i != nil {
		fu.SetWeight(*i)
	}
	return fu
}

// AddWeight adds i to the "weight" field.
func (fu *FriendshipUpdate) AddWeight(i int) *FriendshipUpdate {
	fu.mutation.AddWeight(i)
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *FriendshipUpdate) SetCreatedAt(t time.Time) *FriendshipUpdate {
	fu.mutation.SetCreatedAt(t)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableCreatedAt(t *time.Time) *FriendshipUpdate {
	if t != nil {
		fu.SetCreatedAt(*t)
	}
	return fu
}

// Mutation returns the FriendshipMutation object of the builder.
func (fu *FriendshipUpdate) Mutation() *FriendshipMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FriendshipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, FriendshipMutation](ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FriendshipUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FriendshipUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FriendshipUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FriendshipUpdate) check() error {
	if _, ok := fu.mutation.UserID(); fu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.user"`)
	}
	if _, ok := fu.mutation.FriendID(); fu.mutation.FriendCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.friend"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fu *FriendshipUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FriendshipUpdate {
	fu.modifiers = append(fu.modifiers, modifiers...)
	return fu
}

func (fu *FriendshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Weight(); ok {
		_spec.SetField(friendship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedWeight(); ok {
		_spec.AddField(friendship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
	}
	_spec.Node.Schema = fu.schemaConfig.Friendship
	ctx = internal.NewSchemaConfigContext(ctx, fu.schemaConfig)
	_spec.AddModifiers(fu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FriendshipUpdateOne is the builder for updating a single Friendship entity.
type FriendshipUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FriendshipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetWeight sets the "weight" field.
func (fuo *FriendshipUpdateOne) SetWeight(i int) *FriendshipUpdateOne {
	fuo.mutation.ResetWeight()
	fuo.mutation.SetWeight(i)
	return fuo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableWeight(i *int) *FriendshipUpdateOne {
	if i != nil {
		fuo.SetWeight(*i)
	}
	return fuo
}

// AddWeight adds i to the "weight" field.
func (fuo *FriendshipUpdateOne) AddWeight(i int) *FriendshipUpdateOne {
	fuo.mutation.AddWeight(i)
	return fuo
}

// SetCreatedAt sets the "created_at" field.
func (fuo *FriendshipUpdateOne) SetCreatedAt(t time.Time) *FriendshipUpdateOne {
	fuo.mutation.SetCreatedAt(t)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableCreatedAt(t *time.Time) *FriendshipUpdateOne {
	if t != nil {
		fuo.SetCreatedAt(*t)
	}
	return fuo
}

// Mutation returns the FriendshipMutation object of the builder.
func (fuo *FriendshipUpdateOne) Mutation() *FriendshipMutation {
	return fuo.mutation
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (fuo *FriendshipUpdateOne) Where(ps ...predicate.Friendship) *FriendshipUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FriendshipUpdateOne) Select(field string, fields ...string) *FriendshipUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Friendship entity.
func (fuo *FriendshipUpdateOne) Save(ctx context.Context) (*Friendship, error) {
	return withHooks[*Friendship, FriendshipMutation](ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FriendshipUpdateOne) SaveX(ctx context.Context) *Friendship {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FriendshipUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FriendshipUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FriendshipUpdateOne) check() error {
	if _, ok := fuo.mutation.UserID(); fuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.user"`)
	}
	if _, ok := fuo.mutation.FriendID(); fuo.mutation.FriendCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.friend"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fuo *FriendshipUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FriendshipUpdateOne {
	fuo.modifiers = append(fuo.modifiers, modifiers...)
	return fuo
}

func (fuo *FriendshipUpdateOne) sqlSave(ctx context.Context) (_node *Friendship, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friendship.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendship.FieldID)
		for _, f := range fields {
			if !friendship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friendship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Weight(); ok {
		_spec.SetField(friendship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedWeight(); ok {
		_spec.AddField(friendship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
	}
	_spec.Node.Schema = fuo.schemaConfig.Friendship
	ctx = internal.NewSchemaConfigContext(ctx, fuo.schemaConfig)
	_spec.AddModifiers(fuo.modifiers...)
	_node = &Friendship{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
