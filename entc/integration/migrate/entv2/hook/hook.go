// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/jogly/ent/entc/integration/migrate/entv2"
)

// The BlogFunc type is an adapter to allow the use of ordinary
// function as Blog mutator.
type BlogFunc func(context.Context, *entv2.BlogMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f BlogFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.BlogMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.BlogMutation", m)
}

// The CarFunc type is an adapter to allow the use of ordinary
// function as Car mutator.
type CarFunc func(context.Context, *entv2.CarMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f CarFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.CarMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.CarMutation", m)
}

// The ConversionFunc type is an adapter to allow the use of ordinary
// function as Conversion mutator.
type ConversionFunc func(context.Context, *entv2.ConversionMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f ConversionFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.ConversionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.ConversionMutation", m)
}

// The CustomTypeFunc type is an adapter to allow the use of ordinary
// function as CustomType mutator.
type CustomTypeFunc func(context.Context, *entv2.CustomTypeMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f CustomTypeFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.CustomTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.CustomTypeMutation", m)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *entv2.GroupMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.GroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.GroupMutation", m)
}

// The MediaFunc type is an adapter to allow the use of ordinary
// function as Media mutator.
type MediaFunc func(context.Context, *entv2.MediaMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f MediaFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.MediaMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.MediaMutation", m)
}

// The PetFunc type is an adapter to allow the use of ordinary
// function as Pet mutator.
type PetFunc func(context.Context, *entv2.PetMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f PetFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.PetMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.PetMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *entv2.UserMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.UserMutation", m)
}

// The ZooFunc type is an adapter to allow the use of ordinary
// function as Zoo mutator.
type ZooFunc func(context.Context, *entv2.ZooMutation) (entv2.Value, error)

// Mutate calls f(ctx, m).
func (f ZooFunc) Mutate(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
	if mv, ok := m.(*entv2.ZooMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *entv2.ZooMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, entv2.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entv2.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m entv2.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m entv2.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op entv2.Op) Condition {
	return func(_ context.Context, m entv2.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entv2.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entv2.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m entv2.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk entv2.Hook, cond Condition) entv2.Hook {
	return func(next entv2.Mutator) entv2.Mutator {
		return entv2.MutateFunc(func(ctx context.Context, m entv2.Mutation) (entv2.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, entv2.Delete|entv2.Create)
func On(hk entv2.Hook, op entv2.Op) entv2.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, entv2.Update|entv2.UpdateOne)
func Unless(hk entv2.Hook, op entv2.Op) entv2.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) entv2.Hook {
	return func(entv2.Mutator) entv2.Mutator {
		return entv2.MutateFunc(func(context.Context, entv2.Mutation) (entv2.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []entv2.Hook {
//		return []entv2.Hook{
//			Reject(entv2.Delete|entv2.Update),
//		}
//	}
func Reject(op entv2.Op) entv2.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []entv2.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...entv2.Hook) Chain {
	return Chain{append([]entv2.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() entv2.Hook {
	return func(mutator entv2.Mutator) entv2.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...entv2.Hook) Chain {
	newHooks := make([]entv2.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
