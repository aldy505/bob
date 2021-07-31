package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestTruncate(t *testing.T) {
	t.Run("should be able to create truncate query", func(t *testing.T) {
		sql, _, err := bob.Truncate("users").ToSql()
		if err != nil {
			t.Error(err)
		}

		result := "TRUNCATE \"users\";"
		if sql != result {
			t.Error("sql is not the same as result: ", sql)
		}
	})

	t.Run("should expect an error for no table name", func(t *testing.T) {
		_, _, err := bob.Truncate("").ToSql()
		if err == nil && err.Error() != "TRUNCATE statement must specify a table" {
			t.Error(err)
		}
	})
}
