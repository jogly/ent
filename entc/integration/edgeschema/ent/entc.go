// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/jogly/ent/entc"
	"github.com/jogly/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Header: `
			// Copyright 2019-present Facebook Inc. All rights reserved.
			// This source code is licensed under the Apache 2.0 license found
			// in the LICENSE file in the root directory of this source tree.

			// Code generated by ent, DO NOT EDIT.
		`,
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeatureUpsert,
			gen.FeaturePrivacy,
			gen.FeatureSnapshot,
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
