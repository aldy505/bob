package bob

import (
	"bytes"
	"errors"
	"strings"

	"github.com/lann/builder"
)

type CreateBuilder builder.Builder

type createData struct {
	TableName   string
	IfNotExists bool
	Schema      string
	Columns     []ColumnDef
}

type ColumnDef struct {
	Name   string
	Type   string
	Extras []string
}

func init() {
	builder.Register(CreateBuilder{}, createData{})
}

// Name sets the table name
func (b CreateBuilder) Name(name string) CreateBuilder {
	return builder.Set(b, "TableName", name).(CreateBuilder)
}

// IfNotExists adds IF NOT EXISTS to the query
func (b CreateBuilder) IfNotExists() CreateBuilder {
	return builder.Set(b, "IfNotExists", true).(CreateBuilder)
}

// WithSchema specifies the schema to be used when using the schema-building commands.
func (b CreateBuilder) WithSchema(name string) CreateBuilder {
	return builder.Set(b, "Schema", name).(CreateBuilder)
}

// StringColumn creates a column with VARCHAR(255) data type.
// For SQLite please refer to TextColumn.
func (b CreateBuilder) StringColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "VARCHAR(255)",
		Extras: extras,
	}).(CreateBuilder)
}

// TextColumn creates a column with TEXT data type
func (b CreateBuilder) TextColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "TEXT",
		Extras: extras,
	}).(CreateBuilder)
}

// UUIDColumn only available for PostgreSQL
func (b CreateBuilder) UUIDColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "UUID",
		Extras: extras,
	}).(CreateBuilder)
}

// BooleanColumn only available for PostgreSQL
func (b CreateBuilder) BooleanColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "BOOLEAN",
		Extras: extras,
	}).(CreateBuilder)
}

// IntegerColumn only available for PostgreSQL and SQLite.
// For MySQL and MSSQL, please refer to IntColumn,
func (b CreateBuilder) IntegerColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "INTEGER",
		Extras: extras,
	}).(CreateBuilder)
}

// IntColumn only available for MySQL and MSSQL.
// For PostgreSQL and SQLite please refer to IntegerColumn.
func (b CreateBuilder) IntColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "INT",
		Extras: extras,
	}).(CreateBuilder)
}

func (b CreateBuilder) DateTimeColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "DATETIME",
		Extras: extras,
	}).(CreateBuilder)
}

func (b CreateBuilder) TimeStampColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "TIMESTAMP",
		Extras: extras,
	}).(CreateBuilder)
}

func (b CreateBuilder) TimeColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "TIME",
		Extras: extras,
	}).(CreateBuilder)
}

func (b CreateBuilder) DateColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "DATE",
		Extras: extras,
	}).(CreateBuilder)
}

// JSONColumn only available for MySQL and PostgreSQL.
// For MSSQL please use AddColumn(bob.ColumnDef{Name: "name", Type: "NVARCHAR(1000)"}).
// Not supported for SQLite.
func (b CreateBuilder) JSONColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "JSON",
		Extras: extras,
	}).(CreateBuilder)
}

// JSONBColumn only available for PostgreSQL.
// For MySQL please refer to JSONColumn.
func (b CreateBuilder) JSONBColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "JSONB",
		Extras: extras,
	}).(CreateBuilder)
}

// BlobColumn only available for MySQL and SQLite.
// For PostgreSQL and MSSQL, please use AddColumn(bob.ColumnDef{Name: "name", Type: "BYTEA"}).
func (b CreateBuilder) BlobColumn(name string, extras ...string) CreateBuilder {
	return builder.Append(b, "Columns", ColumnDef{
		Name:   name,
		Type:   "BLOB",
		Extras: extras,
	}).(CreateBuilder)
}

// AddColumn sets custom columns
func (b CreateBuilder) AddColumn(column ColumnDef) CreateBuilder {
	return builder.Append(b, "Columns", column).(CreateBuilder)
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (b CreateBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(b).(createData)
	return data.ToSql()
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (d *createData) ToSql() (sqlStr string, args []interface{}, err error) {
	if len(d.TableName) == 0 || d.TableName == "" {
		err = errors.New("create statements must specify a table")
		return
	}

	if len(d.Columns) == 0 {
		err = errors.New("a table should at least have one column")
		return
	}

	sql := &bytes.Buffer{}

	sql.WriteString("CREATE TABLE ")

	if d.IfNotExists {
		sql.WriteString("IF NOT EXISTS ")
	}

	if d.Schema != "" {
		sql.WriteString("\"" + d.Schema + "\".")
	}

	sql.WriteString("\"" + d.TableName + "\"")
	sql.WriteString(" ")

	var columnTypes []string
	for i := 0; i < len(d.Columns); i++ {
		var column []string
		column = append(column, "\""+d.Columns[i].Name+"\" "+d.Columns[i].Type)
		if len(d.Columns[i].Extras) > 0 {
			column = append(column, strings.Join(d.Columns[i].Extras, " "))
		}
		columnTypes = append(columnTypes, strings.Join(column, " "))
	}

	sql.WriteString("(")
	sql.WriteString(strings.Join(columnTypes, ", "))
	sql.WriteString(");")

	sqlStr = sql.String()
	return
}
