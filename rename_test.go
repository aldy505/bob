package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestRename(t *testing.T) {
	t.Run("should be able to create rename query", func(t *testing.T) {
		sql, _, err := bob.RenameTable("users", "teachers").ToSql()
		if err != nil {
			t.Error(err)
		}

		result := "RENAME TABLE \"users\" TO \"teachers\";"
		if sql != result {
			t.Error("sql is not the same as result: ", sql)
		}
	})

	t.Run("should expect an error for no table name", func(t *testing.T) {
		_, _, err := bob.RenameTable("", "").ToSql()
		if err == nil && err.Error() != "rename statement must specify a table" {
			t.Error(err)
		}
	})
}
