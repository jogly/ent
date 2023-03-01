// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/jogly/ent/entc/integration/edgefield/ent/car"
	"github.com/jogly/ent/entc/integration/edgefield/ent/metadata"
	"github.com/jogly/ent/entc/integration/edgefield/ent/node"
	"github.com/jogly/ent/entc/integration/edgefield/ent/rental"
	"github.com/jogly/ent/entc/integration/edgefield/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	carFields := schema.Car{}.Fields()
	_ = carFields
	// carDescID is the schema descriptor for id field.
	carDescID := carFields[0].Descriptor()
	// car.DefaultID holds the default value on creation for the id field.
	car.DefaultID = carDescID.Default.(func() uuid.UUID)
	metadataFields := schema.Metadata{}.Fields()
	_ = metadataFields
	// metadataDescAge is the schema descriptor for age field.
	metadataDescAge := metadataFields[1].Descriptor()
	// metadata.DefaultAge holds the default value on creation for the age field.
	metadata.DefaultAge = metadataDescAge.Default.(int)
	nodeFields := schema.Node{}.Fields()
	_ = nodeFields
	// nodeDescValue is the schema descriptor for value field.
	nodeDescValue := nodeFields[0].Descriptor()
	// node.DefaultValue holds the default value on creation for the value field.
	node.DefaultValue = nodeDescValue.Default.(int)
	rentalFields := schema.Rental{}.Fields()
	_ = rentalFields
	// rentalDescDate is the schema descriptor for date field.
	rentalDescDate := rentalFields[0].Descriptor()
	// rental.DefaultDate holds the default value on creation for the date field.
	rental.DefaultDate = rentalDescDate.Default.(func() time.Time)
}
