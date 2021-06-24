package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestCreate(t *testing.T) {
	t.Run("should return correct sql string with basic columns and types", func(t *testing.T) {
		sql, _, err := bob.CreateTable("users").Columns("name", "password", "date").Types("varchar(255)", "text", "date").ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE `users` (`name` varchar(255), `password` text, `date` date);"
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
			ToSql()
		if err != nil {
			t.Fatal(err.Error())
		}
		result := "CREATE TABLE `users` (`id` uuid, `name` varchar(255), `password` text, `date` date); ALTER TABLE `users` ADD PRIMARY KEY (`id`); ALTER TABLE `users` ADD UNIQUE (`email`)"
		if sql != result {
			t.Fatal("sql is not equal to result:", sql)
		}
	})
}
