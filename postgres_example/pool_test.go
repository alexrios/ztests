package postgres_example

import (
	"context"
	"github.com/alexrios/ztests/postgres"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	ctx := context.Background()
	t.Run("default setup", func(t *testing.T) {
		testEnv, err := postgres.SetupPGX()
		if err != nil {
			t.Fatal(err)
		}
		defer testEnv.Teardown()

		var now time.Time
		row := testEnv.Pool.QueryRow(ctx, "select now()")
		err = row.Scan(&now)
		if err != nil {
			t.Fatal(err)
		}
		if time.Now().IsZero() {
			t.FailNow()
		}
	})
	t.Run("setup with migration", func(t *testing.T) {
		testEnv, err := postgres.SetupPGX(postgres.Migrate("migrations"))
		if err != nil {
			t.Fatal(err)
		}
		defer testEnv.Teardown()

		var id int64
		row := testEnv.Pool.QueryRow(ctx, "insert into test (column2) values (13) returning column1")
		err = row.Scan(&id)
		if err != nil {
			t.Fatal(err)
		}
		if id == 0 {
			t.FailNow()
		}
	})

}
