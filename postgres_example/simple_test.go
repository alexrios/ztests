package postgres_example

import (
	"context"
	"github.com/alexrios/ztests/postgres"
	"testing"
)

func Test(t *testing.T) {
	ctx := context.Background()
	testEnv, err := postgres.Setup()
	defer testEnv.Teardown()
	conn, err := testEnv.DB.Conn(ctx)
	if err != nil {
		t.Fatal(err)
	}
	var twelve int
	row := conn.QueryRowContext(ctx, "select $1", 12)
	err = row.Scan(&twelve)
	if err != nil {
		t.Fatal(err)
	}
	if twelve != 12 {
		t.FailNow()
	}
}
