package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// NewDB は、DBを生成します。
func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
