package enttest

import (
	"testing"

	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/uniteweb/entkit"
	"github.com/uniteweb/entkit/examples"
	"github.com/uniteweb/entkit/examples/softdelete"
	"github.com/uniteweb/entkit/examples/user"
)

func TestSoftDelete(t *testing.T) {

	drv, err := sql.Open("sqlite3", "file:e1nt?mode=memory&cache=shared&_fk=1")

	if err != nil {
		t.Fatalf("error creating sqlite driver: %v", err)
	}

	client := NewClient(t, WithOptions(examples.Driver(drv)))

	client.Intercept(softdelete.Interceptor()) // add interceptor to client

	// client.Debug().User.Create().

	client.Debug().User.Create().SetName("Alice").SetAge(18).ExecX(t.Context())
	client.Debug().User.Query().Where(user.NameEQ("Alice")).AllX(t.Context())
	client.Debug().User.SoftDelete(t.Context(), 1)

}

func TestPagination(t *testing.T) {
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		t.Fatalf("error creating sqlite driver: %v", err)
	}
	client := NewClient(t, WithOptions(examples.Driver(drv)))

	qry := client.Debug().User.Query()

	entkit.WithPagination(qry, 1, 10).AllX(t.Context())

}

func TestOpsi(t *testing.T) {
	// drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	// if err != nil {
	// 	t.Fatalf("error creating sqlite driver: %v", err)
	// }
	// client := NewClient(t, WithOptions(ent.Driver(drv)))

}
