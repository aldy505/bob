package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

// TODO - do more test

func TestHas(t *testing.T) {
	t.Run("should be able to create a hasTable query", func(t *testing.T) {
		sql, args, err := bob.HasTable("users").ToSQL()
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
		sql, args, err := bob.HasTable("users").HasColumn("name").ToSQL()
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

	t.Run("should be able to create a hasColumn query (but reversed)", func(t *testing.T) {
		sql, args, err := bob.HasColumn("name").HasTable("users").ToSQL()
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

	t.Run("should be able to create a hasTable query with schema", func(t *testing.T) {
		sql, args, err := bob.HasTable("users").WithSchema("private").ToSQL()
		if err != nil {
			t.Fatal(err.Error())
		}

		result := "SELECT * FROM information_schema.tables WHERE table_name = ? AND table_schema = ?;"
		if sql != result {
			t.Fatal("sql is not equal with result:", sql)
		}

		if len(args) != 2 {
			t.Fatal("args is not equal with argsResult:", args)
		}
	})

	t.Run("should be able to have a different placeholder", func(t *testing.T) {
		sql, args, err := bob.HasTable("users").HasColumn("name").PlaceholderFormat(bob.Dollar).ToSQL()
		if err != nil {
			t.Fatal(err.Error())
		}

		result := "SELECT * FROM information_schema.columns WHERE table_name = $1 AND column_name = $2 AND table_schema = current_schema();"
		if sql != result {
			t.Fatal("sql is not equal with result:", sql)
		}

		if len(args) != 2 {
			t.Fatal("args is not equal with argsResult:", args)
		}
	})

	t.Run("should expect an error for no table name", func(t *testing.T) {
		_, _, err := bob.HasTable("").ToSQL()
		if err.Error() != "has statement should have a table name" {
			t.Fatal("error is different:", err.Error())
		}
	})
}
