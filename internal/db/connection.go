package db

import (
	"database/sql"

	"github.com/csvitor-dev/social-media/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Env.CONNECTION_STRING)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
