package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestDrop(t *testing.T) {
	t.Run("should be able to create regular drop query", func (t *testing.T)  {
		sql, _, err := bob.DropTable("users").ToSql()
		if err != nil {
			t.Error(err)
		}

		result := "DROP TABLE \"users\";"
		if sql != result {
			t.Error("sql is not the same as result: ", sql)
		}
	})

	t.Run("should be able to create drop if exists query", func(t *testing.T) {
		sql, _, err := bob.DropTableIfExists("users").ToSql()
		if err != nil {
			t.Error(err)
		}

		result := "DROP TABLE IF EXISTS \"users\";"
		if sql != result {
			t.Error("sql is not the same as result: ", sql)
		}
	})

	t.Run("should expect an error for no table name", func(t *testing.T) {
		_, _, err := bob.DropTableIfExists("").ToSql()
		if err == nil && err.Error() != "drop statement must specify a table" {
			t.Error(err)
		}
	})
} 