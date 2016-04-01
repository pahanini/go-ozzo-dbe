package pgsql

import (
	"github.com/go-ozzo/ozzo-dbx"
)

// PgsqlBuilder is the builder for PostgreSQL databases.
type Ext struct {
	db *DB
}

func (e *Ext) Insert(tablename string, cols dbx.Params) *InsertQuery {
	return NewInsertQuery(e.db).Insert(tablename, cols)
}

func NewExt(db *DB) *Ext {
	return &Ext{db}
}
