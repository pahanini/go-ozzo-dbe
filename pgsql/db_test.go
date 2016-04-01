package pgsql

import (
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPgsql(t *testing.T) {
	db := getPgsqlDb()
	_, err := db.Insert("tests", dbx.Params{
		"code":  "acdc",
		"name":  "AC/DC",
		"price": 100,
	}).Execute()
	assert.Nil(t, err)

	var (
		name  string
		price int
	)

	err = db.Ext().Insert("tests", dbx.Params{
		"code":  "flo",
		"name":  "Pink Floyd",
		"price": 200,
	}).Returning("name", "price").Row(&name, &price)
	assert.Equal(t, "Pink Floyd", name)
	assert.Equal(t, 200, price)
	assert.Nil(t, err)
}

func getPgsqlDb() *DB {
	db, err := Open(os.Getenv("DBE_TEST_DSN"))
	if err != nil {
		panic(err)
	}
	_, err = db.NewQuery(`
		DROP TABLE IF EXISTS tests;
		CREATE TABLE "tests"
		(
			id SERIAL PRIMARY KEY,
			code VARCHAR(4) UNIQUE,
			name VARCHAR(200),
			price INT
		)
	`).Execute()
	if err != nil {
		panic(err)
	}
	return db
}
