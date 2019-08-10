package postgres

import (
	"context"
	"testing"
)

func Test(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	_, err := db.Conn(context.Background())

	if err != nil {
		t.FailNow()
	}
}
