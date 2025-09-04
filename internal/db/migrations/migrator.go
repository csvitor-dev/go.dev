package migrations

import (
	"fmt"
	"strconv"

	"github.com/csvitor-dev/go.dev/internal/config/env"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func getMigratorInstance() (*migrate.Migrate, error) {
	url := fmt.Sprintf("mysql://%s", env.Env.CONNECTION_STRING)
	migrator, err := migrate.New(
		"file://internal/db/migrations",
		url,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create migrator instance: %w", err)
	}
	return migrator, nil
}

func Up() error {
	migrator, err := getMigratorInstance()

	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}
	return nil
}

func Down() error {
	migrator, err := getMigratorInstance()

	if err != nil {
		return err
	}

	if err := migrator.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to revert migrations: %w", err)
	}
	return nil
}

func Status() (string, error) {
	migrator, err := getMigratorInstance()

	if err != nil {
		return "", err
	}
	version, dirty, err := migrator.Version()

	if err == migrate.ErrNilVersion {
		return "No migrations have been applied yet.", nil
	}

	if err != nil {
		return "", err
	}
	status := fmt.Sprintf("Version: %d, Dirty: %s", version, strconv.FormatBool(dirty))
	return status, nil
}

func Force(version int) error {
	migrator, err := getMigratorInstance()

	if err != nil {
		return err
	}
	return migrator.Force(version)
}
