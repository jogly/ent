// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/joegilley/ent/entc/integration/hooks/ent/card"
	"github.com/joegilley/ent/entc/integration/hooks/ent/schema"
	"github.com/joegilley/ent/entc/integration/hooks/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cardMixin := schema.Card{}.Mixin()
	cardMixinHooks0 := cardMixin[0].Hooks()
	cardHooks := schema.Card{}.Hooks()
	card.Hooks[0] = cardMixinHooks0[0]
	card.Hooks[1] = cardHooks[0]
	card.Hooks[2] = cardHooks[1]
	cardFields := schema.Card{}.Fields()
	_ = cardFields
	// cardDescNumber is the schema descriptor for number field.
	cardDescNumber := cardFields[0].Descriptor()
	// card.DefaultNumber holds the default value on creation for the number field.
	card.DefaultNumber = cardDescNumber.Default.(string)
	// card.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	card.NumberValidator = cardDescNumber.Validators[0].(func(string) error)
	// cardDescCreatedAt is the schema descriptor for created_at field.
	cardDescCreatedAt := cardFields[2].Descriptor()
	// card.DefaultCreatedAt holds the default value on creation for the created_at field.
	card.DefaultCreatedAt = cardDescCreatedAt.Default.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescVersion is the schema descriptor for version field.
	userDescVersion := userMixinFields0[0].Descriptor()
	// user.DefaultVersion holds the default value on creation for the version field.
	user.DefaultVersion = userDescVersion.Default.(int)
}

const (
	Version = "(devel)" // Version of ent codegen.
)
