# ozzo-dbe

Database specific extensions for ozzo-dbx

(Its not ready for production)

Extension for https://github.com/go-ozzo/ozzo-dbx. Adds DB-specific features

## Postgres

Usage

```go
// Import package
import (
	"github.com/pahanini/go-ozzo-dbe/pgsql"
)

// Instead of standard ozzo Open function use DB-specific open
db, err = pgsql.Open(t.Dsn())

// Use Ext() method to access to DB-specific methods
var (
    name string
    price int
)

err = db.Ext().Insert("catalog", dbx.Params{
    "name" : "Pink Floyd",
    "price": 200,
}).Returning("name", "price").Row(&name, &price)

```




