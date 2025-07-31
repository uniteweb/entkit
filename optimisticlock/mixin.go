package optimisticlock

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type OptimisticLockMixin struct {
	mixin.Schema
}

func (OptimisticLockMixin) Fields() []ent.Field {

	return []ent.Field{
		field.Int("version").Positive().Default(1).Comment("乐观锁版本号"),
	}
}
