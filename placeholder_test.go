package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestReplacePlaceholder(t *testing.T) {
	t.Run("should be able to replace placeholder to dollar", func(t *testing.T) {
		sql := "INSERT INTO table_name (`col1`, `col2`, `col3`) VALUES (?, ?, ?), (?, ?, ?), (?, ?, ?);"
		result := bob.ReplacePlaceholder(sql, bob.Dollar)
		should := "INSERT INTO table_name (`col1`, `col2`, `col3`) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9);"

		if result != should {
			t.Fatal("result string doesn't match:", result)
		}
	})

	t.Run("should be able to replace placeholder to colon", func(t *testing.T) {
		sql := "INSERT INTO table_name (`col1`, `col2`, `col3`) VALUES (?, ?, ?), (?, ?, ?), (?, ?, ?);"
		result := bob.ReplacePlaceholder(sql, bob.Colon)
		should := "INSERT INTO table_name (`col1`, `col2`, `col3`) VALUES (:1, :2, :3), (:4, :5, :6), (:7, :8, :9);"

		if result != should {
			t.Fatal("result string doesn't match:", result)
		}
	})

	t.Run("should be able to replace placeholder to @p", func(t *testing.T) {
		sql := "INSERT INTO table_name (`col1`, `col2`, `col3`) VALUES (?, ?, ?), (?, ?, ?), (?, ?, ?);"
		result := bob.ReplacePlaceholder(sql, bob.AtP)
		should := "INSERT INTO table_name (`col1`, `col2`, `col3`) VALUES (@p, @p, @p), (@p, @p, @p), (@p, @p, @p);"

		if result != should {
			t.Fatal("result string doesn't match:", result)
		}
	})
}
