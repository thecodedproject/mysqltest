package sqltest

import (
	"database/sql"
	"flag"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

var dsn = flag.String("sqltest_dsn", "root@/test", "data source name for test database")

func OpenMysql(
	t *testing.T,
	schemaPath string,
) *sql.DB {

	pool, err := sql.Open("mysql", *dsn)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Close()
	})

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	s, err := os.ReadFile(schemaPath)
	require.NoError(t, err)

	

	return pool
}

