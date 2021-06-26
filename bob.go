package bob

import "github.com/lann/builder"

// BobBuilderType is the type for BobBuilder
type BobBuilderType builder.Builder

// BobBuilder interface wraps the ToSQL method
type BobBuilder interface {
	ToSQL() (string, []interface{}, error)
}

// CreateTable creates a table with CreateBuilder interface
func (b BobBuilderType) CreateTable(table string) CreateBuilder {
	return CreateBuilder(b).Name(table)
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

// HasTable checks if a table exists with HasBuilder interface
func HasTable(table string) HasBuilder {
	return BobStmtBuilder.HasTable(table)
}

// HasColumn checks if a column exists with HasBuilder interface
func HasColumn(col string) HasBuilder {
	return BobStmtBuilder.HasColumn(col)
}
