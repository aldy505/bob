package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestCreate(t *testing.T) {
	t.Run("should return correct sql string with basic columns and types", func(t *testing.T) {
		sql, _, err := bob.
			CreateTable("users").
			StringColumn("name").
			TextColumn("password").
			DateColumn("date").
			ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE \"users\" (\"name\" VARCHAR(255), \"password\" TEXT, \"date\" DATE);"
		if sql != result {
			t.Fatal("sql is not equal to result:", sql)
		}
	})

	t.Run("should return correct sql with primary key and unique key", func(t *testing.T) {
		sql, _, err := bob.CreateTable("users").
			UUIDColumn("id", "PRIMARY KEY").
			StringColumn("name").
			StringColumn("email", "UNIQUE").
			TextColumn("password").
			DateTimeColumn("date").
			ToSql()

		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE \"users\" (\"id\" UUID PRIMARY KEY, \"name\" VARCHAR(255), \"email\" VARCHAR(255) UNIQUE, \"password\" TEXT, \"date\" DATETIME);"
		if sql != result {
			t.Fatal("sql is not equal to result:", sql)
		}
	})

	t.Run("should be able to have a schema name", func(t *testing.T) {
		sql, _, err := bob.
			CreateTable("users").
			WithSchema("private").
			StringColumn("name").
			TextColumn("password").
			DateColumn("date").
			ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE \"private\".\"users\" (\"name\" VARCHAR(255), \"password\" TEXT, \"date\" DATE);"
		if sql != result {
			t.Fatal("sql is not equal to result:", sql)
		}
	})

	t.Run("should emit error on empty table name", func(t *testing.T) {
		_, _, err := bob.
			CreateTable("").
			StringColumn("name").
			ToSql()
		if err.Error() != "create statements must specify a table" {
			t.Fatal("should throw an error, it didn't:", err.Error())
		}
	})

	t.Run("should emit create if not exists", func(t *testing.T) {
		sql, _, err := bob.
			CreateTableIfNotExists("users").
			TextColumn("name").
			ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE IF NOT EXISTS \"users\" (\"name\" TEXT);"
		if sql != result {
			t.Fatal("sql is not equal to result: ", sql)
		}
	})
}
