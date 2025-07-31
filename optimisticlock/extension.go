package optimisticlock

import (
	"text/template"
	"time"

	_ "embed"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

var (
	//go:embed templates/client.tmpl
	ClientTemplate string
)

type OptimisticLockExtension struct {
	entc.DefaultExtension
	Retry         bool
	RetryDuration time.Duration
}

func WithRetry() Option {

	return func(o *OptimisticLockExtension) {
		o.Retry = true
	}
}

func WithRetryDuration(d time.Duration) Option {

	return func(o *OptimisticLockExtension) {
		o.RetryDuration = d
	}
}

type Option = func(*OptimisticLockExtension)

func (e *OptimisticLockExtension) Templates() []*gen.Template {
	funcMap := template.FuncMap{

		"hasVersionField": func(g *gen.Type) bool {
			for _, f := range g.Fields {
				if f.Name == "version" {
					return true
				}

			}
			return false
		},

		"genRetry": func() bool {
			return e.Retry

		},

		"retryDuration": func() string {
			return ""
		},

		"idType": func(g *gen.Type) string {

			return g.ID.Type.Type.String()

		},
	}
	return []*gen.Template{

		gen.MustParse(gen.NewTemplate("client/additional/optimisticlock").Funcs(funcMap).Parse(ClientTemplate)),
	}
}

func NewExtension(opts ...Option) *OptimisticLockExtension {
	ex := &OptimisticLockExtension{RetryDuration: 15 * time.Millisecond}

	for _, opt := range opts {
		opt(ex)
	}
	return ex
}
