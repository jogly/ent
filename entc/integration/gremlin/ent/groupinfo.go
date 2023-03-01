// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/jogly/ent/dialect/gremlin"
)

// GroupInfo is the model entity for the GroupInfo schema.
type GroupInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupInfoQuery when eager-loading is set.
	Edges GroupInfoEdges `json:"edges"`
}

// GroupInfoEdges holds the relations/edges for other nodes in the graph.
type GroupInfoEdges struct {
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e GroupInfoEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// FromResponse scans the gremlin response data into GroupInfo.
func (gi *GroupInfo) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scangi struct {
		ID       string `json:"id,omitempty"`
		Desc     string `json:"desc,omitempty"`
		MaxUsers int    `json:"max_users,omitempty"`
	}
	if err := vmap.Decode(&scangi); err != nil {
		return err
	}
	gi.ID = scangi.ID
	gi.Desc = scangi.Desc
	gi.MaxUsers = scangi.MaxUsers
	return nil
}

// QueryGroups queries the "groups" edge of the GroupInfo entity.
func (gi *GroupInfo) QueryGroups() *GroupQuery {
	return NewGroupInfoClient(gi.config).QueryGroups(gi)
}

// Update returns a builder for updating this GroupInfo.
// Note that you need to call GroupInfo.Unwrap() before calling this method if this GroupInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (gi *GroupInfo) Update() *GroupInfoUpdateOne {
	return NewGroupInfoClient(gi.config).UpdateOne(gi)
}

// Unwrap unwraps the GroupInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gi *GroupInfo) Unwrap() *GroupInfo {
	_tx, ok := gi.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupInfo is not a transactional entity")
	}
	gi.config.driver = _tx.drv
	return gi
}

// String implements the fmt.Stringer.
func (gi *GroupInfo) String() string {
	var builder strings.Builder
	builder.WriteString("GroupInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gi.ID))
	builder.WriteString("desc=")
	builder.WriteString(gi.Desc)
	builder.WriteString(", ")
	builder.WriteString("max_users=")
	builder.WriteString(fmt.Sprintf("%v", gi.MaxUsers))
	builder.WriteByte(')')
	return builder.String()
}

// GroupInfos is a parsable slice of GroupInfo.
type GroupInfos []*GroupInfo

// FromResponse scans the gremlin response data into GroupInfos.
func (gi *GroupInfos) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scangi []struct {
		ID       string `json:"id,omitempty"`
		Desc     string `json:"desc,omitempty"`
		MaxUsers int    `json:"max_users,omitempty"`
	}
	if err := vmap.Decode(&scangi); err != nil {
		return err
	}
	for _, v := range scangi {
		node := &GroupInfo{ID: v.ID}
		node.Desc = v.Desc
		node.MaxUsers = v.MaxUsers
		*gi = append(*gi, node)
	}
	return nil
}
