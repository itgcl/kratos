// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/SeeMusic/kratos/examples/blog/internal/data/ent/article"
	"github.com/SeeMusic/kratos/examples/blog/internal/data/ent/comment"
	"github.com/SeeMusic/kratos/examples/blog/internal/data/ent/schema"
	"github.com/SeeMusic/kratos/examples/blog/internal/data/ent/tag"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	articleFields := schema.Article{}.Fields()
	_ = articleFields
	// articleDescCreatedAt is the schema descriptor for created_at field.
	articleDescCreatedAt := articleFields[3].Descriptor()
	// article.DefaultCreatedAt holds the default value on creation for the created_at field.
	article.DefaultCreatedAt = articleDescCreatedAt.Default.(func() time.Time)
	// articleDescUpdatedAt is the schema descriptor for updated_at field.
	articleDescUpdatedAt := articleFields[4].Descriptor()
	// article.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	article.DefaultUpdatedAt = articleDescUpdatedAt.Default.(func() time.Time)
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescCreatedAt is the schema descriptor for created_at field.
	commentDescCreatedAt := commentFields[3].Descriptor()
	// comment.DefaultCreatedAt holds the default value on creation for the created_at field.
	comment.DefaultCreatedAt = commentDescCreatedAt.Default.(func() time.Time)
	// commentDescUpdatedAt is the schema descriptor for updated_at field.
	commentDescUpdatedAt := commentFields[4].Descriptor()
	// comment.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	comment.DefaultUpdatedAt = commentDescUpdatedAt.Default.(func() time.Time)
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagFields[3].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagFields[4].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
}
