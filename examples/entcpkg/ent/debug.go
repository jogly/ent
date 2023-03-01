// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import "github.com/jogly/ent/dialect"

func (c *UserClient) Debug() *UserClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	return &UserClient{config: cfg}
}

func (c *UserClient) DebugLog(fn func(...any)) *UserClient {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: fn, debug: true, hooks: c.hooks}
	return &UserClient{config: cfg}
}
