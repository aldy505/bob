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

func (b BobBuilderType) Upsert(table string, dialect int) UpsertBuilder {
	return UpsertBuilder(b).dialect(dialect).into(table)
}

func (b BobBuilderType) DropColumn(table, column string) AlterBuilder {
	return AlterBuilder(b).whatToAlter(alterDropColumn).tableName(table).firstKey(column)
}

func (b BobBuilderType) DropConstraint(table, constraint string) AlterBuilder {
	return AlterBuilder(b).whatToAlter(alterDropConstraint).tableName(table).firstKey(constraint)
}

func (b BobBuilderType) RenameColumn(table, from, to string) AlterBuilder {
	return AlterBuilder(b).whatToAlter(alterRenameColumn).tableName(table).firstKey(from).secondKey(to)
}

func (b BobBuilderType) RenameConstraint(table, from, to string) AlterBuilder {
	return AlterBuilder(b).whatToAlter(alterRenameConstraint).tableName(table).firstKey(from).secondKey(to)
}

// BobStmtBuilder is the parent builder for BobBuilderType
var BobStmtBuilder = BobBuilderType(builder.EmptyBuilder)

// CreateTable creates a table with CreateBuilder interface.
// Refer to README for available column definition types.
//
//      // Note that CREATE TABLE doesn't returns args params.
//      sql, _, err := bob.
//        CreateTable("tableName").
//        // The first parameter is the column's name.
//        // The second parameters and so on forth are extras.
//        StringColumn("id", "NOT NULL", "PRIMARY KEY", "AUTOINCREMENT").
//        StringColumn("email", "NOT NULL", "UNIQUE").
//        // See the list of available column definition types through pkg.go.dev or README.
//        TextColumn("password").
//        // Or add your custom type.
//        AddColumn(bob.ColumnDef{Name: "tableName", Type: "customType", Extras: []string{"NOT NULL"}}).
//        ToSql()
//      if err != nil {
//      // handle your error
//      }
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
//
//       // MySQL example:
//       sql, args, err := bob.
//         // Notice that you should give database dialect on the second params.
//         // Available database dialect are MySQL, PostgreSQL, SQLite, and MSSQL.
//         Upsert("users", bob.MySQL).
//         Columns("name", "email", "age").
//         // You could do multiple Values() call, but I'd suggest to not do it.
//         // Because this is an upsert function, not an insert one.
//         Values("Thomas Mueler", "tmueler@something.com", 25).
//         Replace("age", 25).
//         PlaceholderFormat(bob.Question).
//         ToSql()
//
//       // Another example for PostgreSQL:
//       sql, args, err = bob.
//         Upsert("users", bob.PostgreSQL).
//         Columns("name", "email", "age").
//         Values("Billy Urtha", "billu@something.com", 30).
//         Key("email").
//         Replace("age", 40).
//         PlaceholderFormat(bob.Dollar).
//         ToSql()
//
//       // One more time, for MSSQL / SQL Server:
//       sql, args, err = bob.
//         Upsert("users", bob.MSSQL).
//         Columns("name", "email", "age").
//         Values("George Rust", "georgee@something.com", 19).
//         Key("email", "georgee@something.com").
//         Replace("age", 18).
//         PlaceholderFormat(bob.AtP).
//         ToSql()
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

func DropColumn(table, column string) AlterBuilder {
	return BobStmtBuilder.DropColumn(table, column)
}

func DropConstraint(table, constraint string) AlterBuilder {
	return BobStmtBuilder.DropConstraint(table, constraint)
}

func RenameColumn(table, from, to string) AlterBuilder {
	return BobStmtBuilder.RenameColumn(table, from, to)
}

func RenameConstraint(table, from, to string) AlterBuilder {
	return BobStmtBuilder.RenameConstraint(table, from, to)
}
