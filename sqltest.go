package sqltest

import (
	"database/sql"
	"flag"
	"fmt"
	"math/rand"
	"testing"
	"time"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

var dsn = flag.String("sqltest_dsn", "root@/", "data source name for test database")

func OpenMysql(
	t *testing.T,
	schemaPath string,
) *sql.DB {

	dsnWithOpts := *dsn + "?multiStatements=true&parseTime=true"

	pool, err := sql.Open("mysql", dsnWithOpts)
	require.NoError(t, err)

	// A unique DB name is used for every db connection to avoid
	// concurrency issues between tests (i.e. one test dropping
	// a db which is in use by another test)
	dbName := generateDBName()

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

func generateDBName() string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("golang_test_%x", rand.Uint64())
}

