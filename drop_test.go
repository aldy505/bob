package bob_test

import (
	"testing"

	"github.com/aldy505/bob"
)

func TestDrop_Regular(t *testing.T) {
	sql, _, err := bob.DropTable("users").ToSql()
	if err != nil {
		t.Error(err)
	}

	result := "DROP TABLE \"users\";"
	if sql != result {
		t.Error("sql is not the same as result: ", sql)
	}
}

func TestDrop_IfExists(t *testing.T) {
	sql, _, err := bob.DropTableIfExists("users").ToSql()
	if err != nil {
		t.Error(err)
	}

	result := "DROP TABLE IF EXISTS \"users\";"
	if sql != result {
		t.Error("sql is not the same as result: ", sql)
	}
}

func TestDrop_Cascade(t *testing.T) {
	sql, _, err := bob.DropTable("users").Cascade().ToSql()
	if err != nil {
		t.Error(err)
	}

	result := "DROP TABLE \"users\" CASCADE;"
	if sql != result {
		t.Error("sql is not the same as result: ", sql)
	}
}

func TestDrop_Restrict(t *testing.T) {
	sql, _, err := bob.DropTable("users").Restrict().ToSql()
	if err != nil {
		t.Error(err)
	}

	result := "DROP TABLE \"users\" RESTRICT;"
	if sql != result {
		t.Error("sql is not the same as result: ", sql)
	}
}

func TestDrop_ErrNoTable(t *testing.T) {
	_, _, err := bob.DropTableIfExists("").ToSql()
	if err == nil && err.Error() != "drop statement must specify a table" {
		t.Error(err)
	}
}
