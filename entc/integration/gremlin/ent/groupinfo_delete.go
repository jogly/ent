// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/jogly/ent/dialect/gremlin"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl/__"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl/g"
	"github.com/jogly/ent/entc/integration/gremlin/ent/groupinfo"
	"github.com/jogly/ent/entc/integration/gremlin/ent/predicate"
)

// GroupInfoDelete is the builder for deleting a GroupInfo entity.
type GroupInfoDelete struct {
	config
	hooks    []Hook
	mutation *GroupInfoMutation
}

// Where appends a list predicates to the GroupInfoDelete builder.
func (gid *GroupInfoDelete) Where(ps ...predicate.GroupInfo) *GroupInfoDelete {
	gid.mutation.Where(ps...)
	return gid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gid *GroupInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GroupInfoMutation](ctx, gid.gremlinExec, gid.mutation, gid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gid *GroupInfoDelete) ExecX(ctx context.Context) int {
	n, err := gid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gid *GroupInfoDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gid.gremlin().Query()
	if err := gid.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	gid.mutation.done = true
	return res.ReadInt()
}

func (gid *GroupInfoDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(groupinfo.Label)
	for _, p := range gid.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// GroupInfoDeleteOne is the builder for deleting a single GroupInfo entity.
type GroupInfoDeleteOne struct {
	gid *GroupInfoDelete
}

// Where appends a list predicates to the GroupInfoDelete builder.
func (gido *GroupInfoDeleteOne) Where(ps ...predicate.GroupInfo) *GroupInfoDeleteOne {
	gido.gid.mutation.Where(ps...)
	return gido
}

// Exec executes the deletion query.
func (gido *GroupInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := gido.gid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gido *GroupInfoDeleteOne) ExecX(ctx context.Context) {
	if err := gido.Exec(ctx); err != nil {
		panic(err)
	}
}
