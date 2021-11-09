package bob_test

import (
	"reflect"
	"testing"

	"github.com/aldy505/bob"
)

func TestUpsert_MySQL(t *testing.T) {
	sql, args, err := bob.
		Upsert("users", bob.MySQL).
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
}

func TestUpsert_PostgreSQL(t *testing.T) {
	sql, args, err := bob.
		Upsert("users", bob.PostgreSQL).
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
}

func TestUpsert_SQLite(t *testing.T) {
	sql, args, err := bob.
		Upsert("users", bob.SQLite).
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
}

func TestUpsert_MSSQL(t *testing.T) {
	sql, args, err := bob.
		Upsert("users", bob.MSSQL).
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
}

func TestUpsert_EmitErrors(t *testing.T) {
	t.Run("should emit error without table name", func(t *testing.T) {
		_, _, err := bob.Upsert("", bob.MySQL).ToSql()
		if err == nil && err.Error() != "upsert statement must specify a table" {
			t.Error(err)
		}
	})

	t.Run("should emit error without columns", func(t *testing.T) {
		_, _, err := bob.Upsert("users", bob.PostgreSQL).ToSql()
		if err.Error() != "upsert statement must have at least one column" {
			t.Error(err)
		}
	})

	t.Run("should emit error without values", func(t *testing.T) {
		_, _, err := bob.Upsert("users", bob.PostgreSQL).Columns("name", "email").ToSql()
		if err.Error() != "upsert statements must have at least one set of values" {
			t.Error(err)
		}
	})

	t.Run("should emit error without replaces", func(t *testing.T) {
		_, _, err := bob.Upsert("users", bob.PostgreSQL).Columns("name", "email").Values("James", "james@mail.com").ToSql()
		if err.Error() != "upsert statement must have at least one key value pair to be replaced" {
			t.Error(err)
		}
	})

	t.Run("should emit error without key and value for mssql", func(t *testing.T) {
		_, _, err := bob.Upsert("users", bob.MSSQL).Columns("name", "email").Values("James", "james@mail.com").Replace("name", "Thomas").ToSql()
		if err.Error() != "unique key and value must be provided for MS SQL" {
			t.Error(err)
		}
	})

	t.Run("should emit error without key and value for mssql", func(t *testing.T) {
		_, _, err := bob.Upsert("users", bob.SQLite).Columns("name", "email").Values("James", "james@mail.com").Replace("name", "Thomas").ToSql()
		if err.Error() != "unique key must be provided for PostgreSQL and SQLite" {
			t.Log(err.Error())
			t.Error(err)
		}
	})

	t.Run("should emit error if dialect not supported", func(t *testing.T) {
		_, _, err := bob.Upsert("users", 100).Columns("name", "email").Values("James", "james@mail.com").Replace("name", "Thomas").ToSql()
		if err.Error() != "provided database dialect is not supported" {
			t.Log(err.Error())
			t.Error(err)
		}
	})
}

func TestUpsert_WithoutReplacePlaceHolder(t *testing.T) {
	t.Run("PostgreSQL", func(t *testing.T) {
		sql, args, err := bob.
			Upsert("users", bob.PostgreSQL).
			Columns("name", "email").
			Values("John Doe", "john@doe.com").
			Key("email").
			Replace("name", "John Does").
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

	t.Run("MSSQL", func(t *testing.T) {
		sql, args, err := bob.
			Upsert("users", bob.MSSQL).
			Columns("name", "email").
			Values("John Doe", "john@doe.com").
			Key("email", "john@doe.com").
			Replace("name", "John Does").
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

	t.Run("SQLite", func(t *testing.T) {
		sql, args, err := bob.
			Upsert("users", bob.SQLite).
			Columns("name", "email").
			Values("John Doe", "john@doe.com").
			Key("email").
			Replace("name", "John Does").
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
}
