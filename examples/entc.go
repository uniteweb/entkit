//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/uniteweb/entkit/optimisticlock"
	"github.com/uniteweb/entkit/softdelete"
)

func main() {

	var err error

	softdelete, err := softdelete.NewExtension()

	if err != nil {
		log.Fatalf("error creating softdelete extension: %v", err)
	}

	optimisticlock := optimisticlock.NewExtension(optimisticlock.WithRetry())

	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureIntercept, // must set interceptor feature to use soft delete extension
		},
	},
		entc.Extensions(softdelete, optimisticlock),
	)

	if err != nil {
		panic(err)
	}
}
