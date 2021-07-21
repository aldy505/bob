package bob

import (
	"errors"

	"github.com/lann/builder"
)

var ErrEmptyTable = errors.New("sql: no rows in result set")
var ErrEmptyTablePgx = errors.New("no rows in result set")

// BobBuilderType is the type for BobBuilder
type BobBuilderType builder.Builder

// BobBuilder interface wraps the ToSql method
type BobBuilder interface {
	ToSql() (string, []interface{}, error)
}

// CreateTable creates a table with CreateBuilder interface
func (b BobBuilderType) CreateTable(table string) CreateBuilder {
	return CreateBuilder(b).Name(table)
}

func (b BobBuilderType) CreateTableIfNotExists(table string) CreateBuilder {
	return CreateBuilder(b).Name(table).IfNotExists()
}

// HasTable checks if a table exists with HasBuilder interface
func (b BobBuilderType) HasTable(table string) HasBuilder {
	return HasBuilder(b).HasTable(table)
}

// HasColumn checks if a column exists with HasBuilder interface
func (b BobBuilderType) HasColumn(column string) HasBuilder {
	return HasBuilder(b).HasColumn(column)
}

// BobStmtBuilder is the parent builder for BobBuilderType
var BobStmtBuilder = BobBuilderType(builder.EmptyBuilder)

// CreateTable creates a table with CreateBuilder interface
func CreateTable(table string) CreateBuilder {
	return BobStmtBuilder.CreateTable(table)
}

// CreateTableIfNotExists creates a table with CreateBuilder interface, if the table doesn't exists
func CreateTableIfNotExists(table string) CreateBuilder {
	return BobStmtBuilder.CreateTableIfNotExists(table)
}

// HasTable checks if a table exists with HasBuilder interface
func HasTable(table string) HasBuilder {
	return BobStmtBuilder.HasTable(table)
}

// HasColumn checks if a column exists with HasBuilder interface
func HasColumn(col string) HasBuilder {
	return BobStmtBuilder.HasColumn(col)
}
