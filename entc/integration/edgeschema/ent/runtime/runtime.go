// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"
	"time"

	"github.com/jogly/ent/entc/integration/edgeschema/ent/attachedfile"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/friendship"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/group"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/relationship"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/role"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/roleuser"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/schema"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/tweetlike"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/tweettag"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/user"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/usergroup"
	"github.com/jogly/ent/entc/integration/edgeschema/ent/usertweet"
	"github.com/google/uuid"

	"github.com/jogly/ent"
	"github.com/jogly/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	attachedfileFields := schema.AttachedFile{}.Fields()
	_ = attachedfileFields
	// attachedfileDescAttachTime is the schema descriptor for attach_time field.
	attachedfileDescAttachTime := attachedfileFields[0].Descriptor()
	// attachedfile.DefaultAttachTime holds the default value on creation for the attach_time field.
	attachedfile.DefaultAttachTime = attachedfileDescAttachTime.Default.(func() time.Time)
	friendshipFields := schema.Friendship{}.Fields()
	_ = friendshipFields
	// friendshipDescWeight is the schema descriptor for weight field.
	friendshipDescWeight := friendshipFields[0].Descriptor()
	// friendship.DefaultWeight holds the default value on creation for the weight field.
	friendship.DefaultWeight = friendshipDescWeight.Default.(int)
	// friendshipDescCreatedAt is the schema descriptor for created_at field.
	friendshipDescCreatedAt := friendshipFields[1].Descriptor()
	// friendship.DefaultCreatedAt holds the default value on creation for the created_at field.
	friendship.DefaultCreatedAt = friendshipDescCreatedAt.Default.(func() time.Time)
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.DefaultName holds the default value on creation for the name field.
	group.DefaultName = groupDescName.Default.(string)
	relationship.Policy = privacy.NewPolicies(schema.Relationship{})
	relationship.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := relationship.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	relationshipFields := schema.Relationship{}.Fields()
	_ = relationshipFields
	// relationshipDescWeight is the schema descriptor for weight field.
	relationshipDescWeight := relationshipFields[0].Descriptor()
	// relationship.DefaultWeight holds the default value on creation for the weight field.
	relationship.DefaultWeight = relationshipDescWeight.Default.(int)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[1].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	roleuserFields := schema.RoleUser{}.Fields()
	_ = roleuserFields
	// roleuserDescCreatedAt is the schema descriptor for created_at field.
	roleuserDescCreatedAt := roleuserFields[0].Descriptor()
	// roleuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	roleuser.DefaultCreatedAt = roleuserDescCreatedAt.Default.(func() time.Time)
	tweetlike.Policy = privacy.NewPolicies(schema.TweetLike{})
	tweetlike.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := tweetlike.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	tweetlikeFields := schema.TweetLike{}.Fields()
	_ = tweetlikeFields
	// tweetlikeDescLikedAt is the schema descriptor for liked_at field.
	tweetlikeDescLikedAt := tweetlikeFields[0].Descriptor()
	// tweetlike.DefaultLikedAt holds the default value on creation for the liked_at field.
	tweetlike.DefaultLikedAt = tweetlikeDescLikedAt.Default.(func() time.Time)
	tweettagFields := schema.TweetTag{}.Fields()
	_ = tweettagFields
	// tweettagDescAddedAt is the schema descriptor for added_at field.
	tweettagDescAddedAt := tweettagFields[1].Descriptor()
	// tweettag.DefaultAddedAt holds the default value on creation for the added_at field.
	tweettag.DefaultAddedAt = tweettagDescAddedAt.Default.(func() time.Time)
	// tweettagDescID is the schema descriptor for id field.
	tweettagDescID := tweettagFields[0].Descriptor()
	// tweettag.DefaultID holds the default value on creation for the id field.
	tweettag.DefaultID = tweettagDescID.Default.(func() uuid.UUID)
	user.Policy = privacy.NewPolicies(schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	usergroupFields := schema.UserGroup{}.Fields()
	_ = usergroupFields
	// usergroupDescJoinedAt is the schema descriptor for joined_at field.
	usergroupDescJoinedAt := usergroupFields[0].Descriptor()
	// usergroup.DefaultJoinedAt holds the default value on creation for the joined_at field.
	usergroup.DefaultJoinedAt = usergroupDescJoinedAt.Default.(func() time.Time)
	usertweetFields := schema.UserTweet{}.Fields()
	_ = usertweetFields
	// usertweetDescCreatedAt is the schema descriptor for created_at field.
	usertweetDescCreatedAt := usertweetFields[0].Descriptor()
	// usertweet.DefaultCreatedAt holds the default value on creation for the created_at field.
	usertweet.DefaultCreatedAt = usertweetDescCreatedAt.Default.(func() time.Time)
}

const (
	Version = "(devel)" // Version of ent codegen.
)
