package postgres_example

import (
	"context"
	"testing"
	"ztests/postgres"
)

func Test(t *testing.T) {
	ctx := context.Background()
	db, teardown := postgres.Setup(t)
	defer teardown()
	conn, err := db.Conn(ctx)
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
