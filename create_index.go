package bob

import (
	"errors"
	"strings"

	"github.com/lann/builder"
)

type IndexBuilder builder.Builder

type indexData struct {
	Unique bool
	Spatial bool
	Fulltext bool
	Name string
	TableName string
	Columns []IndexColumn
	IfNotExists bool
}

type IndexColumn struct {
	Name string
	Extras []string
	Collate string
}

func init() {
	builder.Register(IndexBuilder{}, indexData{})
}

func (i IndexBuilder) Unique() IndexBuilder {
	return builder.Set(i, "Unique", true).(IndexBuilder)
}

func (i IndexBuilder) Spatial() IndexBuilder {
	return builder.Set(i, "Spatial", true).(IndexBuilder)
}

func (i IndexBuilder) Fulltext() IndexBuilder {
	return builder.Set(i, "Fulltext", true).(IndexBuilder)
}

func (i IndexBuilder) name(name string) IndexBuilder {
	return builder.Set(i, "Name", name).(IndexBuilder)
}

func (i IndexBuilder) ifNotExists() IndexBuilder {
	return builder.Set(i, "IfNotExists", true).(IndexBuilder)
}

func (i IndexBuilder) On(table string) IndexBuilder {
	return builder.Set(i, "TableName", table).(IndexBuilder)
}

func (i IndexBuilder) Columns(column IndexColumn) IndexBuilder {
	return builder.Append(i, "Columns", column).(IndexBuilder)
}

func (i IndexBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(i).(indexData)
	return data.ToSql()
}

func (i *indexData) ToSql() (sqlStr string, args []interface{}, err error) {
	if i.Name == "" {
		err = errors.New("index name is required on create index statement")
		return
	}
	if i.TableName == "" {
		err = errors.New("a table name must be specified on create index statement")
		return
	}

	if len(i.Columns) == 0 {
		err = errors.New("should at least specify one column for create index statement")
		return
	}
	
	var sql strings.Builder

	sql.WriteString("CREATE ")

	if i.Unique {
		sql.WriteString("UNIQUE ")
	}

	if i.Fulltext {
		sql.WriteString("FULLTEXT ")
	}

	if i.Spatial {
		sql.WriteString("SPATIAL ")
	}

	sql.WriteString("INDEX ")

	if i.IfNotExists {
		sql.WriteString("IF NOT EXISTS ")
	}

	sql.WriteString(i.Name + " ")

	sql.WriteString("ON ")

	sql.WriteString(i.TableName + " ")

	var columns []string
	for _, column := range i.Columns {
		var colBuilder strings.Builder
		colBuilder.WriteString(column.Name)
		if column.Collate != "" {
			colBuilder.WriteString(" COLLATE " + column.Collate)
		}
		if len(column.Extras) > 0 {
			colBuilder.WriteString(" " + strings.Join(column.Extras, " "))
		}
		columns = append(columns, colBuilder.String())
	}

	sql.WriteString("(")
	sql.WriteString(strings.Join(columns, ", "))
	sql.WriteString(");")
	
	sqlStr = sql.String()
	return
}