package softdelete

import (
	_ "embed"
	"strings"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

var (

	//go:embed templates/interceptor.go.tmpl
	InterceptorTemplate string

	//go:embed templates/client.tmpl
	ClientTemplate string
)

const (
	AnnotationKey = "SoftDelete"
)

type SoftDeleteExtension struct {
	entc.DefaultExtension
}

type SoftDeleteAnnotation struct {
	IDType   string
	Name     string
	PkgAlias string
}

func NewExstension() (*SoftDeleteExtension, error) {

	return &SoftDeleteExtension{}, nil
}
func (SoftDeleteExtension) Hooks() []gen.Hook {

	return []gen.Hook{
		Hook(),
	}
}

func (SoftDeleteExtension) Templates() []*gen.Template {

	gens := []*gen.Template{}

	gens = append(gens, gen.MustParse(gen.NewTemplate("softdelete/interceptor").Parse(InterceptorTemplate)))

	gens = append(gens, gen.MustParse(gen.NewTemplate("client/additional/softdelete").Parse(ClientTemplate)))

	return gens

}
func Hook() gen.Hook {

	return func(next gen.Generator) gen.Generator {

		return gen.GenerateFunc(func(g *gen.Graph) error {

			for _, n := range g.Nodes {

				_, ok := n.FieldBy(func(f *gen.Field) bool {

					return f.Name == "deleted_at"
				})

				if !ok {
					continue
				}
				ident := n.ID.Type.String()

				schemaName := n.Name

				n.Annotations.Set(AnnotationKey, SoftDeleteAnnotation{
					IDType:   ident,
					Name:     schemaName,
					PkgAlias: strings.ToLower(schemaName),
				})
			}
			return next.Generate(g)
		})
	}
}
