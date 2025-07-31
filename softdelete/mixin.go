package softdelete

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

const SoftDeleteField = "deleted_at"

type softDeleteKey struct{}

func SkipSoftDelete(ctx context.Context) context.Context {

	return context.WithValue(ctx, softDeleteKey{}, true)
}

type SoftDeleteMixin struct {
	mixin.Schema
}

func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time(SoftDeleteField).Optional().Nillable().Comment("软删除时间"),
	}
}

func (SoftDeleteMixin) Index() []ent.Index {
	return []ent.Index{
		index.Fields(SoftDeleteField),
	}
}

// func (SoftDeleteMixin) Interceptors() []ent.Interceptor {

// 	return []ent.Interceptor{

// 		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
// 			return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {

// 				if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {

// 					return next.Query(ctx, q)
// 				}

// 				value, err := next.Query(ctx, q)

// 				return value, err

// 			})
// 		}),
// 	}
// }

// type query[T any,  P ~func(*sql.Sel)]
