// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package viewer

import (
	"context"

	"github.com/jogly/ent/examples/privacytenant/ent"
)

// Role for viewer actions.
type Role int

// List of roles.
const (
	_ Role = 1 << iota
	Admin
	View
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Admin() bool         // If viewer is admin.
	Tenant() (int, bool) // Tenant identifier.
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	T    *ent.Tenant
	Role Role // Attached roles.
}

func (v UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

func (v UserViewer) Tenant() (int, bool) {
	if v.T != nil {
		return v.T.ID, true
	}
	return 0, false
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}
