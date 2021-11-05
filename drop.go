package bob

import (
	"errors"
	"strings"

	"github.com/lann/builder"
)

type DropBuilder builder.Builder

type dropData struct {
	TableName string
	IfExists  bool
	Cascade   bool
	Restrict  bool
}

func init() {
	builder.Register(DropBuilder{}, dropData{})
}

// DropTable sets which table to be dropped
func (b DropBuilder) dropTable(name string) DropBuilder {
	return builder.Set(b, "TableName", name).(DropBuilder)
}

func (b DropBuilder) ifExists() DropBuilder {
	return builder.Set(b, "IfExists", true).(DropBuilder)
}

func (b DropBuilder) Cascade() DropBuilder {
	return builder.Set(b, "Cascade", true).(DropBuilder)
}

func (b DropBuilder) Restrict() DropBuilder {
	return builder.Set(b, "Restrict", true).(DropBuilder)
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (b DropBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(b).(dropData)
	return data.ToSql()
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (d *dropData) ToSql() (sqlStr string, args []interface{}, err error) {
	if len(d.TableName) == 0 || d.TableName == "" {
		err = errors.New("drop statement must specify a table")
	}
	var sql strings.Builder

	sql.WriteString("DROP TABLE ")

	if d.IfExists {
		sql.WriteString("IF EXISTS ")
	}

	sql.WriteString("\"" + d.TableName + "\"")

	if d.Cascade {
		sql.WriteString(" CASCADE")
	} else if d.Restrict {
		sql.WriteString(" RESTRICT")
	}

	sql.WriteString(";")

	sqlStr = sql.String()
	return
}
