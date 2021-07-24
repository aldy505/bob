package bob

import (
	"errors"

	"github.com/lann/builder"
)

type TruncateBuilder builder.Builder

type truncateData struct {
	TableName string
}

func init() {
	builder.Register(TruncateBuilder{}, truncateData{})
}

// Truncate sets which table to be dropped
func (b TruncateBuilder) Truncate(name string) TruncateBuilder {
	return builder.Set(b, "TableName", name).(TruncateBuilder)
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (b TruncateBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(b).(truncateData)
	return data.ToSql()
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (d *truncateData) ToSql() (sqlStr string, args []interface{}, err error) {
	if len(d.TableName) == 0 || d.TableName == "" {
		err = errors.New("truncate statement must specify a table")
	}
	sqlStr = "TRUNCATE \""+d.TableName+"\";"
	return
}