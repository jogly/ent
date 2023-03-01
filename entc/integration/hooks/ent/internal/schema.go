// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/jogly/ent/entc/integration/hooks/ent/schema","Package":"github.com/jogly/ent/entc/integration/hooks/ent","Schemas":[{"name":"Card","config":{"Table":""},"edges":[{"name":"owner","type":"User","ref_name":"cards","unique":true,"inverse":true}],"fields":[{"name":"number","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":"unknown","default_kind":24,"immutable":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"Exact name written on card"},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"in_hook","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"InHook is a mandatory field that is set by the hook."},{"name":"expired_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0}}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":false,"MixinIndex":0},{"Index":1,"MixedIn":false,"MixinIndex":0}],"interceptors":[{"Index":0,"MixedIn":false,"MixinIndex":0}]},{"name":"Pet","config":{"Table":""},"edges":[{"name":"owner","type":"User","ref_name":"pets","unique":true,"inverse":true}],"fields":[{"name":"delete_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":0}]},{"name":"User","config":{"Table":""},"edges":[{"name":"cards","type":"Card"},{"name":"pets","type":"Pet"},{"name":"friends","type":"User"},{"name":"best_friend","type":"User","unique":true}],"fields":[{"name":"version","type":{"Type":12,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":0,"default_kind":2,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"worth","type":{"Type":17,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"password","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"sensitive":true},{"name":"active","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":true,"default_kind":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":false,"MixinIndex":0}]}],"Features":["intercept","schema/snapshot"]}`
