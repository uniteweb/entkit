package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/uniteweb/entkit"
	"github.com/uniteweb/entkit/optimisticlock"
	"github.com/uniteweb/entkit/softdelete"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {

	return []ent.Mixin{softdelete.SoftDeleteMixin{}, entkit.TimeMixin{}}
}

// Fields of the User.
func (User) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").Comment("用户名"),
		field.Int("age").Positive().Comment("年龄"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

type UserVersion struct {
	User
}

func (UserVersion) Mixin() []ent.Mixin {

	return []ent.Mixin{
		optimisticlock.OptimisticLockMixin{},
	}

}
