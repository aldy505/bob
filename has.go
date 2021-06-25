package bob

import (
	"bytes"
	"errors"

	"github.com/lann/builder"
)

// TODO - The whole file is a todo
// Meant to find two things: HasTable and HasColumn(s)

type HasBuilder builder.Builder

type hasData struct {
	Name        string
	Column      string
	Schema      string
	Placeholder string
}

func init() {
	builder.Register(HasBuilder{}, hasData{})
}

func (h HasBuilder) HasTable(table string) HasBuilder {
	return builder.Set(h, "Name", table).(HasBuilder)
}

func (h HasBuilder) HasColumn(column string) HasBuilder {
	return builder.Set(h, "Column", column).(HasBuilder)
}

func (h HasBuilder) WithSchema(schema string) HasBuilder {
	return builder.Set(h, "Schema", schema).(HasBuilder)
}

func (h HasBuilder) PlaceholderFormat(f string) HasBuilder {
	return builder.Set(h, "Placeholder", f).(HasBuilder)
}

func (h HasBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(h).(hasData)
	return data.ToSql()
}

func (d *hasData) ToSql() (sqlStr string, args []interface{}, err error) {
	sql := &bytes.Buffer{}
	if d.Name == "" {
		err = errors.New("has statement should have a table name")
		return
	}

	if d.Column != "" && d.Name != "" {
		// search for column
		sql.WriteString("SELECT * FROM information_schema.columns WHERE table_name = ? AND column_name = ?")
	} else if d.Name != "" && d.Column == "" {
		sql.WriteString("SELECT * FROM information_schema.tables WHERE table_name = ?")
	}

	if d.Schema != "" {
		sql.WriteString(" AND table_schema = ?;")
	} else {
		sql.WriteString(" AND table_schema = current_schema();")
	}

	sqlStr = ReplacePlaceholder(sql.String(), d.Placeholder)
	args = createArgs(d.Name, d.Column, d.Schema)
	return
}
