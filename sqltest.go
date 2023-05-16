package sqltest

import (
	"database/sql"
	"flag"
	"testing"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

var dsn = flag.String("sqltest_dsn", "root@/", "data source name for test database")

func OpenMysql(
	t *testing.T,
	schemaPath string,
) *sql.DB {

	dsnWithMultistatement := *dsn + "?multiStatements=true"

	pool, err := sql.Open("mysql", dsnWithMultistatement)
	require.NoError(t, err)

	dbName := "golang_test"

	t.Cleanup(func() {

		_, err := pool.Exec("drop database if exists " + dbName + ";")
		require.NoError(t, err)

		pool.Close()
	})

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	_, err = pool.Exec(
		"drop database if exists " + dbName + ";" +
		"create database " + dbName + ";" +
		"use " + dbName + ";",
	)
	require.NoError(t, err)

	s, err := os.ReadFile(schemaPath)
	require.NoError(t, err)

	_, err = pool.Exec(string(s))
	require.NoError(t, err, "error executing schema file")

	return pool
}

