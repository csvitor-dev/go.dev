package db

import (
	"database/sql"

	"github.com/csvitor-dev/social-media/api/src/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB = nil
)

func Connect() (*sql.DB, error) {
	if DB != nil {
		return DB, nil
	}
	var err error
	DB, err = sql.Open("mysql", config.ConnectionString)

	if err != nil {
		return nil, err
	}

	if err = DB.Ping(); err != nil {
		DB.Close()
		return nil, err
	}
	return DB, nil
}