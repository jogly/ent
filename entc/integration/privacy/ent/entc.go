// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/jogly/ent/entc"
	"github.com/jogly/ent/entc/gen"
)

func main() {
	// A usage for custom options to configure the code generator to use
	// an extension and inject external dependencies in the generated API.
	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyType(&http.Client{}),
		),
		entc.FeatureNames(
			"privacy",
			"entql",
			"schema/snapshot",
		),
	}
	err := entc.Generate("./schema", &gen.Config{
		Header: `
			// Copyright 2019-present Facebook Inc. All rights reserved.
			// This source code is licensed under the Apache 2.0 license found
			// in the LICENSE file in the root directory of this source tree.

			// Code generated by ent, DO NOT EDIT.
		`,
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
