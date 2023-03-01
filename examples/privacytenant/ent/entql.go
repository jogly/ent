// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/jogly/ent/examples/privacytenant/ent/group"
	"github.com/jogly/ent/examples/privacytenant/ent/predicate"
	"github.com/jogly/ent/examples/privacytenant/ent/tenant"
	"github.com/jogly/ent/examples/privacytenant/ent/user"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entql"
	"github.com/jogly/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   group.Table,
			Columns: group.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		},
		Type: "Group",
		Fields: map[string]*sqlgraph.FieldSpec{
			group.FieldTenantID: {Type: field.TypeInt, Column: group.FieldTenantID},
			group.FieldName:     {Type: field.TypeString, Column: group.FieldName},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tenant.Table,
			Columns: tenant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tenant.FieldID,
			},
		},
		Type: "Tenant",
		Fields: map[string]*sqlgraph.FieldSpec{
			tenant.FieldName: {Type: field.TypeString, Column: tenant.FieldName},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldTenantID: {Type: field.TypeInt, Column: user.FieldTenantID},
			user.FieldName:     {Type: field.TypeString, Column: user.FieldName},
			user.FieldFoods:    {Type: field.TypeJSON, Column: user.FieldFoods},
		},
	}
	graph.MustAddE(
		"tenant",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   group.TenantTable,
			Columns: []string{group.TenantColumn},
			Bidi:    false,
		},
		"Group",
		"Tenant",
	)
	graph.MustAddE(
		"users",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
		},
		"Group",
		"User",
	)
	graph.MustAddE(
		"tenant",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.TenantTable,
			Columns: []string{user.TenantColumn},
			Bidi:    false,
		},
		"User",
		"Tenant",
	)
	graph.MustAddE(
		"groups",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
		},
		"User",
		"Group",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (gq *GroupQuery) addPredicate(pred func(s *sql.Selector)) {
	gq.predicates = append(gq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GroupQuery builder.
func (gq *GroupQuery) Filter() *GroupFilter {
	return &GroupFilter{config: gq.config, predicateAdder: gq}
}

// addPredicate implements the predicateAdder interface.
func (m *GroupMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GroupMutation builder.
func (m *GroupMutation) Filter() *GroupFilter {
	return &GroupFilter{config: m.config, predicateAdder: m}
}

// GroupFilter provides a generic filtering capability at runtime for GroupQuery.
type GroupFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GroupFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *GroupFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(group.FieldID))
}

// WhereTenantID applies the entql int predicate on the tenant_id field.
func (f *GroupFilter) WhereTenantID(p entql.IntP) {
	f.Where(p.Field(group.FieldTenantID))
}

// WhereName applies the entql string predicate on the name field.
func (f *GroupFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(group.FieldName))
}

// WhereHasTenant applies a predicate to check if query has an edge tenant.
func (f *GroupFilter) WhereHasTenant() {
	f.Where(entql.HasEdge("tenant"))
}

// WhereHasTenantWith applies a predicate to check if query has an edge tenant with a given conditions (other predicates).
func (f *GroupFilter) WhereHasTenantWith(preds ...predicate.Tenant) {
	f.Where(entql.HasEdgeWith("tenant", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasUsers applies a predicate to check if query has an edge users.
func (f *GroupFilter) WhereHasUsers() {
	f.Where(entql.HasEdge("users"))
}

// WhereHasUsersWith applies a predicate to check if query has an edge users with a given conditions (other predicates).
func (f *GroupFilter) WhereHasUsersWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("users", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (tq *TenantQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TenantQuery builder.
func (tq *TenantQuery) Filter() *TenantFilter {
	return &TenantFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TenantMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TenantMutation builder.
func (m *TenantMutation) Filter() *TenantFilter {
	return &TenantFilter{config: m.config, predicateAdder: m}
}

// TenantFilter provides a generic filtering capability at runtime for TenantQuery.
type TenantFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TenantFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TenantFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(tenant.FieldID))
}

// WhereName applies the entql string predicate on the name field.
func (f *TenantFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(tenant.FieldName))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *UserFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(user.FieldID))
}

// WhereTenantID applies the entql int predicate on the tenant_id field.
func (f *UserFilter) WhereTenantID(p entql.IntP) {
	f.Where(p.Field(user.FieldTenantID))
}

// WhereName applies the entql string predicate on the name field.
func (f *UserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WhereFoods applies the entql json.RawMessage predicate on the foods field.
func (f *UserFilter) WhereFoods(p entql.BytesP) {
	f.Where(p.Field(user.FieldFoods))
}

// WhereHasTenant applies a predicate to check if query has an edge tenant.
func (f *UserFilter) WhereHasTenant() {
	f.Where(entql.HasEdge("tenant"))
}

// WhereHasTenantWith applies a predicate to check if query has an edge tenant with a given conditions (other predicates).
func (f *UserFilter) WhereHasTenantWith(preds ...predicate.Tenant) {
	f.Where(entql.HasEdgeWith("tenant", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasGroups applies a predicate to check if query has an edge groups.
func (f *UserFilter) WhereHasGroups() {
	f.Where(entql.HasEdge("groups"))
}

// WhereHasGroupsWith applies a predicate to check if query has an edge groups with a given conditions (other predicates).
func (f *UserFilter) WhereHasGroupsWith(preds ...predicate.Group) {
	f.Where(entql.HasEdgeWith("groups", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
