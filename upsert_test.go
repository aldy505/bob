package bob_test

import (
	"reflect"
	"testing"

	"github.com/aldy505/bob"
)

func TestUpsert(t *testing.T) {
	t.Run("should be able to generate upsert query for mysql", func(t *testing.T) {
		sql, args, err := bob.
			Upsert("users", bob.Mysql).
			Columns("name", "email").
			Values("John Doe", "john@doe.com").
			Replace("name", "John Does").
			ToSql()
		if err != nil {
			t.Error(err)
		}

		desiredSql := "INSERT INTO \"users\" (\"name\", \"email\") VALUES (?, ?) ON DUPLICATE KEY UPDATE \"name\" = ?;"
		desiredArgs := []interface{}{"John Doe", "john@doe.com", "John Does"}

		if sql != desiredSql {
			t.Error("sql is not the same as result: ", sql)
		}
		if !reflect.DeepEqual(args, desiredArgs) {
			t.Error("args is not the same as result: ", args)
		}
	})

	t.Run("should be able to generate upsert query for postgres", func(t *testing.T) {
		sql, args, err := bob.
			Upsert("users", bob.Postgresql).
			Columns("name", "email").
			Values("John Doe", "john@doe.com").
			Key("email").
			Replace("name", "John Does").
			PlaceholderFormat(bob.Dollar).
			ToSql()
		if err != nil {
			t.Error(err)
		}

		desiredSql := "INSERT INTO \"users\" (\"name\", \"email\") VALUES ($1, $2) ON CONFLICT (\"email\") DO UPDATE SET \"name\" = $3;"
		desiredArgs := []interface{}{"John Doe", "john@doe.com", "John Does"}

		if sql != desiredSql {
			t.Error("sql is not the same as result: ", sql)
		}
		if !reflect.DeepEqual(args, desiredArgs) {
			t.Error("args is not the same as result: ", args)
		}
	})

	t.Run("should be able to generate upsert query for sqlite", func(t *testing.T) {
		sql, args, err := bob.
			Upsert("users", bob.Sqlite).
			Columns("name", "email").
			Values("John Doe", "john@doe.com").
			Key("email").
			Replace("name", "John Does").
			PlaceholderFormat(bob.Question).
			ToSql()
		if err != nil {
			t.Error(err)
		}

		desiredSql := "INSERT INTO \"users\" (\"name\", \"email\") VALUES (?, ?) ON CONFLICT (\"email\") DO UPDATE SET \"name\" = ?;"
		desiredArgs := []interface{}{"John Doe", "john@doe.com", "John Does"}

		if sql != desiredSql {
			t.Error("sql is not the same as result: ", sql)
		}
		if !reflect.DeepEqual(args, desiredArgs) {
			t.Error("args is not the same as result: ", args)
		}
	})

	t.Run("should be able to generate upsert query for mssql", func(t *testing.T) {
		sql, args, err := bob.
			Upsert("users", bob.MSSql).
			Columns("name", "email").
			Values("John Doe", "john@doe.com").
			Key("email", "john@doe.com").
			Replace("name", "John Does").
			PlaceholderFormat(bob.AtP).
			ToSql()
		if err != nil {
			t.Error(err)
		}

		desiredSql := "IF NOT EXISTS (SELECT * FROM \"users\" WHERE \"email\" = @p1) INSERT INTO \"users\" (\"name\", \"email\") VALUES (@p2, @p3) ELSE UPDATE \"users\" SET \"name\" = @p4 WHERE \"email\" = @p5;"
		desiredArgs := []interface{}{"john@doe.com", "John Doe", "john@doe.com", "John Does", "john@doe.com"}

		if sql != desiredSql {
			t.Error("sql is not the same as result: ", sql)
		}
		if !reflect.DeepEqual(args, desiredArgs) {
			t.Error("args is not the same as result: ", args)
		}
	})
}