package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestDropColumn(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		sql, _, err := bob.DropColumn("users", "name").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		expected := "ALTER TABLE users DROP COLUMN name"
		if sql != expected {
			t.Fatalf("Expected %s, got %s", expected, sql)
		}
	})

	t.Run("Complex", func(t *testing.T) {
		sql, _, err := bob.DropColumn("users", "name").Suffix("CASCADE").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		expected := "ALTER TABLE users DROP COLUMN name CASCADE"
		if sql != expected {
			t.Fatalf("Got: %s, Expected: %s", sql, expected)
		}
	})
}

func TestDropConstraint(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		sql, _, err := bob.DropConstraint("users", "name").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		expected := "ALTER TABLE users DROP CONSTRAINT name"
		if sql != expected {
			t.Fatalf("Expected %s, got %s", expected, sql)
		}
	})

	t.Run("Complex", func(t *testing.T) {
		sql, _, err := bob.DropConstraint("users", "name").Suffix("CASCADE").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		expected := "ALTER TABLE users DROP CONSTRAINT name CASCADE"
		if sql != expected {
			t.Fatalf("Got: %s, Expected: %s", sql, expected)
		}
	})
}

func TestRenameColumn(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		sql, _, err := bob.RenameColumn("users", "name", "full_name").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		expected := "ALTER TABLE users RENAME COLUMN name TO full_name"
		if sql != expected {
			t.Fatalf("Expected %s, got %s", expected, sql)
		}
	})

	t.Run("Complex", func(t *testing.T) {
		sql, _, err := bob.RenameColumn("users", "name", "full_name").Suffix("CASCADE").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		expected := "ALTER TABLE users RENAME COLUMN name TO full_name CASCADE"
		if sql != expected {
			t.Fatalf("Got: %s, Expected: %s", sql, expected)
		}
	})
}

func TestRenameConstraint(t *testing.T) {
	t.Run("Plain", func(t *testing.T) {
		sql, _, err := bob.RenameConstraint("users", "name", "full_name").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		expected := "ALTER TABLE users RENAME CONSTRAINT name TO full_name"
		if sql != expected {
			t.Fatalf("Expected %s, got %s", expected, sql)
		}
	})

	t.Run("Complex", func(t *testing.T) {
		sql, _, err := bob.RenameConstraint("users", "name", "full_name").Suffix("CASCADE").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}

		expected := "ALTER TABLE users RENAME CONSTRAINT name TO full_name CASCADE"
		if sql != expected {
			t.Fatalf("Got: %s, Expected: %s", sql, expected)
		}
	})
}

func TestAlter_Error(t *testing.T) {
	_, _, err := bob.DropColumn("", "").ToSql()
	if err.Error() != "table name must not be empty" {
		t.Fatal("Expected error: table name must not be empty. Got:", err.Error())
	}

	_, _, err = bob.DropColumn("users", "").ToSql()
	if err.Error() != "the second argument must not be empty" {
		t.Fatal("Expected error: the second argument must not be empty. Got:", err.Error())
	}
}
