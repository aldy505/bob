package bob

import "github.com/lann/builder"

type BobBuilderType builder.Builder

type BobBuilder interface {
	ToSql() (string, []interface{}, error)
}

func (b BobBuilderType) CreateTable(table string) CreateBuilder {
	return CreateBuilder(b).Name(table)
}

func (b BobBuilderType) HasTable(table string) HasBuilder {
	return HasBuilder(b).HasTable(table)
}

func (b BobBuilderType) HasColumn(column string) HasBuilder {
	return HasBuilder(b).HasColumn(column)
}

var BobStmtBuilder = BobBuilderType(builder.EmptyBuilder)

func CreateTable(table string) CreateBuilder {
	return BobStmtBuilder.CreateTable(table)
}

func HasTable(table string) HasBuilder {
	return BobStmtBuilder.HasTable(table)
}

func HasColumn(col string) HasBuilder {
	return BobStmtBuilder.HasColumn(col)
}
