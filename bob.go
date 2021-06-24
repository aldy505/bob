package bob

import "github.com/lann/builder"

type BobBuilderType builder.Builder

type BobBuilder interface {
	ToSql() (string, []interface{}, error)
}

func (b BobBuilderType) CreateTable(table string) CreateBuilder {
	return CreateBuilder(b).Name(table)
}

var BobStmtBuilder = BobBuilderType(builder.EmptyBuilder)

func CreateTable(table string) CreateBuilder {
	return BobStmtBuilder.CreateTable(table)
}
