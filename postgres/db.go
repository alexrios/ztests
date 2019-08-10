package postgres

import (
	"database/sql"
	"github.com/alexrios/canned-containers/databases/postgres"
	"testing"
)

func Setup(t *testing.T) (db *sql.DB, teardown func()) {
	dbCtx, err := postgresdb.DefaultPostgresContainer().CreateContainerContext()
	if err != nil {
		t.Fatal(err)
	}
	return dbCtx.Conn, func() {
		if dbCtx.Conn != nil {
			_ = dbCtx.Conn.Close()
		}
		_ = dbCtx.Container.Terminate(dbCtx.Ctx)
	}
}
