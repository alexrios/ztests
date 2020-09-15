package postgres_example

import (
	"context"
	"github.com/alexrios/ztests/postgres"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	ctx := context.Background()
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
}
