# Bob - SQL Query Builder

[![Go Reference][pkg-go-dev-badge]][pkg-go-dev-link]
[![Go Report Card][go-report-badge]][go-report-link]
![GitHub][license-badge]
[![CodeFactor][codefactor-badge]][codefactor-link]
[![codecov][codecov-badge]][codecov-link]
[![Codacy Badge][codacy-badge]][codacy-link]
[![Test and coverage][actions-badge]][actions-link]

Bob is an SQL builder library initially made as an extension for [Squirrel][squirrel-url]
with functionality like [Knex][knex-url] (from the Node.js world). Squirrel itself
doesn't provide other types of queries for creating a table, upsert,
and some other things. Bob is meant to fill those gaps.

The difference between Bob and Squirrel is that Bob is solely a query builder.
The users have to execute and manage the SQL connection themselves.
Meaning there is no ExecWith() function implemented in Bob, as you can
find it on Squirrel.

The purpose of an SQL query builder is to prevent any typo or mistypes
on the SQL queries. Although also with that reason, Bob might not always
have the right query for you, depending on what you are doing with the
SQL query. It might sometimes be better for you to write the SQL query
yourself, if your problem is specific and needs some micro-tweaks.

With that being said, I hope you enjoy using Bob and consider starring or
reporting any issues regarding the usage of Bob in your projects.

Oh, and of course, heavily inspired by [Bob the Builder][bob-wikipedia].

## Usage

```go
import "github.com/aldy505/bob"
```

Like any other Go projects when you're using Go modules, just put that
text right there on the top of your projects, do `go mod tidy` and
you are good to go.

Either way, I'm not 100% confident enough to say that this thing is
production ready. But, the way I see it, it's good enough to be used
on a production-level application. In fact, I'm using it on one of my
current projects that's getting around 100-200 hits per day.

If you have any feature request or improvement ideas for the project,
please kindly open an issue

### Create a table

```go
import "github.com/aldy505/bob"

func main() {
  // Note that CREATE TABLE doesn't returns args params.
  sql, _, err := bob.
    CreateTable("tableName").
    // The first parameter is the column's name.
    // The second parameter and so on forth are extras.
    StringColumn("id", "NOT NULL", "PRIMARY KEY", "AUTOINCREMENT").
    StringColumn("email", "NOT NULL", "UNIQUE").
    // See the list of available column definition types through pkg.go.dev or scroll down below.
    TextColumn("password").
    // Or add your custom type.
    AddColumn(bob.ColumnDef{Name: "tableName", Type: "customType", Extras: []string{"NOT NULL"}}).
    ToSql()
  if err != nil {
    // handle your error
  }
}
```

Available column definition types (please be aware that some only works on certain database):

- `StringColumn()` - Default to `VARCHAR(255)`
- `TextColumn()` - Default to `TEXT`
- `UUIDColumn()` - Defaults to `UUID`
- `BooleanColumn()` - Defaults to `BOOLEAN`
- `IntegerColumn()` - Defaults to `INTEGER`. Postgres and SQLite only.
- `IntColumn()` - Defaults to `INT`. MySQL and MSSQL only.
- `RealColumn()` - Defaults to `REAL`. Postgres, MSSQL, and SQLite only.
- `FloatColumn()` - Defaults to `FLOAT`. Postgres and SQLite only.
- `DateTimeColumn()` - Defaults to `DATETIME`.
- `TimeStampColumn()` - Defaults to `TIMESTAMP`.
- `TimeColumn()` - Defaults to `TIME`.
- `DateColumn()` - Defaults to `DATE`.
- `JSONColumn()` - Defaults to `JSON`. MySQL and Postgres only.
- `JSONBColumn()` - Defaults to `JSONB`. Postgres only.
- `BlobColumn()` - Defaults to `BLOB`. MySQL and SQLite only.

For any other types, please use `AddColumn()`.

Another builder of `bob.CreateTableIfNotExists()` is also available.

### Create index

```go
func main() {
  sql, _, err := bob.
    CreateIndex("idx_email").
    On("users").
    // To create a CREATE UNIQUE INDEX ...
    Unique().
    // Method "Spatial()" and "FullText()" are also available.
    // You can specify as many columns as you like.
    Columns(bob.IndexColumn{Name: "email", Collate: "DEFAULT", Extras: []string{"ASC"}}).
    ToSql()
  if err != nil {
    log.Fatal(err)
  }
}
```

Another builder of `bob.CreateIndexIfNotExists()` is also available.

### Check if a table exists

```go
func main() {
  sql, args, err := bob.HasTable("users").ToSql()
  if err != nil {
    log.Fatal(err)
  }
}
```

### Check if a column exists

```go
func main() {
  sql, args, err := bob.HasColumn("email").ToSql()
  if err != nil {
    log.Fatal(err)
  }
}
```

### Drop table

```go
func main() {
  sql, _, err := bob.DropTable("users").ToSql()
  if err != nil {
    log.Fatal(err)
  }
  // sql = "DROP TABLE users;"

  sql, _, err = bob.DropTableIfExists("users").ToSql()
  if err != nil {
    log.Fatal(err)
  }
  // sql = "DROP TABLE IF EXISTS users;"

  sql, _, err = bob.DropTable("users").Cascade().ToSql()
  if err != nil {
    log.Fatal(err)
  }
  // sql = "DROP TABLE users CASCADE;"

  sql, _, err = bob.DropTable("users").Restrict().ToSql()
  if err != nil {
    log.Fatal(err)
  }
  // sql = "DROP TABLE users RESTRICT;"
}
```

### Truncate table

```go
func main() {
  sql, _, err := bob.Truncate("users").ToSql()
  if err != nil {
    log.Fatal(err)
  }
}
```

### Rename table

```go
func main() {
  sql, _, err := bob.RenameTable("users", "people").ToSql()
  if err != nil {
    log.Fatal(err)
  }
}
```

### Upsert

```go
func main() {
  sql, args, err := bob.
    // Notice that you should give database dialect on the second params.
    // Available database dialect are MySQL, PostgreSQL, SQLite, and MSSQL.
    Upsert("users", bob.MySQL).
    Columns("name", "email", "age").
    // You could do multiple Values() call, but I'd suggest to not do it.
    // Because this is an upsert function, not an insert one.
    Values("Thomas Mueler", "tmueler@something.com", 25).
    Replace("age", 25).
    ToSql()

  // Another example for PostgreSQL
  sql, args, err = bob.
    Upsert("users", bob.PostgreSQL).
    Columns("name", "email", "age").
    Values("Billy Urtha", "billu@something.com", 30).
    Key("email").
    Replace("age", 40).
    ToSql()

  // One more time, for MSSQL / SQL Server.
  sql, args, err = bob.
    Upsert("users", bob.MSSQL).
    Columns("name", "email", "age").
    Values("George Rust", "georgee@something.com", 19).
    Key("email", "georgee@something.com").
    Replace("age", 18).
    ToSql()
}
```

### Placeholder format / Dialect

Default placeholder is a question mark (MySQL-like). If you want to change it, simply use something like this:

```go
func main() {
  // Option 1
  sql, args, err := bob.HasTable("users").PlaceholderFormat(bob.Dollar).ToSql()
  if err != nil {
    log.Fatal(err)
  }

  // Option 2
  sql, args, err = bob.HasTable("users").ToSql()
  if err != nil {
    log.Fatal(err)
  }
  correctPlaceholder := bob.ReplacePlaceholder(sql, bob.Dollar)
}
```

Available placeholder formats:

- `bob.Question` - `INSERT INTO "users" (name) VALUES (?)`
- `bob.Dollar` - `INSERT INTO "users" (name) VALUES ($1)`
- `bob.Colon` - `INSERT INTO "users" (name) VALUES (:1)`
- `bob.AtP` - `INSERT INTO "users" (name) VALUES (@p1)`

### With pgx (PostgreSQL)

```go
import (
  "context"
  "log"
  "strings"

  "github.com/aldy505/bob"
  "github.com/jackc/pgx/v4"
)

func main() {
  db := pgx.Connect()

  // Check if a table exists
  sql, args, err = bob.HasTable("users").PlaceholderFormat(bob.Dollar).ToSql()
  if err != nil {
    log.Fatal(err)
  }

  var hasTableUsers bool
  err = db.QueryRow(context.Background(), sql, args...).Scan(&hasTableUsers)
  if err != nil {
    if err == bob.ErrEmptyTablePg {
      hasTableUsers = false
    } else {
      log.Fatal(err)
    }
  }

  if !hasTableUsers {
    // Create "users" table
    sql, _, err := bob.
      CreateTable("users").
      IntegerColumn("id", "PRIMARY KEY", "SERIAL").
      StringColumn("name", "NOT NULL").
      TextColumn("password", "NOT NULL").
      DateColumn("created_at").
      ToSql()
    if err != nil {
      log.Fatal(err)
    }

    _, err = db.Query(context.Background(), splitQuery[i])
    if err != nil {
      log.Fatal(err)
    }

    // Create another table, this time with CREATE TABLE IF NOT EXISTS
    sql, _, err := bob.
      CreateTableIfNotExists("inventory").
      UUIDColumn("id", "PRIMARY KEY").
      IntegerColumn("userID", "FOREIGN KEY REFERENCES users(id)").
      JSONColumn("items").
      IntegerColumn("quantity").
      ToSql()
    if err != nil {
      log.Fatal(err)
    }

    _, err = db.Query(context.Background(), inventoryQuery[i])
    if err != nil {
      log.Fatal(err)
    }
  }
}
```

## Features

- `bob.CreateTable(tableName)` - Basic SQL create table
- `bob.CreateTableIfNotExists(tableName)` - Create table if not exists
- `bob.CreateIndex(indexName)` - Basic SQL create index
- `bob.CreateIndexIfNotExists(tableName)` - Create index if not exists
- `bob.HasTable(tableName)` - Checks if column exists (return error if false, check example above for error handling)
- `bob.HasColumn(columnName)` - Check if a column exists on current table
- `bob.DropTable(tableName)` - Drop a table (`drop table "users"`)
- `bob.DropTableIfExists(tableName)` - Drop a table if exists (`drop table if exists "users"`)
- `bob.RenameTable(currentTable, desiredName)` - Rename a table (`rename table "users" to "people"`)
- `bob.Truncate(tableName)` - Truncate a table (`truncate "users"`)
- `bob.Upsert(tableName, dialect)` - UPSERT function (`insert into "users" ("name", "email") values (?, ?) on duplicate key update email = ?`)

## Contributing

Contributions are always welcome! As long as you add a test for your changes.

## License

Bob is licensed under [MIT license](./LICENSE)

[squirrel-url]: https://github.com/Masterminds/squirrel
[knex-url]: https://knexjs.org/
[bob-wikipedia]: https://en.wikipedia.org/wiki/Bob_the_Builder
[pkg-go-dev-badge]: https://pkg.go.dev/badge/github.com/aldy505/bob.svg
[pkg-go-dev-link]: https://pkg.go.dev/github.com/aldy505/bob
[go-report-badge]: https://goreportcard.com/badge/github.com/aldy505/bob
[go-report-link]: https://goreportcard.com/report/github.com/aldy505/bob
[license-badge]: https://img.shields.io/github/license/aldy505/bob
[codefactor-link]: https://www.codefactor.io/repository/github/aldy505/bob
[codefactor-badge]: https://www.codefactor.io/repository/github/aldy505/bob/badge
[codecov-badge]: https://codecov.io/gh/aldy505/bob/branch/master/graph/badge.svg?token=Noeexg5xEJ
[codecov-link]: https://codecov.io/gh/aldy505/bob
[codacy-badge]: https://app.codacy.com/project/badge/Grade/9b78970127c74c1a923533e05f65848d
[codacy-link]: https://www.codacy.com/gh/aldy505/bob/dashboard?utm_source=github.com&utm_medium=referral&utm_content=aldy505/bob&utm_campaign=Badge_Grade
[actions-badge]: https://github.com/aldy505/bob/actions/workflows/coverage.yml/badge.svg
[actions-link]: https://github.com/aldy505/bob/actions/workflows/coverage.yml
