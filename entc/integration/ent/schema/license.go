// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/jogly/ent"
	"github.com/jogly/ent/dialect"
	"github.com/jogly/ent/schema/field"
	"github.com/jogly/ent/schema/mixin"

	"ariga.io/atlas/sql/postgres"
)

// License holds the schema definition for the License entity.
type License struct {
	ent.Schema
}

func (License) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the License.
func (License) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.Postgres: postgres.TypeBigSerial,
			}),
	}
}
