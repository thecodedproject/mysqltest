package sqltest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/thecodedproject/sqltest"
)

func TestOpenMysql_SingleTableSchema(t *testing.T) {

	schemaPath := "testdata/single_table.sql"

	t.Run("Ping DB", func(t *testing.T) {
		db := sqltest.OpenMysql(t, schemaPath)
		err := db.Ping()
		require.NoError(t, err)
	})

	t.Run("Insert some records and retrieve", func(t *testing.T) {
		db := sqltest.OpenMysql(t, schemaPath)

		_, err := db.Exec("insert into mytable values (1, 'hello')")
		require.NoError(t, err)

		_, err = db.Exec("insert into mytable values (2, 'world')")
		require.NoError(t, err)

		var count int
		err = db.QueryRow("select count(*) from mytable").Scan(&count)
		require.NoError(t, err)
		require.Equal(t, 2, count)

		var val1 string
		err = db.QueryRow("select s from mytable where id=1").Scan(&val1)
		require.NoError(t, err)
		assert.Equal(t, val1, "hello")

		var val2 string
		err = db.QueryRow("select s from mytable where id=2").Scan(&val2)
		require.NoError(t, err)
		assert.Equal(t, val2, "world")
	})

	t.Run("Query table after connection finds no records", func(t *testing.T) {
		db := sqltest.OpenMysql(t, schemaPath)

		r, err := db.Query("select * from mytable")
		require.NoError(t, err)

		require.False(t, r.Next(), "mytable should contain no rows")
	})
}

func TestOpenMysql_MultiTableSchema(t *testing.T) {

	schemaPath := "testdata/multi_table.sql"

	t.Run("Ping DB", func(t *testing.T) {
		db := sqltest.OpenMysql(t, schemaPath)
		err := db.Ping()
		require.NoError(t, err)
	})
}
