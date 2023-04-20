package sqltest_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/thecodedproject/sqltest"
)

func TestConnect(t *testing.T) {

	testCases := []struct{
		Name string
		SchemaPath string
	}{
		{
			Name: "empty",
			SchemaPath: "temp.sql",
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {

			db := sqltest.OpenMysql(t, test.SchemaPath)

			err := db.Ping()
			require.NoError(t, err)
		})
	}
}
