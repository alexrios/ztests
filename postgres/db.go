package postgres

import (
	"database/sql"
	"github.com/alexrios/canned-containers/databases/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"testing"
	"time"
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
		if dbCtx.Ctx != nil {
			_ = dbCtx.Container.Terminate(dbCtx.Ctx)
		}
	}
}

func SetupAndMigrate(t *testing.T, migrationPath string) (db *sql.DB, teardown func()) {
	dbCtx, err := postgresdb.DefaultPostgresContainer().CreateContainerContext()
	if err != nil {
		t.Fatal(err)
	}

	var mig *migrate.Migrate
	mig, err = migrate.New("file://"+migrationPath, dbCtx.ConnStr)
	if err != nil {
		t.Fatal(err)
	}
	if err = mig.Up(); err != nil {
		t.Fatal(err)
	}

	return dbCtx.Conn, func() {
		_ = mig.Down()
		if dbCtx.Conn != nil {
			_ = dbCtx.Conn.Close()
		}
		if dbCtx.Ctx != nil {
			_ = dbCtx.Container.Terminate(dbCtx.Ctx)
		}
	}
}

func SetupWithTimeout(t *testing.T, timeout time.Duration) (db *sql.DB, teardown func()) {
	container := postgresdb.DefaultPostgresContainer()
	container.WithTimeout(timeout)
	dbCtx, err := container.CreateContainerContext()
	if err != nil {
		t.Fatal(err)
	}
	return dbCtx.Conn, func() {
		if dbCtx.Conn != nil {
			_ = dbCtx.Conn.Close()
		}
		if dbCtx.Ctx != nil {
			_ = dbCtx.Container.Terminate(dbCtx.Ctx)
		}
	}
}

func SetupAndMigrateWithTimeout(t *testing.T, migrationPath string, timeout time.Duration) (db *sql.DB, teardown func()) {
	container := postgresdb.DefaultPostgresContainer()
	container.WithTimeout(timeout)
	dbCtx, err := container.CreateContainerContext()
	if err != nil {
		t.Fatal(err)
	}

	var mig *migrate.Migrate
	mig, err = migrate.New("file://"+migrationPath, dbCtx.ConnStr)
	if err != nil {
		t.Fatal(err)
	}
	if err = mig.Up(); err != nil {
		t.Fatal(err)
	}

	return dbCtx.Conn, func() {
		_ = mig.Down()
		if dbCtx.Conn != nil {
			_ = dbCtx.Conn.Close()
		}
		if dbCtx.Ctx != nil {
			_ = dbCtx.Container.Terminate(dbCtx.Ctx)
		}
	}
}
