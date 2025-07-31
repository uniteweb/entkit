package entkit

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now()).Comment("创建时间").Immutable(),
		field.Time("updated_at").Default(time.Now()).Comment("更新时间").UpdateDefault(time.Now()),
	}
}
