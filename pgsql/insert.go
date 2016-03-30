package pgsql

import (
	"github.com/go-ozzo/ozzo-dbx"
	"fmt"
	"sort"
	"strings"
	"database/sql"
)

// InsertIntoQuery represents a Postgres-specific INSERT query.
type InsertQuery struct {
	db *DB

	table     string
	cols      dbx.Params
	conflict  []string
	returning []string
}


// NewInsertQuery creates a new InsertQuery instance.
func NewInsertQuery(db *DB) *InsertQuery {
	return &InsertQuery{
		db: db,
	}
}

func (q *InsertQuery) Insert(table string, cols dbx.Params) *InsertQuery {
	q.table = table
	q.cols = cols
	return q
}

func (q *InsertQuery) Returning(cols ...string) *InsertQuery {
	q.returning = cols
	return q
}

// Build builds the query and returns an executable Query object.
func (q *InsertQuery) Build() *dbx.Query {
	names := []string{}
	for name := range q.cols {
		names = append(names, name)
	}
	sort.Strings(names)

	params := dbx.Params{}
	columns := []string{}
	values := []string{}
	for _, name := range names {
		columns = append(columns, q.db.QuoteColumnName(name))
		value := q.cols[name]
		if e, ok := value.(dbx.Expression); ok {
			values = append(values, e.Build(q.db.DB, params))
		} else {
			values = append(values, fmt.Sprintf("{:p%v}", len(params)))
			params[fmt.Sprintf("p%v", len(params))] = value
		}
	}

	var sql string
	if len(names) == 0 {
		sql = fmt.Sprintf("INSERT INTO %v DEFAULT VALUES", q.db.QuoteTableName(q.table))
	} else {
		sql = fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)",
			q.db.QuoteTableName(q.table),
			strings.Join(columns, ", "),
			strings.Join(values, ", "),
		)
	}

	if len(q.returning) > 0 {
		names = []string{}
		for _, name := range q.returning {
			names = append(names, q.db.QuoteColumnName(name))
		}
		sql += fmt.Sprintf(" RETURNING %v", strings.Join(names, ", "))
	}

	return q.db.NewQuery(sql).Bind(params)
}

func (q *InsertQuery) One(a interface{}) error {
	return q.Build().One(a)
}

func (q *InsertQuery) All(slice interface{}) error {
	return q.Build().All(slice)
}

func (q *InsertQuery) Rows() (*dbx.Rows, error) {
	return q.Build().Rows()
}

func (s *InsertQuery) Row(a ...interface{}) error {
	return s.Build().Row(a...)
}

func (s *InsertQuery) Execute(a ...interface{}) (sql.Result, error) {
	return s.Build().Execute()
}





