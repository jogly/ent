// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/predicate"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/usergroup"
	"github.com/jogly/ent/schema/field"
)

// UserGroupDelete is the builder for deleting a UserGroup entity.
type UserGroupDelete struct {
	config
	hooks    []Hook
	mutation *UserGroupMutation
}

// Where appends a list predicates to the UserGroupDelete builder.
func (ugd *UserGroupDelete) Where(ps ...predicate.UserGroup) *UserGroupDelete {
	ugd.mutation.Where(ps...)
	return ugd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ugd *UserGroupDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UserGroupMutation](ctx, ugd.sqlExec, ugd.mutation, ugd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ugd *UserGroupDelete) ExecX(ctx context.Context) int {
	n, err := ugd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ugd *UserGroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usergroup.Table, sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeInt))
	if ps := ugd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ugd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ugd.mutation.done = true
	return affected, err
}

// UserGroupDeleteOne is the builder for deleting a single UserGroup entity.
type UserGroupDeleteOne struct {
	ugd *UserGroupDelete
}

// Where appends a list predicates to the UserGroupDelete builder.
func (ugdo *UserGroupDeleteOne) Where(ps ...predicate.UserGroup) *UserGroupDeleteOne {
	ugdo.ugd.mutation.Where(ps...)
	return ugdo
}

// Exec executes the deletion query.
func (ugdo *UserGroupDeleteOne) Exec(ctx context.Context) error {
	n, err := ugdo.ugd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usergroup.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ugdo *UserGroupDeleteOne) ExecX(ctx context.Context) {
	if err := ugdo.Exec(ctx); err != nil {
		panic(err)
	}
}
