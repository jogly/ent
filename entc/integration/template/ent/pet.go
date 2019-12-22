// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/template/ent/pet"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// LicensedAt holds the value of the "licensed_at" field.
	LicensedAt *time.Time `json:"licensed_at,omitempty"`
}

// FromRows scans the sql response data into Pet.
func (pe *Pet) FromRows(rows *sql.Rows) error {
	var scanpe struct {
		ID         int
		Age        sql.NullInt64
		LicensedAt sql.NullTime
	}
	// the order here should be the same as in the `pet.Columns`.
	if err := rows.Scan(
		&scanpe.ID,
		&scanpe.Age,
		&scanpe.LicensedAt,
	); err != nil {
		return err
	}
	pe.ID = scanpe.ID
	pe.Age = int(scanpe.Age.Int64)
	if scanpe.LicensedAt.Valid {
		pe.LicensedAt = new(time.Time)
		*pe.LicensedAt = scanpe.LicensedAt.Time
	}
	return nil
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&sql.NullInt64{},
		&sql.NullTime{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(values ...interface{}) error {
	if m, n := len(values), len(pet.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pe.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[0])
	} else if value.Valid {
		pe.Age = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field licensed_at", values[1])
	} else if value.Valid {
		pe.LicensedAt = new(time.Time)
		*pe.LicensedAt = value.Time
	}
	return nil
}

// QueryOwner queries the owner edge of the Pet.
func (pe *Pet) QueryOwner() *UserQuery {
	return (&PetClient{pe.config}).QueryOwner(pe)
}

// Update returns a builder for updating this Pet.
// Note that, you need to call Pet.Unwrap() before calling this method, if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pe.Age))
	if v := pe.LicensedAt; v != nil {
		builder.WriteString(", licensed_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

// FromRows scans the sql response data into Pets.
func (pe *Pets) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanpe := &Pet{}
		if err := scanpe.FromRows(rows); err != nil {
			return err
		}
		*pe = append(*pe, scanpe)
	}
	return nil
}

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
