// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package item

import (
	"github.com/jogly/ent/dialect/gremlin/graph/dsl"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl/__"
	"github.com/jogly/ent/dialect/gremlin/graph/dsl/p"
	"github.com/jogly/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.EQ(v))
	})
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.EQ(v))
	})
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.NEQ(v))
	})
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.Within(vs...))
	})
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.Without(vs...))
	})
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.GT(v))
	})
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.GTE(v))
	})
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.LT(v))
	})
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.LTE(v))
	})
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.Containing(v))
	})
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.StartingWith(v))
	})
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.Has(Label, FieldText, p.EndingWith(v))
	})
}

// TextIsNil applies the IsNil predicate on the "text" field.
func TextIsNil() predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldText)
	})
}

// TextNotNil applies the NotNil predicate on the "text" field.
func TextNotNil() predicate.Item {
	return predicate.Item(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldText)
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(func(tr *dsl.Traversal) {
		trs := make([]any, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Item) predicate.Item {
	return predicate.Item(func(tr *dsl.Traversal) {
		trs := make([]any, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Item) predicate.Item {
	return predicate.Item(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
