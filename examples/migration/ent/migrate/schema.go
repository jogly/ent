// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"github.com/jogly/ent/dialect/entsql"
	"github.com/jogly/ent/dialect/sql/schema"
	"github.com/jogly/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "owner_id", Type: field.TypeInt, Default: 0},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_users_cards",
				Columns:    []*schema.Column{CardsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "best_friend_id", Type: field.TypeUUID, Unique: true, Nullable: true, Default: "00000000-0000-0000-0000-000000000000"},
		{Name: "owner_id", Type: field.TypeInt, Default: 0},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pets_pets_best_friend",
				Columns:    []*schema.Column{PetsColumns[1]},
				RefColumns: []*schema.Column{PetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "pets_users_owner",
				Columns:    []*schema.Column{PetsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeFloat64},
		{Name: "name", Type: field.TypeString},
		{Name: "tags", Type: field.TypeJSON, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		PetsTable,
		UsersTable,
	}
)

func init() {
	CardsTable.ForeignKeys[0].RefTable = UsersTable
	PetsTable.ForeignKeys[0].RefTable = PetsTable
	PetsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.Annotation = &entsql.Annotation{
		Check: "age > 0",
	}
	UsersTable.Annotation.Checks = map[string]string{
		"name_not_empty": "name <> ''",
	}
}
