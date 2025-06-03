package database

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	DB   *sql.DB
	once sync.Once
)

func New(url string) (*sql.DB, error) {
	var err error
	once.Do(func() {
		DB, err = sql.Open("postgres", url)
	})

	return DB, err
}
