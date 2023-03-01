// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/jogly/ent/entc/integration/customid/ent"
	"github.com/jogly/ent/entql"

	"github.com/jogly/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AccountQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AccountQueryRuleFunc func(context.Context, *ent.AccountQuery) error

// EvalQuery return f(ctx, q).
func (f AccountQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AccountQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AccountQuery", q)
}

// The AccountMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AccountMutationRuleFunc func(context.Context, *ent.AccountMutation) error

// EvalMutation calls f(ctx, m).
func (f AccountMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AccountMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AccountMutation", m)
}

// The BlobQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BlobQueryRuleFunc func(context.Context, *ent.BlobQuery) error

// EvalQuery return f(ctx, q).
func (f BlobQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BlobQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BlobQuery", q)
}

// The BlobMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BlobMutationRuleFunc func(context.Context, *ent.BlobMutation) error

// EvalMutation calls f(ctx, m).
func (f BlobMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BlobMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BlobMutation", m)
}

// The BlobLinkQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type BlobLinkQueryRuleFunc func(context.Context, *ent.BlobLinkQuery) error

// EvalQuery return f(ctx, q).
func (f BlobLinkQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BlobLinkQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.BlobLinkQuery", q)
}

// The BlobLinkMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type BlobLinkMutationRuleFunc func(context.Context, *ent.BlobLinkMutation) error

// EvalMutation calls f(ctx, m).
func (f BlobLinkMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.BlobLinkMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.BlobLinkMutation", m)
}

// The CarQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CarQueryRuleFunc func(context.Context, *ent.CarQuery) error

// EvalQuery return f(ctx, q).
func (f CarQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CarQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CarQuery", q)
}

// The CarMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CarMutationRuleFunc func(context.Context, *ent.CarMutation) error

// EvalMutation calls f(ctx, m).
func (f CarMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CarMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CarMutation", m)
}

// The DeviceQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DeviceQueryRuleFunc func(context.Context, *ent.DeviceQuery) error

// EvalQuery return f(ctx, q).
func (f DeviceQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DeviceQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DeviceQuery", q)
}

// The DeviceMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DeviceMutationRuleFunc func(context.Context, *ent.DeviceMutation) error

// EvalMutation calls f(ctx, m).
func (f DeviceMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DeviceMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DeviceMutation", m)
}

// The DocQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DocQueryRuleFunc func(context.Context, *ent.DocQuery) error

// EvalQuery return f(ctx, q).
func (f DocQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DocQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DocQuery", q)
}

// The DocMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DocMutationRuleFunc func(context.Context, *ent.DocMutation) error

// EvalMutation calls f(ctx, m).
func (f DocMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DocMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DocMutation", m)
}

// The GroupQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupQueryRuleFunc func(context.Context, *ent.GroupQuery) error

// EvalQuery return f(ctx, q).
func (f GroupQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GroupQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GroupQuery", q)
}

// The GroupMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupMutationRuleFunc func(context.Context, *ent.GroupMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GroupMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GroupMutation", m)
}

// The IntSIDQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type IntSIDQueryRuleFunc func(context.Context, *ent.IntSIDQuery) error

// EvalQuery return f(ctx, q).
func (f IntSIDQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.IntSIDQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.IntSIDQuery", q)
}

// The IntSIDMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type IntSIDMutationRuleFunc func(context.Context, *ent.IntSIDMutation) error

// EvalMutation calls f(ctx, m).
func (f IntSIDMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.IntSIDMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.IntSIDMutation", m)
}

// The LinkQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LinkQueryRuleFunc func(context.Context, *ent.LinkQuery) error

// EvalQuery return f(ctx, q).
func (f LinkQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LinkQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LinkQuery", q)
}

// The LinkMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LinkMutationRuleFunc func(context.Context, *ent.LinkMutation) error

// EvalMutation calls f(ctx, m).
func (f LinkMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LinkMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LinkMutation", m)
}

// The MixinIDQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MixinIDQueryRuleFunc func(context.Context, *ent.MixinIDQuery) error

// EvalQuery return f(ctx, q).
func (f MixinIDQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MixinIDQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MixinIDQuery", q)
}

// The MixinIDMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MixinIDMutationRuleFunc func(context.Context, *ent.MixinIDMutation) error

// EvalMutation calls f(ctx, m).
func (f MixinIDMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MixinIDMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MixinIDMutation", m)
}

// The NoteQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type NoteQueryRuleFunc func(context.Context, *ent.NoteQuery) error

// EvalQuery return f(ctx, q).
func (f NoteQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NoteQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.NoteQuery", q)
}

// The NoteMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type NoteMutationRuleFunc func(context.Context, *ent.NoteMutation) error

// EvalMutation calls f(ctx, m).
func (f NoteMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.NoteMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.NoteMutation", m)
}

// The OtherQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OtherQueryRuleFunc func(context.Context, *ent.OtherQuery) error

// EvalQuery return f(ctx, q).
func (f OtherQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OtherQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.OtherQuery", q)
}

// The OtherMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OtherMutationRuleFunc func(context.Context, *ent.OtherMutation) error

// EvalMutation calls f(ctx, m).
func (f OtherMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.OtherMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.OtherMutation", m)
}

// The PetQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PetQueryRuleFunc func(context.Context, *ent.PetQuery) error

// EvalQuery return f(ctx, q).
func (f PetQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PetQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PetQuery", q)
}

// The PetMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PetMutationRuleFunc func(context.Context, *ent.PetMutation) error

// EvalMutation calls f(ctx, m).
func (f PetMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PetMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PetMutation", m)
}

// The RevisionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RevisionQueryRuleFunc func(context.Context, *ent.RevisionQuery) error

// EvalQuery return f(ctx, q).
func (f RevisionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RevisionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RevisionQuery", q)
}

// The RevisionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RevisionMutationRuleFunc func(context.Context, *ent.RevisionMutation) error

// EvalMutation calls f(ctx, m).
func (f RevisionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RevisionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RevisionMutation", m)
}

// The SessionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SessionQueryRuleFunc func(context.Context, *ent.SessionQuery) error

// EvalQuery return f(ctx, q).
func (f SessionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SessionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SessionQuery", q)
}

// The SessionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SessionMutationRuleFunc func(context.Context, *ent.SessionMutation) error

// EvalMutation calls f(ctx, m).
func (f SessionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SessionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SessionMutation", m)
}

// The TokenQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TokenQueryRuleFunc func(context.Context, *ent.TokenQuery) error

// EvalQuery return f(ctx, q).
func (f TokenQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TokenQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TokenQuery", q)
}

// The TokenMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TokenMutationRuleFunc func(context.Context, *ent.TokenMutation) error

// EvalMutation calls f(ctx, m).
func (f TokenMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TokenMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TokenMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *ent.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *ent.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.UserMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AccountQuery:
		return q.Filter(), nil
	case *ent.BlobQuery:
		return q.Filter(), nil
	case *ent.BlobLinkQuery:
		return q.Filter(), nil
	case *ent.CarQuery:
		return q.Filter(), nil
	case *ent.DeviceQuery:
		return q.Filter(), nil
	case *ent.DocQuery:
		return q.Filter(), nil
	case *ent.GroupQuery:
		return q.Filter(), nil
	case *ent.IntSIDQuery:
		return q.Filter(), nil
	case *ent.LinkQuery:
		return q.Filter(), nil
	case *ent.MixinIDQuery:
		return q.Filter(), nil
	case *ent.NoteQuery:
		return q.Filter(), nil
	case *ent.OtherQuery:
		return q.Filter(), nil
	case *ent.PetQuery:
		return q.Filter(), nil
	case *ent.RevisionQuery:
		return q.Filter(), nil
	case *ent.SessionQuery:
		return q.Filter(), nil
	case *ent.TokenQuery:
		return q.Filter(), nil
	case *ent.UserQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AccountMutation:
		return m.Filter(), nil
	case *ent.BlobMutation:
		return m.Filter(), nil
	case *ent.BlobLinkMutation:
		return m.Filter(), nil
	case *ent.CarMutation:
		return m.Filter(), nil
	case *ent.DeviceMutation:
		return m.Filter(), nil
	case *ent.DocMutation:
		return m.Filter(), nil
	case *ent.GroupMutation:
		return m.Filter(), nil
	case *ent.IntSIDMutation:
		return m.Filter(), nil
	case *ent.LinkMutation:
		return m.Filter(), nil
	case *ent.MixinIDMutation:
		return m.Filter(), nil
	case *ent.NoteMutation:
		return m.Filter(), nil
	case *ent.OtherMutation:
		return m.Filter(), nil
	case *ent.PetMutation:
		return m.Filter(), nil
	case *ent.RevisionMutation:
		return m.Filter(), nil
	case *ent.SessionMutation:
		return m.Filter(), nil
	case *ent.TokenMutation:
		return m.Filter(), nil
	case *ent.UserMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
