# mysqltest

A package for testing code which integrates with SQL DBs.

## usage

```
import (
  "testing"

  "github.com/codedproject/mysqltest"
)

func TestSome(t *testing.T) {

  db := mysqltest.OpenMysql(t, "path/to/schema.sql")

  // then use `db` as normal; e.g.
  ... := db.Exec("SOME QUERY")
}

```

where `path/to/schema.sql` contains the table schemas:
```
create table sometable (
  i int
);
```

## detail

1. inside the test a connection to a mysql db is created; this will:
  * Connect the mysql deamon
  * Create a new `golang_test` db
  * Read the schema file and exectue its contents on the new db

2. perform test any testing on your DB code

3. on `t.Cleanup` the `golang_test` db is dropped so test results are reproducable

**note:** Beware the scope of db connection affects when the db will be dropped.
E.g. in a table test - probably best to create the db connection in each test to avoid any inserted records being visble to subsequent tests:

## TODO

* expand to allow other dbs
