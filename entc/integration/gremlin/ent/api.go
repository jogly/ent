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

// Api is the model entity for the Api schema.
type Api struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromResponse scans the gremlin response data into Api.
func (a *Api) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scana struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scana); err != nil {
		return err
	}
	a.ID = scana.ID
	return nil
}

// Update returns a builder for updating this Api.
// Note that you need to call Api.Unwrap() before calling this method if this Api
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Api) Update() *APIUpdateOne {
	return NewAPIClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Api entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Api) Unwrap() *Api {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Api is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Api) String() string {
	var builder strings.Builder
	builder.WriteString("Api(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Apis is a parsable slice of Api.
type Apis []*Api

// FromResponse scans the gremlin response data into Apis.
func (a *Apis) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scana []struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scana); err != nil {
		return err
	}
	for _, v := range scana {
		node := &Api{ID: v.ID}
		*a = append(*a, node)
	}
	return nil
}
