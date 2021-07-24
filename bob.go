package bob

import (
	"errors"

	"github.com/lann/builder"
)

// ErrEmptyTable is a common database/sql error if a table is empty or no rows is returned by the query.
var ErrEmptyTable = errors.New("sql: no rows in result set")
// ErrEmptyTable is a common pgx error if a table is empty or no rows is returned by the query.
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

// CreateTableIfNotExists creates a table with CreateBuilder interface, if the table doesn't exists.
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

// DropTable drops (delete contents & remove) a table from the database.
func (b BobBuilderType) DropTable(table string) DropBuilder {
	return DropBuilder(b).DropTable(table)
}

// DropTable drops (delete contents & remove) a table from the database if the table exists.
func (b BobBuilderType) DropTableIfExists(table string) DropBuilder {
	return DropBuilder(b).DropTable(table).IfExists()
}

// RenameTable simply renames an exisisting table.
func (b BobBuilderType) RenameTable(from, to string) RenameBuilder {
	return RenameBuilder(b).From(from).To(to)
}

// Truncate performs TRUNCATE function. It deletes all contents from a table but not deleting the table.
func (b BobBuilderType) Truncate(table string) TruncateBuilder {
	return TruncateBuilder(b).Truncate(table)
}

// BobStmtBuilder is the parent builder for BobBuilderType
var BobStmtBuilder = BobBuilderType(builder.EmptyBuilder)

// CreateTable creates a table with CreateBuilder interface.
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