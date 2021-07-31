package bob

import (
	"errors"

	"github.com/lann/builder"
)

type RenameBuilder builder.Builder

type renameData struct {
	From string
	To string
}

func init() {
	builder.Register(RenameBuilder{}, renameData{})
}

// from sets existing table name
func (b RenameBuilder) from(name string) RenameBuilder {
	return builder.Set(b, "From", name).(RenameBuilder)
}

// to sets desired table name
func (b RenameBuilder) to(name string) RenameBuilder {
	return builder.Set(b, "To", name).(RenameBuilder)
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (b RenameBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(b).(renameData)
	return data.ToSql()
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (d *renameData) ToSql() (sqlStr string, args []interface{}, err error) {
	if len(d.From) == 0 || d.From == "" || len(d.To) == 0 || d.To == "" {
		err = errors.New("rename statement must specify a table")
	}
	sqlStr = "RENAME TABLE \""+d.From+"\" TO \""+d.To+"\";"
	return
}