// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/joegilley/ent/entc/integration/cascadelete/ent/comment"
	"github.com/joegilley/ent/entc/integration/cascadelete/ent/post"
	"github.com/joegilley/ent/entc/integration/cascadelete/ent/predicate"
	"github.com/joegilley/ent/entc/integration/cascadelete/ent/user"

	"github.com/joegilley/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeComment = "Comment"
	TypePost    = "Post"
	TypeUser    = "User"
)

// CommentMutation represents an operation that mutates the Comment nodes in the graph.
type CommentMutation struct {
	config
	op            Op
	typ           string
	id            *int
	text          *string
	clearedFields map[string]struct{}
	post          *int
	clearedpost   bool
	done          bool
	oldValue      func(context.Context) (*Comment, error)
	predicates    []predicate.Comment
}

var _ ent.Mutation = (*CommentMutation)(nil)

// commentOption allows management of the mutation configuration using functional options.
type commentOption func(*CommentMutation)

// newCommentMutation creates new mutation for the Comment entity.
func newCommentMutation(c config, op Op, opts ...commentOption) *CommentMutation {
	m := &CommentMutation{
		config:        c,
		op:            op,
		typ:           TypeComment,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCommentID sets the ID field of the mutation.
func withCommentID(id int) commentOption {
	return func(m *CommentMutation) {
		var (
			err   error
			once  sync.Once
			value *Comment
		)
		m.oldValue = func(ctx context.Context) (*Comment, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Comment.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withComment sets the old Comment of the mutation.
func withComment(node *Comment) commentOption {
	return func(m *CommentMutation) {
		m.oldValue = func(context.Context) (*Comment, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CommentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CommentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CommentMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CommentMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Comment.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetText sets the "text" field.
func (m *CommentMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *CommentMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Comment entity.
// If the Comment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommentMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *CommentMutation) ResetText() {
	m.text = nil
}

// SetPostID sets the "post_id" field.
func (m *CommentMutation) SetPostID(i int) {
	m.post = &i
}

// PostID returns the value of the "post_id" field in the mutation.
func (m *CommentMutation) PostID() (r int, exists bool) {
	v := m.post
	if v == nil {
		return
	}
	return *v, true
}

// OldPostID returns the old "post_id" field's value of the Comment entity.
// If the Comment object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CommentMutation) OldPostID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPostID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPostID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPostID: %w", err)
	}
	return oldValue.PostID, nil
}

// ResetPostID resets all changes to the "post_id" field.
func (m *CommentMutation) ResetPostID() {
	m.post = nil
}

// ClearPost clears the "post" edge to the Post entity.
func (m *CommentMutation) ClearPost() {
	m.clearedpost = true
}

// PostCleared reports if the "post" edge to the Post entity was cleared.
func (m *CommentMutation) PostCleared() bool {
	return m.clearedpost
}

// PostIDs returns the "post" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// PostID instead. It exists only for internal usage by the builders.
func (m *CommentMutation) PostIDs() (ids []int) {
	if id := m.post; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetPost resets all changes to the "post" edge.
func (m *CommentMutation) ResetPost() {
	m.post = nil
	m.clearedpost = false
}

// Where appends a list predicates to the CommentMutation builder.
func (m *CommentMutation) Where(ps ...predicate.Comment) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CommentMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Comment).
func (m *CommentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CommentMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.text != nil {
		fields = append(fields, comment.FieldText)
	}
	if m.post != nil {
		fields = append(fields, comment.FieldPostID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CommentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case comment.FieldText:
		return m.Text()
	case comment.FieldPostID:
		return m.PostID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CommentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case comment.FieldText:
		return m.OldText(ctx)
	case comment.FieldPostID:
		return m.OldPostID(ctx)
	}
	return nil, fmt.Errorf("unknown Comment field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CommentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case comment.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case comment.FieldPostID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPostID(v)
		return nil
	}
	return fmt.Errorf("unknown Comment field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CommentMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CommentMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CommentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Comment numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CommentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CommentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CommentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Comment nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CommentMutation) ResetField(name string) error {
	switch name {
	case comment.FieldText:
		m.ResetText()
		return nil
	case comment.FieldPostID:
		m.ResetPostID()
		return nil
	}
	return fmt.Errorf("unknown Comment field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CommentMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.post != nil {
		edges = append(edges, comment.EdgePost)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CommentMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case comment.EdgePost:
		if id := m.post; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CommentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CommentMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CommentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedpost {
		edges = append(edges, comment.EdgePost)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CommentMutation) EdgeCleared(name string) bool {
	switch name {
	case comment.EdgePost:
		return m.clearedpost
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CommentMutation) ClearEdge(name string) error {
	switch name {
	case comment.EdgePost:
		m.ClearPost()
		return nil
	}
	return fmt.Errorf("unknown Comment unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CommentMutation) ResetEdge(name string) error {
	switch name {
	case comment.EdgePost:
		m.ResetPost()
		return nil
	}
	return fmt.Errorf("unknown Comment edge %s", name)
}

// PostMutation represents an operation that mutates the Post nodes in the graph.
type PostMutation struct {
	config
	op              Op
	typ             string
	id              *int
	text            *string
	clearedFields   map[string]struct{}
	author          *int
	clearedauthor   bool
	comments        map[int]struct{}
	removedcomments map[int]struct{}
	clearedcomments bool
	done            bool
	oldValue        func(context.Context) (*Post, error)
	predicates      []predicate.Post
}

var _ ent.Mutation = (*PostMutation)(nil)

// postOption allows management of the mutation configuration using functional options.
type postOption func(*PostMutation)

// newPostMutation creates new mutation for the Post entity.
func newPostMutation(c config, op Op, opts ...postOption) *PostMutation {
	m := &PostMutation{
		config:        c,
		op:            op,
		typ:           TypePost,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPostID sets the ID field of the mutation.
func withPostID(id int) postOption {
	return func(m *PostMutation) {
		var (
			err   error
			once  sync.Once
			value *Post
		)
		m.oldValue = func(ctx context.Context) (*Post, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Post.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPost sets the old Post of the mutation.
func withPost(node *Post) postOption {
	return func(m *PostMutation) {
		m.oldValue = func(context.Context) (*Post, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PostMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PostMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PostMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *PostMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Post.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetText sets the "text" field.
func (m *PostMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *PostMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *PostMutation) ResetText() {
	m.text = nil
}

// SetAuthorID sets the "author_id" field.
func (m *PostMutation) SetAuthorID(i int) {
	m.author = &i
}

// AuthorID returns the value of the "author_id" field in the mutation.
func (m *PostMutation) AuthorID() (r int, exists bool) {
	v := m.author
	if v == nil {
		return
	}
	return *v, true
}

// OldAuthorID returns the old "author_id" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldAuthorID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAuthorID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAuthorID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAuthorID: %w", err)
	}
	return oldValue.AuthorID, nil
}

// ClearAuthorID clears the value of the "author_id" field.
func (m *PostMutation) ClearAuthorID() {
	m.author = nil
	m.clearedFields[post.FieldAuthorID] = struct{}{}
}

// AuthorIDCleared returns if the "author_id" field was cleared in this mutation.
func (m *PostMutation) AuthorIDCleared() bool {
	_, ok := m.clearedFields[post.FieldAuthorID]
	return ok
}

// ResetAuthorID resets all changes to the "author_id" field.
func (m *PostMutation) ResetAuthorID() {
	m.author = nil
	delete(m.clearedFields, post.FieldAuthorID)
}

// ClearAuthor clears the "author" edge to the User entity.
func (m *PostMutation) ClearAuthor() {
	m.clearedauthor = true
}

// AuthorCleared reports if the "author" edge to the User entity was cleared.
func (m *PostMutation) AuthorCleared() bool {
	return m.AuthorIDCleared() || m.clearedauthor
}

// AuthorIDs returns the "author" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AuthorID instead. It exists only for internal usage by the builders.
func (m *PostMutation) AuthorIDs() (ids []int) {
	if id := m.author; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAuthor resets all changes to the "author" edge.
func (m *PostMutation) ResetAuthor() {
	m.author = nil
	m.clearedauthor = false
}

// AddCommentIDs adds the "comments" edge to the Comment entity by ids.
func (m *PostMutation) AddCommentIDs(ids ...int) {
	if m.comments == nil {
		m.comments = make(map[int]struct{})
	}
	for i := range ids {
		m.comments[ids[i]] = struct{}{}
	}
}

// ClearComments clears the "comments" edge to the Comment entity.
func (m *PostMutation) ClearComments() {
	m.clearedcomments = true
}

// CommentsCleared reports if the "comments" edge to the Comment entity was cleared.
func (m *PostMutation) CommentsCleared() bool {
	return m.clearedcomments
}

// RemoveCommentIDs removes the "comments" edge to the Comment entity by IDs.
func (m *PostMutation) RemoveCommentIDs(ids ...int) {
	if m.removedcomments == nil {
		m.removedcomments = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.comments, ids[i])
		m.removedcomments[ids[i]] = struct{}{}
	}
}

// RemovedComments returns the removed IDs of the "comments" edge to the Comment entity.
func (m *PostMutation) RemovedCommentsIDs() (ids []int) {
	for id := range m.removedcomments {
		ids = append(ids, id)
	}
	return
}

// CommentsIDs returns the "comments" edge IDs in the mutation.
func (m *PostMutation) CommentsIDs() (ids []int) {
	for id := range m.comments {
		ids = append(ids, id)
	}
	return
}

// ResetComments resets all changes to the "comments" edge.
func (m *PostMutation) ResetComments() {
	m.comments = nil
	m.clearedcomments = false
	m.removedcomments = nil
}

// Where appends a list predicates to the PostMutation builder.
func (m *PostMutation) Where(ps ...predicate.Post) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PostMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Post).
func (m *PostMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PostMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.text != nil {
		fields = append(fields, post.FieldText)
	}
	if m.author != nil {
		fields = append(fields, post.FieldAuthorID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PostMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case post.FieldText:
		return m.Text()
	case post.FieldAuthorID:
		return m.AuthorID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PostMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case post.FieldText:
		return m.OldText(ctx)
	case post.FieldAuthorID:
		return m.OldAuthorID(ctx)
	}
	return nil, fmt.Errorf("unknown Post field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) SetField(name string, value ent.Value) error {
	switch name {
	case post.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case post.FieldAuthorID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAuthorID(v)
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PostMutation) AddedFields() []string {
	var fields []string
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PostMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Post numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PostMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(post.FieldAuthorID) {
		fields = append(fields, post.FieldAuthorID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PostMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PostMutation) ClearField(name string) error {
	switch name {
	case post.FieldAuthorID:
		m.ClearAuthorID()
		return nil
	}
	return fmt.Errorf("unknown Post nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PostMutation) ResetField(name string) error {
	switch name {
	case post.FieldText:
		m.ResetText()
		return nil
	case post.FieldAuthorID:
		m.ResetAuthorID()
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PostMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.author != nil {
		edges = append(edges, post.EdgeAuthor)
	}
	if m.comments != nil {
		edges = append(edges, post.EdgeComments)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PostMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case post.EdgeAuthor:
		if id := m.author; id != nil {
			return []ent.Value{*id}
		}
	case post.EdgeComments:
		ids := make([]ent.Value, 0, len(m.comments))
		for id := range m.comments {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PostMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedcomments != nil {
		edges = append(edges, post.EdgeComments)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PostMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case post.EdgeComments:
		ids := make([]ent.Value, 0, len(m.removedcomments))
		for id := range m.removedcomments {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PostMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedauthor {
		edges = append(edges, post.EdgeAuthor)
	}
	if m.clearedcomments {
		edges = append(edges, post.EdgeComments)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PostMutation) EdgeCleared(name string) bool {
	switch name {
	case post.EdgeAuthor:
		return m.clearedauthor
	case post.EdgeComments:
		return m.clearedcomments
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PostMutation) ClearEdge(name string) error {
	switch name {
	case post.EdgeAuthor:
		m.ClearAuthor()
		return nil
	}
	return fmt.Errorf("unknown Post unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PostMutation) ResetEdge(name string) error {
	switch name {
	case post.EdgeAuthor:
		m.ResetAuthor()
		return nil
	case post.EdgeComments:
		m.ResetComments()
		return nil
	}
	return fmt.Errorf("unknown Post edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	posts         map[int]struct{}
	removedposts  map[int]struct{}
	clearedposts  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddPostIDs adds the "posts" edge to the Post entity by ids.
func (m *UserMutation) AddPostIDs(ids ...int) {
	if m.posts == nil {
		m.posts = make(map[int]struct{})
	}
	for i := range ids {
		m.posts[ids[i]] = struct{}{}
	}
}

// ClearPosts clears the "posts" edge to the Post entity.
func (m *UserMutation) ClearPosts() {
	m.clearedposts = true
}

// PostsCleared reports if the "posts" edge to the Post entity was cleared.
func (m *UserMutation) PostsCleared() bool {
	return m.clearedposts
}

// RemovePostIDs removes the "posts" edge to the Post entity by IDs.
func (m *UserMutation) RemovePostIDs(ids ...int) {
	if m.removedposts == nil {
		m.removedposts = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.posts, ids[i])
		m.removedposts[ids[i]] = struct{}{}
	}
}

// RemovedPosts returns the removed IDs of the "posts" edge to the Post entity.
func (m *UserMutation) RemovedPostsIDs() (ids []int) {
	for id := range m.removedposts {
		ids = append(ids, id)
	}
	return
}

// PostsIDs returns the "posts" edge IDs in the mutation.
func (m *UserMutation) PostsIDs() (ids []int) {
	for id := range m.posts {
		ids = append(ids, id)
	}
	return
}

// ResetPosts resets all changes to the "posts" edge.
func (m *UserMutation) ResetPosts() {
	m.posts = nil
	m.clearedposts = false
	m.removedposts = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.posts != nil {
		edges = append(edges, user.EdgePosts)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePosts:
		ids := make([]ent.Value, 0, len(m.posts))
		for id := range m.posts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedposts != nil {
		edges = append(edges, user.EdgePosts)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgePosts:
		ids := make([]ent.Value, 0, len(m.removedposts))
		for id := range m.removedposts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedposts {
		edges = append(edges, user.EdgePosts)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgePosts:
		return m.clearedposts
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgePosts:
		m.ResetPosts()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
