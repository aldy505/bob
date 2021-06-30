package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestCreate(t *testing.T) {
	t.Run("should return correct sql string with basic columns and types", func(t *testing.T) {
		sql, _, err := bob.CreateTable("users").Columns("name", "password", "date").Types("varchar(255)", "text", "date").ToSQL()
		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE \"users\" (\"name\" varchar(255), \"password\" text, \"date\" date);"
		if sql != result {
			t.Fatal("sql is not equal to result:", sql)
		}
	})

	t.Run("should return correct sql with primary key and unique key", func(t *testing.T) {
		sql, _, err := bob.CreateTable("users").
			Columns("id", "name", "email", "password", "date").
			Types("uuid", "varchar(255)", "varchar(255)", "text", "date").
			Primary("id").
			Unique("email").
			ToSQL()
		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE \"users\" (\"id\" uuid, \"name\" varchar(255), \"email\" varchar(255), \"password\" text, \"date\" date); ALTER TABLE \"users\" ADD PRIMARY KEY (\"id\"); ALTER TABLE \"users\" ADD UNIQUE (\"email\");"
		if sql != result {
			t.Fatal("sql is not equal to result:", sql)
		}
	})

	t.Run("should be able to have a schema name", func(t *testing.T) {
		sql, _, err := bob.CreateTable("users").WithSchema("private").Columns("name", "password", "date").Types("varchar(255)", "text", "date").ToSQL()
		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE \"private\".\"users\" (\"name\" varchar(255), \"password\" text, \"date\" date);"
		if sql != result {
			t.Fatal("sql is not equal to result:", sql)
		}
	})

	t.Run("should emit error on unmatched column and types length", func(t *testing.T) {
		_, _, err := bob.CreateTable("users").
			Columns("id", "name", "email", "password", "date").
			Types("uuid", "varchar(255)", "varchar(255)", "date").
			ToSQL()
		if err.Error() != "columns and types should have equal length" {
			t.Fatal("should throw an error, it didn't:", err.Error())
		}
	})

	t.Run("should emit error on empty table name", func(t *testing.T) {
		_, _, err := bob.CreateTable("").Columns("name").Types("text").ToSQL()
		if err.Error() != "create statements must specify a table" {
			t.Fatal("should throw an error, it didn't:", err.Error())
		}
	})

	t.Run("should emit error for primary key not in columns", func(t *testing.T) {
		_, _, err := bob.CreateTable("users").Columns("name").Types("text").Primary("id").ToSQL()
		if err.Error() != "supplied primary column name doesn't exists on columns" {
			t.Fatal("should throw an error, it didn't:", err.Error())
		}
	})

	t.Run("should emit error for unique key not in columns", func(t *testing.T) {
		_, _, err := bob.CreateTable("users").Columns("name").Types("text").Unique("id").ToSQL()
		if err.Error() != "supplied unique column name doesn't exists on columns" {
			t.Fatal("should throw an error, it didn't:", err.Error())
		}
	})
}
