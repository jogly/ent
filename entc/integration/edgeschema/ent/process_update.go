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
	"github.com/jogly/ent/entc/integration/edgeschema/ent/attachedfile"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/file"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/predicate"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/process"
	"github.com/jogly/ent/schema/field"
)

// ProcessUpdate is the builder for updating Process entities.
type ProcessUpdate struct {
	config
	hooks    []Hook
	mutation *ProcessMutation
}

// Where appends a list predicates to the ProcessUpdate builder.
func (pu *ProcessUpdate) Where(ps ...predicate.Process) *ProcessUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (pu *ProcessUpdate) AddFileIDs(ids ...int) *ProcessUpdate {
	pu.mutation.AddFileIDs(ids...)
	return pu
}

// AddFiles adds the "files" edges to the File entity.
func (pu *ProcessUpdate) AddFiles(f ...*File) *ProcessUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.AddFileIDs(ids...)
}

// AddAttachedFileIDs adds the "attached_files" edge to the AttachedFile entity by IDs.
func (pu *ProcessUpdate) AddAttachedFileIDs(ids ...int) *ProcessUpdate {
	pu.mutation.AddAttachedFileIDs(ids...)
	return pu
}

// AddAttachedFiles adds the "attached_files" edges to the AttachedFile entity.
func (pu *ProcessUpdate) AddAttachedFiles(a ...*AttachedFile) *ProcessUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddAttachedFileIDs(ids...)
}

// Mutation returns the ProcessMutation object of the builder.
func (pu *ProcessUpdate) Mutation() *ProcessMutation {
	return pu.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (pu *ProcessUpdate) ClearFiles() *ProcessUpdate {
	pu.mutation.ClearFiles()
	return pu
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (pu *ProcessUpdate) RemoveFileIDs(ids ...int) *ProcessUpdate {
	pu.mutation.RemoveFileIDs(ids...)
	return pu
}

// RemoveFiles removes "files" edges to File entities.
func (pu *ProcessUpdate) RemoveFiles(f ...*File) *ProcessUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.RemoveFileIDs(ids...)
}

// ClearAttachedFiles clears all "attached_files" edges to the AttachedFile entity.
func (pu *ProcessUpdate) ClearAttachedFiles() *ProcessUpdate {
	pu.mutation.ClearAttachedFiles()
	return pu
}

// RemoveAttachedFileIDs removes the "attached_files" edge to AttachedFile entities by IDs.
func (pu *ProcessUpdate) RemoveAttachedFileIDs(ids ...int) *ProcessUpdate {
	pu.mutation.RemoveAttachedFileIDs(ids...)
	return pu
}

// RemoveAttachedFiles removes "attached_files" edges to AttachedFile entities.
func (pu *ProcessUpdate) RemoveAttachedFiles(a ...*AttachedFile) *ProcessUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveAttachedFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProcessUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ProcessMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProcessUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProcessUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProcessUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProcessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(process.Table, process.Columns, sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   process.FilesTable,
			Columns: process.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		createE := &AttachedFileCreate{config: pu.config, mutation: newAttachedFileMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedFilesIDs(); len(nodes) > 0 && !pu.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   process.FilesTable,
			Columns: process.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AttachedFileCreate{config: pu.config, mutation: newAttachedFileMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   process.FilesTable,
			Columns: process.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AttachedFileCreate{config: pu.config, mutation: newAttachedFileMutation(pu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.AttachedFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   process.AttachedFilesTable,
			Columns: []string{process.AttachedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachedfile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedAttachedFilesIDs(); len(nodes) > 0 && !pu.mutation.AttachedFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   process.AttachedFilesTable,
			Columns: []string{process.AttachedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachedfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AttachedFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   process.AttachedFilesTable,
			Columns: []string{process.AttachedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachedfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{process.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProcessUpdateOne is the builder for updating a single Process entity.
type ProcessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProcessMutation
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (puo *ProcessUpdateOne) AddFileIDs(ids ...int) *ProcessUpdateOne {
	puo.mutation.AddFileIDs(ids...)
	return puo
}

// AddFiles adds the "files" edges to the File entity.
func (puo *ProcessUpdateOne) AddFiles(f ...*File) *ProcessUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.AddFileIDs(ids...)
}

// AddAttachedFileIDs adds the "attached_files" edge to the AttachedFile entity by IDs.
func (puo *ProcessUpdateOne) AddAttachedFileIDs(ids ...int) *ProcessUpdateOne {
	puo.mutation.AddAttachedFileIDs(ids...)
	return puo
}

// AddAttachedFiles adds the "attached_files" edges to the AttachedFile entity.
func (puo *ProcessUpdateOne) AddAttachedFiles(a ...*AttachedFile) *ProcessUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddAttachedFileIDs(ids...)
}

// Mutation returns the ProcessMutation object of the builder.
func (puo *ProcessUpdateOne) Mutation() *ProcessMutation {
	return puo.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (puo *ProcessUpdateOne) ClearFiles() *ProcessUpdateOne {
	puo.mutation.ClearFiles()
	return puo
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (puo *ProcessUpdateOne) RemoveFileIDs(ids ...int) *ProcessUpdateOne {
	puo.mutation.RemoveFileIDs(ids...)
	return puo
}

// RemoveFiles removes "files" edges to File entities.
func (puo *ProcessUpdateOne) RemoveFiles(f ...*File) *ProcessUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.RemoveFileIDs(ids...)
}

// ClearAttachedFiles clears all "attached_files" edges to the AttachedFile entity.
func (puo *ProcessUpdateOne) ClearAttachedFiles() *ProcessUpdateOne {
	puo.mutation.ClearAttachedFiles()
	return puo
}

// RemoveAttachedFileIDs removes the "attached_files" edge to AttachedFile entities by IDs.
func (puo *ProcessUpdateOne) RemoveAttachedFileIDs(ids ...int) *ProcessUpdateOne {
	puo.mutation.RemoveAttachedFileIDs(ids...)
	return puo
}

// RemoveAttachedFiles removes "attached_files" edges to AttachedFile entities.
func (puo *ProcessUpdateOne) RemoveAttachedFiles(a ...*AttachedFile) *ProcessUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveAttachedFileIDs(ids...)
}

// Where appends a list predicates to the ProcessUpdate builder.
func (puo *ProcessUpdateOne) Where(ps ...predicate.Process) *ProcessUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProcessUpdateOne) Select(field string, fields ...string) *ProcessUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Process entity.
func (puo *ProcessUpdateOne) Save(ctx context.Context) (*Process, error) {
	return withHooks[*Process, ProcessMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProcessUpdateOne) SaveX(ctx context.Context) *Process {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProcessUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProcessUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProcessUpdateOne) sqlSave(ctx context.Context) (_node *Process, err error) {
	_spec := sqlgraph.NewUpdateSpec(process.Table, process.Columns, sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Process.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, process.FieldID)
		for _, f := range fields {
			if !process.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != process.FieldID {
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
	if puo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   process.FilesTable,
			Columns: process.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		createE := &AttachedFileCreate{config: puo.config, mutation: newAttachedFileMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedFilesIDs(); len(nodes) > 0 && !puo.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   process.FilesTable,
			Columns: process.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AttachedFileCreate{config: puo.config, mutation: newAttachedFileMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   process.FilesTable,
			Columns: process.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AttachedFileCreate{config: puo.config, mutation: newAttachedFileMutation(puo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.AttachedFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   process.AttachedFilesTable,
			Columns: []string{process.AttachedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachedfile.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedAttachedFilesIDs(); len(nodes) > 0 && !puo.mutation.AttachedFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   process.AttachedFilesTable,
			Columns: []string{process.AttachedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachedfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AttachedFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   process.AttachedFilesTable,
			Columns: []string{process.AttachedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachedfile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Process{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{process.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
