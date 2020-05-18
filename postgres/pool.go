package postgres

import (
	postgresdb "github.com/alexrios/canned-containers/databases/postgres"
	pgx "github.com/jackc/pgx/v4/pgxpool"
	"testing"
)

func SetupPool(t *testing.T) (db *pgx.Pool, teardown func()) {
	dbCtx, err := postgresdb.DefaultPostgresContainer().CreatePoolContainerContext()
	if err != nil {
		t.Fatal(err)
	}
	return dbCtx.Pool, func() {
		if dbCtx.Pool != nil {
			dbCtx.Pool.Close()
		}
		if dbCtx.Ctx != nil {
			_ = dbCtx.Container.Terminate(dbCtx.Ctx)
		}
	}
}
