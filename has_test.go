package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

// TODO - do more test

func TestHas(t *testing.T) {
	t.Run("should be able to create a hasTable query", func(t *testing.T) {
		sql, args, err := bob.HasTable("users").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		result := "SELECT * FROM information_schema.tables WHERE table_name = ? AND table_schema = current_schema();"
		if sql != result {
			t.Fatal("sql is not equal with result:", sql)
		}

		if len(args) != 1 {
			t.Fatal("args is not equal with argsResult:", args)
		}
	})

	t.Run("should be able to create a hasColumn query", func(t *testing.T) {
		sql, args, err := bob.HasTable("users").HasColumn("name").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		result := "SELECT * FROM information_schema.columns WHERE table_name = ? AND column_name = ? AND table_schema = current_schema();"
		if sql != result {
			t.Fatal("sql is not equal with result:", sql)
		}

		if len(args) != 2 {
			t.Fatal("args is not equal with argsResult:", args)
		}
	})
}
