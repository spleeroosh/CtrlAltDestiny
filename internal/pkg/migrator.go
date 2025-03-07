package migrator

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/pressly/goose/v3"
	"github.com/pressly/goose/v3/lock"
)

type Migrator struct {
	fs embed.FS
	db *sql.DB
}

func NewMigrator(db *sql.DB, fs embed.FS) Migrator {
	return Migrator{
		db: db,
		fs: fs,
		//logger: logger,
	}
}

func (m Migrator) Up() error {
	// Ensure database is available before proceeding with migrations
	fmt.Println("KEK")
	err := m.ping(m.db)
	if err != nil {
		fmt.Println(err)
		return err
	}

	locker, err := lock.NewPostgresSessionLocker()
	if err != nil {
		return fmt.Errorf("new locker: %w", err)
	}

	prov, err := goose.NewProvider(goose.DialectPostgres, m.db, m.fs, goose.WithSessionLocker(locker))
	if err != nil {
		return fmt.Errorf("new provider: %w", err)
	}

	fmt.Println("starting migration...")

	// Perform migrations up
	if _, err := prov.Up(context.Background()); err != nil {
		if errors.Is(err, goose.ErrNoNextVersion) {
			//m.logger.Info().Msg("no migrations to apply")
			return nil
		}
		fmt.Println("migration failed")
		return fmt.Errorf("migration error: %w", err)
	}

	fmt.Println("migration succeeded")
	return nil
}

func (m Migrator) ping(stdDB *sql.DB) error {
	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.InitialInterval = 2 * time.Second
	expBackoff.MaxInterval = 5 * time.Second
	expBackoff.MaxElapsedTime = 5 * time.Minute
	if err := backoff.Retry(func() error {
		if err := stdDB.Ping(); err != nil {
			fmt.Println("ping:", err)
			return err
		}
		return nil
	}, expBackoff); err != nil {
		return fmt.Errorf("database ping attempts failed: %w", err)
	}
	return nil
}
