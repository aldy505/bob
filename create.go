package bob

import (
	"bytes"
	"errors"
	"strings"

	"github.com/aldy505/bob/util"
	"github.com/lann/builder"
)

type CreateBuilder builder.Builder

type createData struct {
	TableName string
	Schema    string
	Columns   []string
	Types     []string
	Primary   string
	Unique    string
	NotNull   []string
}

func init() {
	builder.Register(CreateBuilder{}, createData{})
}

func (b CreateBuilder) Name(name string) CreateBuilder {
	return builder.Set(b, "TableName", name).(CreateBuilder)
}

func (b CreateBuilder) WithSchema(name string) CreateBuilder {
	return builder.Set(b, "Schema", name).(CreateBuilder)
}

func (b CreateBuilder) Columns(cols ...string) CreateBuilder {
	return builder.Set(b, "Columns", cols).(CreateBuilder)
}

func (b CreateBuilder) Types(types ...string) CreateBuilder {
	return builder.Set(b, "Types", types).(CreateBuilder)
}

func (b CreateBuilder) Primary(column string) CreateBuilder {
	return builder.Set(b, "Primary", column).(CreateBuilder)
}

func (b CreateBuilder) Unique(column string) CreateBuilder {
	return builder.Set(b, "Unique", column).(CreateBuilder)
}

func (b CreateBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(b).(createData)
	return data.ToSql()
}

func (d *createData) ToSql() (sqlStr string, args []interface{}, err error) {
	if len(d.TableName) == 0 || d.TableName == "" {
		err = errors.New("create statements must specify a table")
		return
	}

	if (len(d.Columns) != len(d.Types)) && len(d.Columns) > 0 {
		err = errors.New("columns and types should have equal length")
		return
	}

	sql := &bytes.Buffer{}

	sql.WriteString("CREATE TABLE ")

	if d.Schema != "" {
		sql.WriteString("`" + d.Schema + "`.")
	}

	sql.WriteString("`" + d.TableName + "`")
	sql.WriteString(" ")

	var columnTypes []string
	for i := 0; i < len(d.Columns); i++ {
		columnTypes = append(columnTypes, "`"+d.Columns[i]+"` "+d.Types[i])
	}

	sql.WriteString("(")
	sql.WriteString(strings.Join(columnTypes, ", "))
	sql.WriteString(");")

	if d.Primary != "" {
		if !util.IsIn(d.Columns, d.Primary) {
			err = errors.New("supplied primary column name doesn't exists on columns")
			return
		}
		sql.WriteString(" ")
		sql.WriteString("ALTER TABLE `" + d.TableName + "` ADD PRIMARY KEY (`" + d.Primary + "`);")
	}

	if d.Unique != "" {
		if !util.IsIn(d.Columns, d.Unique) {
			err = errors.New("supplied unique column name doesn't exists on columns")
			return
		}
		sql.WriteString(" ")
		sql.WriteString("ALTER TABLE `" + d.TableName + "` ADD UNIQUE (`" + d.Unique + "`);")
	}
	sqlStr = sql.String()
	return
}
