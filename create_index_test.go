package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestCreateIndex(t *testing.T) {
	sql, _, err := bob.
		CreateIndexIfNotExists("email_idx").
		On("users").
		Unique().
		Spatial().
		Fulltext().
		Columns(bob.IndexColumn{Name: "email"}).
		ToSql()
	if err != nil {
		t.Fatal(err.Error())
	}

	result := "CREATE UNIQUE FULLTEXT SPATIAL INDEX IF NOT EXISTS email_idx ON users (email);"
	if sql != result {
		t.Fatal("sql is not equal to result:", sql)
	}
}

func TestCreateIndex_Simple(t *testing.T) {
	sql, _, err := bob.
		CreateIndex("idx_email").
		On("users").
		Columns(bob.IndexColumn{Name: "email", Collate: "DEFAULT", Extras: []string{"ASC"}}).
		Columns(bob.IndexColumn{Name: "name", Extras: []string{"DESC"}}).
		ToSql()
	if err != nil {
		t.Fatal(err.Error())
	}

	result := "CREATE INDEX idx_email ON users (email COLLATE DEFAULT ASC, name DESC);"
	if sql != result {
		t.Fatal("sql is not equal to result:", sql)
	}
}

func TestCreateIndex_Error(t *testing.T) {
	t.Run("without index", func(t *testing.T) {
		_, _, err := bob.
			CreateIndex("").
			On("users").
			Columns(bob.IndexColumn{Name: "email", Collate: "DEFAULT", Extras: []string{"ASC"}}).
			Columns(bob.IndexColumn{Name: "name", Extras: []string{"DESC"}}).
			ToSql()
		if err == nil {
			t.Fatal("error is nil")
		}

		if err.Error() != "index name is required on create index statement" {
			t.Fatal("error is not equal to result:", err.Error())
		}
	})

	t.Run("without table name", func(t *testing.T) {
		_, _, err := bob.
			CreateIndex("name").
			On("").
			Columns(bob.IndexColumn{Name: "email", Collate: "DEFAULT", Extras: []string{"ASC"}}).
			Columns(bob.IndexColumn{Name: "name", Extras: []string{"DESC"}}).
			ToSql()
		if err == nil {
			t.Fatal("error is nil")
		}

		if err.Error() != "a table name must be specified on create index statement" {
			t.Fatal("error is not equal to result:", err.Error())
		}
	})

	t.Run("without columns", func(t *testing.T) {
		_, _, err := bob.
			CreateIndex("name").
			On("users").
			ToSql()
		if err == nil {
			t.Fatal("error is nil")
		}

		if err.Error() != "should at least specify one column for create index statement" {
			t.Fatal("error is not equal to result:", err.Error())
		}
	})
}
