// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package tweetlike

import (
	"time"

	"github.com/jogly/ent"
)

const (
	// Label holds the string label denoting the tweetlike type in the database.
	Label = "tweet_like"
	// FieldLikedAt holds the string denoting the liked_at field in the database.
	FieldLikedAt = "liked_at"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTweetID holds the string denoting the tweet_id field in the database.
	FieldTweetID = "tweet_id"
	// EdgeTweet holds the string denoting the tweet edge name in mutations.
	EdgeTweet = "tweet"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// TweetFieldID holds the string denoting the ID field of the Tweet.
	TweetFieldID = "id"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// Table holds the table name of the tweetlike in the database.
	Table = "tweet_likes"
	// TweetTable is the table that holds the tweet relation/edge.
	TweetTable = "tweet_likes"
	// TweetInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	TweetInverseTable = "tweets"
	// TweetColumn is the table column denoting the tweet relation/edge.
	TweetColumn = "tweet_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "tweet_likes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for tweetlike fields.
var Columns = []string{
	FieldLikedAt,
	FieldUserID,
	FieldTweetID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/jogly/ent/entc/integration/edgeschema/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultLikedAt holds the default value on creation for the "liked_at" field.
	DefaultLikedAt func() time.Time
)
