// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package enttest

import (
	"context"

	"github.com/jogly/ent/entc/integration/migrate/entv1"
	// required by schema hooks.
	_ "github.com/jogly/ent/entc/integration/migrate/entv1/runtime"

	"github.com/jogly/ent/dialect/sql/schema"
	"github.com/jogly/ent/entc/integration/migrate/entv1/migrate"
)

type (
	// TestingT is the interface that is shared between
	// testing.T and testing.B and used by enttest.
	TestingT interface {
		FailNow()
		Error(...any)
	}

	// Option configures client creation.
	Option func(*options)

	options struct {
		opts        []entv1.Option
		migrateOpts []schema.MigrateOption
	}
)

// WithOptions forwards options to client creation.
func WithOptions(opts ...entv1.Option) Option {
	return func(o *options) {
		o.opts = append(o.opts, opts...)
	}
}

// WithMigrateOptions forwards options to auto migration.
func WithMigrateOptions(opts ...schema.MigrateOption) Option {
	return func(o *options) {
		o.migrateOpts = append(o.migrateOpts, opts...)
	}
}

func newOptions(opts []Option) *options {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// Open calls entv1.Open and auto-run migration.
func Open(t TestingT, driverName, dataSourceName string, opts ...Option) *entv1.Client {
	o := newOptions(opts)
	c, err := entv1.Open(driverName, dataSourceName, o.opts...)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	migrateSchema(t, c, o)
	return c
}

// NewClient calls entv1.NewClient and auto-run migration.
func NewClient(t TestingT, opts ...Option) *entv1.Client {
	o := newOptions(opts)
	c := entv1.NewClient(o.opts...)
	migrateSchema(t, c, o)
	return c
}
func migrateSchema(t TestingT, c *entv1.Client, o *options) {
	tables, err := schema.CopyTables(migrate.Tables)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := migrate.Create(context.Background(), c.Schema, tables, o.migrateOpts...); err != nil {
		t.Error(err)
		t.FailNow()
	}
}
