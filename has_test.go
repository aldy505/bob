package bob_test

import (
	"reflect"
	"testing"

	"github.com/aldy505/bob"
)

func TestHasTable(t *testing.T) {
	sql, args, err := bob.HasTable("users").ToSql()
	if err != nil {
		t.Fatal(err.Error())
	}

	result := "SELECT * FROM information_schema.tables WHERE table_name = ? AND table_schema = current_schema();"
	if sql != result {
		t.Fatal("sql is not equal with result:", sql)
	}
	argsResult := []interface{}{"users"}
	if !reflect.DeepEqual(args, argsResult) {
		t.Fatal("args is not equal with argsResult:", args)
	}
}

func TestHasColumn(t *testing.T) {
	t.Run("should be able to create a hasColumn query", func(t *testing.T) {
		sql, args, err := bob.HasTable("users").HasColumn("name").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		result := "SELECT * FROM information_schema.columns WHERE table_name = ? AND column_name = ? AND table_schema = current_schema();"
		if sql != result {
			t.Fatal("sql is not equal with result:", sql)
		}

		argsResult := []interface{}{"users", "name"}
		if !reflect.DeepEqual(args, argsResult) {
			t.Fatal("args is not equal with argsResult:", args)
		}
	})

	t.Run("should be able to create a hasColumn query (but reversed)", func(t *testing.T) {
		sql, args, err := bob.HasColumn("name").HasTable("users").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		result := "SELECT * FROM information_schema.columns WHERE table_name = ? AND column_name = ? AND table_schema = current_schema();"
		if sql != result {
			t.Fatal("sql is not equal with result:", sql)
		}

		argsResult := []interface{}{"users", "name"}
		if !reflect.DeepEqual(args, argsResult) {
			t.Fatal("args is not equal with argsResult:", args)
		}
	})
}

func TestHas_Schema(t *testing.T) {
	sql, args, err := bob.HasTable("users").WithSchema("private").ToSql()
	if err != nil {
		t.Fatal(err.Error())
	}

	result := "SELECT * FROM information_schema.tables WHERE table_name = ? AND table_schema = ?;"
	if sql != result {
		t.Fatal("sql is not equal with result:", sql)
	}

	argsResult := []interface{}{"users", "private"}
	if !reflect.DeepEqual(args, argsResult) {
		t.Fatal("args is not equal with argsResult:", args)
	}
}

func TestHas_PlaceholderFormats(t *testing.T) {
	sql, args, err := bob.HasTable("users").HasColumn("name").PlaceholderFormat(bob.Dollar).ToSql()
	if err != nil {
		t.Fatal(err.Error())
	}

	result := "SELECT * FROM information_schema.columns WHERE table_name = $1 AND column_name = $2 AND table_schema = current_schema();"
	if sql != result {
		t.Fatal("sql is not equal with result:", sql)
	}

	argsResult := []interface{}{"users", "name"}
	if !reflect.DeepEqual(args, argsResult) {
		t.Fatal("args is not equal with argsResult:", args)
	}
}

func TestHas_EmitError(t *testing.T) {
	_, _, err := bob.HasTable("").ToSql()
	if err.Error() != "has statement should have a table name" {
		t.Fatal("error is different:", err.Error())
	}
}
