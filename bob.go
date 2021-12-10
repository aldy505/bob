// Bob is an SQL builder library initially made as an extension for Squirrel
// with functionality like Knex (from the Node.js world). Squirrel itself
// doesn't provide other types of queries for creating a table, upsert,
// and some other things. Bob is meant to fill those gaps.
//
// The different between Bob and Squirrel is that Bob is solely a query builder.
// The users have to execute and manage the SQL connection themself.
// Meaning there are no ExecWith() function implemented on Bob, as you can
// find it on Squirrel.
//
// The purpose of an SQL query builder is to prevent any typo or mistypes
// on the SQL queries. Although also with that reason, Bob might not always
// have the right query for you, depending on what you are doing with the
// SQL query. It might sometimes be better for you to write the SQL query
// yourself, if your problem is specific and needs some micro-tweaks.
//
// With that being said, I hope you enjoy using Bob and consider starring or
// reporting any issues regarding the usage of Bob in your projects.
//
// MIT License
//
// Copyright (c) 2021-present Reinaldy Rafli and Bob collaborators
//
package bob

import (
	"errors"

	"github.com/lann/builder"
)

// ErrEmptyTable is a common database/sql error if a table is empty or no rows is returned by the query.
var ErrEmptyTable = errors.New("sql: no rows in result set")

// ErrEmptyTable is a common pgx error if a table is empty or no rows is returned by the query.
var ErrEmptyTablePgx = errors.New("no rows in result set")

// ErrDialectNotSupported tells you whether the dialect is supported or not.
var ErrDialectNotSupported = errors.New("provided database dialect is not supported")

const (
	MySQL int = iota
	PostgreSQL
	SQLite
	MSSQL
)

// BobBuilderType is the type for BobBuilder
type BobBuilderType builder.Builder

// BobBuilder interface wraps the ToSql method
type BobBuilder interface {
	ToSql() (string, []interface{}, error)
}

// CreateTable creates a table with CreateBuilder interface
func (b BobBuilderType) CreateTable(table string) CreateBuilder {
	return CreateBuilder(b).name(table)
}

// CreateTableIfNotExists creates a table with CreateBuilder interface, if the table doesn't exists.
func (b BobBuilderType) CreateTableIfNotExists(table string) CreateBuilder {
	return CreateBuilder(b).name(table).ifNotExists()
}

// CreateIndex creates an index with CreateIndexBuilder interface.
func (b BobBuilderType) CreateIndex(name string) IndexBuilder {
	return IndexBuilder(b).name(name)
}

// CreateIndexIfNotExists creates an index with CreateIndexBuilder interface, if the index doesn't exists.
func (b BobBuilderType) CreateIndexIfNotExists(name string) IndexBuilder {
	return IndexBuilder(b).name(name).ifNotExists()
}

// HasTable checks if a table exists with HasBuilder interface
func (b BobBuilderType) HasTable(table string) HasBuilder {
	return HasBuilder(b).HasTable(table)
}

// HasColumn checks if a column exists with HasBuilder interface
func (b BobBuilderType) HasColumn(column string) HasBuilder {
	return HasBuilder(b).HasColumn(column)
}

// DropTable drops (delete contents & remove) a table from the database.
func (b BobBuilderType) DropTable(table string) DropBuilder {
	return DropBuilder(b).dropTable(table)
}

// DropTable drops (delete contents & remove) a table from the database if the table exists.
func (b BobBuilderType) DropTableIfExists(table string) DropBuilder {
	return DropBuilder(b).dropTable(table).ifExists()
}

// RenameTable simply renames an exisisting table.
func (b BobBuilderType) RenameTable(from, to string) RenameBuilder {
	return RenameBuilder(b).from(from).to(to)
}

// Truncate performs TRUNCATE function. It deletes all contents from a table but not deleting the table.
func (b BobBuilderType) Truncate(table string) TruncateBuilder {
	return TruncateBuilder(b).truncate(table)
}

// Upsert upserts a row into a table.
func (b BobBuilderType) Upsert(table string, dialect int) UpsertBuilder {
	return UpsertBuilder(b).dialect(dialect).into(table)
}

// DropTable drops (delete contents & remove) a table from the database if the table exists.
func (b BobBuilderType) DropColumn(table, column string) AlterBuilder {
	return AlterBuilder(b).whatToAlter(alterDropColumn).tableName(table).firstKey(column)
}

// DropConstraint drops (delete contents & remove) a constraint from the database.
func (b BobBuilderType) DropConstraint(table, constraint string) AlterBuilder {
	return AlterBuilder(b).whatToAlter(alterDropConstraint).tableName(table).firstKey(constraint)
}

// RenameColumn simply renames an exisisting column.
func (b BobBuilderType) RenameColumn(table, from, to string) AlterBuilder {
	return AlterBuilder(b).whatToAlter(alterRenameColumn).tableName(table).firstKey(from).secondKey(to)
}

// RenameConstraint simply renames an exisisting constraint.
func (b BobBuilderType) RenameConstraint(table, from, to string) AlterBuilder {
	return AlterBuilder(b).whatToAlter(alterRenameConstraint).tableName(table).firstKey(from).secondKey(to)
}

// BobStmtBuilder is the parent builder for BobBuilderType
var BobStmtBuilder = BobBuilderType(builder.EmptyBuilder)

// CreateTable creates a table with CreateBuilder interface.
// Refer to README for available column definition types.
func CreateTable(table string) CreateBuilder {
	return BobStmtBuilder.CreateTable(table)
}

// CreateTableIfNotExists creates a table with CreateBuilder interface, if the table doesn't exists.
func CreateTableIfNotExists(table string) CreateBuilder {
	return BobStmtBuilder.CreateTableIfNotExists(table)
}

// HasTable checks if a table exists with HasBuilder interface.
func HasTable(table string) HasBuilder {
	return BobStmtBuilder.HasTable(table)
}

// HasColumn checks if a column exists with HasBuilder interface.
func HasColumn(col string) HasBuilder {
	return BobStmtBuilder.HasColumn(col)
}

// DropTable drops (delete contents & remove) a table from the database.
func DropTable(table string) DropBuilder {
	return BobStmtBuilder.DropTable(table)
}

// DropTable drops (delete contents & remove) a table from the database if the table exists.
func DropTableIfExists(table string) DropBuilder {
	return BobStmtBuilder.DropTableIfExists(table)
}

// RenameTable simply renames an exisisting table.
func RenameTable(from, to string) RenameBuilder {
	return BobStmtBuilder.RenameTable(from, to)
}

// Truncate performs TRUNCATE function. It deletes all contents from a table but not deleting the table.
func Truncate(table string) TruncateBuilder {
	return BobStmtBuilder.Truncate(table)
}

// Upsert performs a UPSERT query with specified database dialect.
// Supported database includes MySQL, PostgreSQL, SQLite and MSSQL.
func Upsert(table string, dialect int) UpsertBuilder {
	return BobStmtBuilder.Upsert(table, dialect)
}

// CreateIndex creates an index with CreateIndexBuilder interface.
func CreateIndex(name string) IndexBuilder {
	return BobStmtBuilder.CreateIndex(name)
}

// CreateIndexIfNotExists creates an index with CreateIndexBuilder interface, if the index doesn't exists.
func CreateIndexIfNotExists(name string) IndexBuilder {
	return BobStmtBuilder.CreateIndexIfNotExists(name)
}

// DropColumn drops (delete contents & remove) a column from the table.
func DropColumn(table, column string) AlterBuilder {
	return BobStmtBuilder.DropColumn(table, column)
}

// DropConstraint drops (delete contents & remove) a constraint from the table.
func DropConstraint(table, constraint string) AlterBuilder {
	return BobStmtBuilder.DropConstraint(table, constraint)
}

// RenameColumn simply renames an exisisting column.
func RenameColumn(table, from, to string) AlterBuilder {
	return BobStmtBuilder.RenameColumn(table, from, to)
}

// RenameConstraint simply renames an exisisting constraint.
func RenameConstraint(table, from, to string) AlterBuilder {
	return BobStmtBuilder.RenameConstraint(table, from, to)
}
