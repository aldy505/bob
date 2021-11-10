package bob

import (
	"errors"
	"strings"

	"github.com/lann/builder"
)

type AlterBuilder builder.Builder

type alter int

const (
	alterDropColumn alter = iota
	alterDropConstraint
	alterRenameColumn
	alterRenameConstraint
)

type alterData struct {
	What      alter
	TableName string
	FirstKey  string
	SecondKey string
	Suffix    string
}

func init() {
	builder.Register(AlterBuilder{}, alterData{})
}

func (b AlterBuilder) whatToAlter(what alter) AlterBuilder {
	return builder.Set(b, "What", what).(AlterBuilder)
}

func (b AlterBuilder) tableName(table string) AlterBuilder {
	return builder.Set(b, "TableName", table).(AlterBuilder)
}

func (b AlterBuilder) firstKey(key string) AlterBuilder {
	return builder.Set(b, "FirstKey", key).(AlterBuilder)
}

func (b AlterBuilder) secondKey(key string) AlterBuilder {
	return builder.Set(b, "SecondKey", key).(AlterBuilder)
}

func (b AlterBuilder) Suffix(any string) AlterBuilder {
	return builder.Set(b, "Suffix", any).(AlterBuilder)
}

func (b AlterBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(b).(alterData)
	return data.ToSql()
}

func (d *alterData) ToSql() (sqlStr string, args []interface{}, err error) {
	if d.TableName == "" {
		err = errors.New("table name must not be empty")
		return
	}

	if d.FirstKey == "" {
		err = errors.New("the second argument must not be empty")
		return
	}

	var sql strings.Builder

	sql.WriteString("ALTER TABLE ")

	sql.WriteString(d.TableName + " ")

	switch d.What {
	case alterDropColumn:
		sql.WriteString("DROP COLUMN " + d.FirstKey)
	case alterDropConstraint:
		sql.WriteString("DROP CONSTRAINT " + d.FirstKey)
	case alterRenameColumn:
		sql.WriteString("RENAME COLUMN " + d.FirstKey + " TO " + d.SecondKey)
	case alterRenameConstraint:
		sql.WriteString("RENAME CONSTRAINT " + d.FirstKey + " TO " + d.SecondKey)
	}

	if d.Suffix != "" {
		sql.WriteString(" " + d.Suffix)
	}

	sqlStr = sql.String()
	return
}
