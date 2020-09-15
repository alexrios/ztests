package postgres

import (
	"database/sql"
	"github.com/alexrios/canned-containers/databases/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Options struct {
	migrationPath string
	timeout       string
}

func (o Options) HasMigrations() bool {
	return len(o.migrationPath) != 0
}

func (o Options) HasTimeout() bool {
	return len(o.timeout) != 0
}

type Option func(*Options)

func Migrate(path string) Option {
	return func(option *Options) {
		option.migrationPath = path
	}
}

func Timeout(duration string) Option {
	return func(option *Options) {
		option.timeout = duration
	}
}

type DBTestEnv struct {
	DB       *sql.DB
	Teardown func()
}

type DBPoolTestEnv struct {
	Pool     *pgxpool.Pool
	Teardown func()
}

func SetupPGX(options ...Option) (*DBPoolTestEnv, error) {
	defaultOptions := Options{}

	opt := &defaultOptions
	for _, op := range options {
		op(opt)
	}

	container := postgresdb.DefaultPostgresContainer()
	if opt.HasTimeout() {
		duration, err := time.ParseDuration(opt.timeout)
		if err != nil {
			return nil, err
		}
		container.WithTimeout(duration)
	}
	hasMigrations := opt.HasMigrations()

	dbCtx, err := container.CreatePoolContainerContext()
	if err != nil {
		return nil, err
	}

	var mig *migrate.Migrate
	if hasMigrations {
		mig, err = migrate.New("file://"+opt.migrationPath, dbCtx.ConnStr)
		if err != nil {
			return nil, err
		}
		if err = mig.Up(); err != nil {
			return nil, err
		}
	}
	return &DBPoolTestEnv{
		Pool: dbCtx.Pool,
		Teardown: func() {
			if hasMigrations {
				_ = mig.Down()
			}
			if dbCtx.Pool != nil {
				dbCtx.Pool.Close()
			}
			if dbCtx.Ctx != nil {
				_ = dbCtx.Container.Terminate(dbCtx.Ctx)
			}
		},
	}, nil
}

func Setup(options ...Option) (*DBTestEnv, error) {
	defaultOptions := Options{}

	opt := &defaultOptions
	for _, op := range options {
		op(opt)
	}

	container := postgresdb.DefaultPostgresContainer()
	if opt.HasTimeout() {
		duration, err := time.ParseDuration(opt.timeout)
		if err != nil {
			return nil, err
		}
		container.WithTimeout(duration)
	}
	hasMigrations := opt.HasMigrations()

	dbCtx, err := container.CreateContainerContext()
	if err != nil {
		return nil, err
	}

	var mig *migrate.Migrate
	if hasMigrations {
		mig, err = migrate.New("file://"+opt.migrationPath, dbCtx.ConnStr)
		if err != nil {
			return nil, err
		}
		if err = mig.Up(); err != nil {
			return nil, err
		}
	}
	return &DBTestEnv{
		DB: dbCtx.Conn,
		Teardown: func() {
			if hasMigrations {
				_ = mig.Down()
			}
			if dbCtx.Conn != nil {
				_ = dbCtx.Conn.Close()
			}
			if dbCtx.Ctx != nil {
				_ = dbCtx.Container.Terminate(dbCtx.Ctx)
			}
		},
	}, nil
}
