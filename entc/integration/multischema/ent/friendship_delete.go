// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entc/integration/multischema/ent/friendship"
	"github.com/jogly/ent/entc/integration/multischema/ent/internal"
	"github.com/jogly/ent/entc/integration/multischema/ent/predicate"
	"github.com/jogly/ent/schema/field"
)

// FriendshipDelete is the builder for deleting a Friendship entity.
type FriendshipDelete struct {
	config
	hooks    []Hook
	mutation *FriendshipMutation
}

// Where appends a list predicates to the FriendshipDelete builder.
func (fd *FriendshipDelete) Where(ps ...predicate.Friendship) *FriendshipDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FriendshipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, FriendshipMutation](ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FriendshipDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FriendshipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(friendship.Table, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	_spec.Node.Schema = fd.schemaConfig.Friendship
	ctx = internal.NewSchemaConfigContext(ctx, fd.schemaConfig)
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FriendshipDeleteOne is the builder for deleting a single Friendship entity.
type FriendshipDeleteOne struct {
	fd *FriendshipDelete
}

// Where appends a list predicates to the FriendshipDelete builder.
func (fdo *FriendshipDeleteOne) Where(ps ...predicate.Friendship) *FriendshipDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FriendshipDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{friendship.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FriendshipDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}
