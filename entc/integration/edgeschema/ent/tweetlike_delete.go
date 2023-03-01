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
	"github.com/jogly/ent/entc/integration/edgeschema/ent/tweetlike"
)

// TweetLikeDelete is the builder for deleting a TweetLike entity.
type TweetLikeDelete struct {
	config
	hooks    []Hook
	mutation *TweetLikeMutation
}

// Where appends a list predicates to the TweetLikeDelete builder.
func (tld *TweetLikeDelete) Where(ps ...predicate.TweetLike) *TweetLikeDelete {
	tld.mutation.Where(ps...)
	return tld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tld *TweetLikeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TweetLikeMutation](ctx, tld.sqlExec, tld.mutation, tld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tld *TweetLikeDelete) ExecX(ctx context.Context) int {
	n, err := tld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tld *TweetLikeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tweetlike.Table, nil)
	if ps := tld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tld.mutation.done = true
	return affected, err
}

// TweetLikeDeleteOne is the builder for deleting a single TweetLike entity.
type TweetLikeDeleteOne struct {
	tld *TweetLikeDelete
}

// Where appends a list predicates to the TweetLikeDelete builder.
func (tldo *TweetLikeDeleteOne) Where(ps ...predicate.TweetLike) *TweetLikeDeleteOne {
	tldo.tld.mutation.Where(ps...)
	return tldo
}

// Exec executes the deletion query.
func (tldo *TweetLikeDeleteOne) Exec(ctx context.Context) error {
	n, err := tldo.tld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tweetlike.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tldo *TweetLikeDeleteOne) ExecX(ctx context.Context) {
	if err := tldo.Exec(ctx); err != nil {
		panic(err)
	}
}
