package pgsql

import (
	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

// DB specific db structure
type DB struct {
	*dbx.DB
	ext *Ext
}

// Returns DB-specific query builder
func (db *DB) Ext() *Ext {
	return db.ext
}

// Opens db specific database
func Open(dsn string) (db *DB, err error) {
	base, err := dbx.Open("postgres", dsn)
	if err != nil {
		return
	}
	db = &DB{
		base,
		nil,
	}
	db.ext = NewExt(db)
	return
}
