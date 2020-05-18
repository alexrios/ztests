package postgres_example

import (
	"context"
	"github.com/alexrios/ztests/postgres"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	ctx := context.Background()
	pool, teardown := postgres.SetupPool(t)
	defer teardown()

	var now time.Time
	row := pool.QueryRow(ctx, "select now()")
	err := row.Scan(&now)
	if err != nil {
		t.Fatal(err)
	}
	if time.Now().IsZero() {
		t.FailNow()
	}
}
