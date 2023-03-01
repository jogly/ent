// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/jogly/ent/dialect/sql"
	"github.com/jogly/ent/dialect/sql/sqlgraph"
	"github.com/jogly/ent/entc/integration/ent/pet"
	"github.com/jogly/ent/entc/integration/ent/predicate"
	"github.com/jogly/ent/entc/integration/ent/user"
	"github.com/jogly/ent/schema/field"
	"github.com/google/uuid"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	hooks     []Hook
	mutation  *PetMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PetUpdate builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAge sets the "age" field.
func (pu *PetUpdate) SetAge(f float64) *PetUpdate {
	pu.mutation.ResetAge()
	pu.mutation.SetAge(f)
	return pu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (pu *PetUpdate) SetNillableAge(f *float64) *PetUpdate {
	if f != nil {
		pu.SetAge(*f)
	}
	return pu
}

// AddAge adds f to the "age" field.
func (pu *PetUpdate) AddAge(f float64) *PetUpdate {
	pu.mutation.AddAge(f)
	return pu
}

// SetName sets the "name" field.
func (pu *PetUpdate) SetName(s string) *PetUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetUUID sets the "uuid" field.
func (pu *PetUpdate) SetUUID(u uuid.UUID) *PetUpdate {
	pu.mutation.SetUUID(u)
	return pu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (pu *PetUpdate) SetNillableUUID(u *uuid.UUID) *PetUpdate {
	if u != nil {
		pu.SetUUID(*u)
	}
	return pu
}

// ClearUUID clears the value of the "uuid" field.
func (pu *PetUpdate) ClearUUID() *PetUpdate {
	pu.mutation.ClearUUID()
	return pu
}

// SetNickname sets the "nickname" field.
func (pu *PetUpdate) SetNickname(s string) *PetUpdate {
	pu.mutation.SetNickname(s)
	return pu
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (pu *PetUpdate) SetNillableNickname(s *string) *PetUpdate {
	if s != nil {
		pu.SetNickname(*s)
	}
	return pu
}

// ClearNickname clears the value of the "nickname" field.
func (pu *PetUpdate) ClearNickname() *PetUpdate {
	pu.mutation.ClearNickname()
	return pu
}

// SetTrained sets the "trained" field.
func (pu *PetUpdate) SetTrained(b bool) *PetUpdate {
	pu.mutation.SetTrained(b)
	return pu
}

// SetNillableTrained sets the "trained" field if the given value is not nil.
func (pu *PetUpdate) SetNillableTrained(b *bool) *PetUpdate {
	if b != nil {
		pu.SetTrained(*b)
	}
	return pu
}

// SetTeamID sets the "team" edge to the User entity by ID.
func (pu *PetUpdate) SetTeamID(id int) *PetUpdate {
	pu.mutation.SetTeamID(id)
	return pu
}

// SetNillableTeamID sets the "team" edge to the User entity by ID if the given value is not nil.
func (pu *PetUpdate) SetNillableTeamID(id *int) *PetUpdate {
	if id != nil {
		pu = pu.SetTeamID(*id)
	}
	return pu
}

// SetTeam sets the "team" edge to the User entity.
func (pu *PetUpdate) SetTeam(u *User) *PetUpdate {
	return pu.SetTeamID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pu *PetUpdate) SetOwnerID(id int) *PetUpdate {
	pu.mutation.SetOwnerID(id)
	return pu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(id *int) *PetUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the "owner" edge to the User entity.
func (pu *PetUpdate) SetOwner(u *User) *PetUpdate {
	return pu.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (pu *PetUpdate) Mutation() *PetMutation {
	return pu.mutation
}

// ClearTeam clears the "team" edge to the User entity.
func (pu *PetUpdate) ClearTeam() *PetUpdate {
	pu.mutation.ClearTeam()
	return pu
}

// ClearOwner clears the "owner" edge to the User entity.
func (pu *PetUpdate) ClearOwner() *PetUpdate {
	pu.mutation.ClearOwner()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PetMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PetUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PetUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Age(); ok {
		_spec.SetField(pet.FieldAge, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedAge(); ok {
		_spec.AddField(pet.FieldAge, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.UUID(); ok {
		_spec.SetField(pet.FieldUUID, field.TypeUUID, value)
	}
	if pu.mutation.UUIDCleared() {
		_spec.ClearField(pet.FieldUUID, field.TypeUUID)
	}
	if value, ok := pu.mutation.Nickname(); ok {
		_spec.SetField(pet.FieldNickname, field.TypeString, value)
	}
	if pu.mutation.NicknameCleared() {
		_spec.ClearField(pet.FieldNickname, field.TypeString)
	}
	if value, ok := pu.mutation.Trained(); ok {
		_spec.SetField(pet.FieldTrained, field.TypeBool, value)
	}
	if pu.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.TeamTable,
			Columns: []string{pet.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.TeamTable,
			Columns: []string{pet.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PetMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetAge sets the "age" field.
func (puo *PetUpdateOne) SetAge(f float64) *PetUpdateOne {
	puo.mutation.ResetAge()
	puo.mutation.SetAge(f)
	return puo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableAge(f *float64) *PetUpdateOne {
	if f != nil {
		puo.SetAge(*f)
	}
	return puo
}

// AddAge adds f to the "age" field.
func (puo *PetUpdateOne) AddAge(f float64) *PetUpdateOne {
	puo.mutation.AddAge(f)
	return puo
}

// SetName sets the "name" field.
func (puo *PetUpdateOne) SetName(s string) *PetUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetUUID sets the "uuid" field.
func (puo *PetUpdateOne) SetUUID(u uuid.UUID) *PetUpdateOne {
	puo.mutation.SetUUID(u)
	return puo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableUUID(u *uuid.UUID) *PetUpdateOne {
	if u != nil {
		puo.SetUUID(*u)
	}
	return puo
}

// ClearUUID clears the value of the "uuid" field.
func (puo *PetUpdateOne) ClearUUID() *PetUpdateOne {
	puo.mutation.ClearUUID()
	return puo
}

// SetNickname sets the "nickname" field.
func (puo *PetUpdateOne) SetNickname(s string) *PetUpdateOne {
	puo.mutation.SetNickname(s)
	return puo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableNickname(s *string) *PetUpdateOne {
	if s != nil {
		puo.SetNickname(*s)
	}
	return puo
}

// ClearNickname clears the value of the "nickname" field.
func (puo *PetUpdateOne) ClearNickname() *PetUpdateOne {
	puo.mutation.ClearNickname()
	return puo
}

// SetTrained sets the "trained" field.
func (puo *PetUpdateOne) SetTrained(b bool) *PetUpdateOne {
	puo.mutation.SetTrained(b)
	return puo
}

// SetNillableTrained sets the "trained" field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableTrained(b *bool) *PetUpdateOne {
	if b != nil {
		puo.SetTrained(*b)
	}
	return puo
}

// SetTeamID sets the "team" edge to the User entity by ID.
func (puo *PetUpdateOne) SetTeamID(id int) *PetUpdateOne {
	puo.mutation.SetTeamID(id)
	return puo
}

// SetNillableTeamID sets the "team" edge to the User entity by ID if the given value is not nil.
func (puo *PetUpdateOne) SetNillableTeamID(id *int) *PetUpdateOne {
	if id != nil {
		puo = puo.SetTeamID(*id)
	}
	return puo
}

// SetTeam sets the "team" edge to the User entity.
func (puo *PetUpdateOne) SetTeam(u *User) *PetUpdateOne {
	return puo.SetTeamID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (puo *PetUpdateOne) SetOwnerID(id int) *PetUpdateOne {
	puo.mutation.SetOwnerID(id)
	return puo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(id *int) *PetUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the "owner" edge to the User entity.
func (puo *PetUpdateOne) SetOwner(u *User) *PetUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (puo *PetUpdateOne) Mutation() *PetMutation {
	return puo.mutation
}

// ClearTeam clears the "team" edge to the User entity.
func (puo *PetUpdateOne) ClearTeam() *PetUpdateOne {
	puo.mutation.ClearTeam()
	return puo
}

// ClearOwner clears the "owner" edge to the User entity.
func (puo *PetUpdateOne) ClearOwner() *PetUpdateOne {
	puo.mutation.ClearOwner()
	return puo
}

// Where appends a list predicates to the PetUpdate builder.
func (puo *PetUpdateOne) Where(ps ...predicate.Pet) *PetUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PetUpdateOne) Select(field string, fields ...string) *PetUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pet entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	return withHooks[*Pet, PetMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PetUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PetUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PetUpdateOne) sqlSave(ctx context.Context) (_node *Pet, err error) {
	_spec := sqlgraph.NewUpdateSpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pet.FieldID)
		for _, f := range fields {
			if !pet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Age(); ok {
		_spec.SetField(pet.FieldAge, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedAge(); ok {
		_spec.AddField(pet.FieldAge, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.UUID(); ok {
		_spec.SetField(pet.FieldUUID, field.TypeUUID, value)
	}
	if puo.mutation.UUIDCleared() {
		_spec.ClearField(pet.FieldUUID, field.TypeUUID)
	}
	if value, ok := puo.mutation.Nickname(); ok {
		_spec.SetField(pet.FieldNickname, field.TypeString, value)
	}
	if puo.mutation.NicknameCleared() {
		_spec.ClearField(pet.FieldNickname, field.TypeString)
	}
	if value, ok := puo.mutation.Trained(); ok {
		_spec.SetField(pet.FieldTrained, field.TypeBool, value)
	}
	if puo.mutation.TeamCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.TeamTable,
			Columns: []string{pet.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   pet.TeamTable,
			Columns: []string{pet.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Pet{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
