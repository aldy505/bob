# Bob - SQL Query Builder

[![Go Reference](https://pkg.go.dev/badge/github.com/aldy505/bob.svg)](https://pkg.go.dev/github.com/aldy505/bob) [![Go Report Card](https://goreportcard.com/badge/github.com/aldy505/bob)](https://goreportcard.com/report/github.com/aldy505/bob) ![GitHub](https://img.shields.io/github/license/aldy505/bob) [![CodeFactor](https://www.codefactor.io/repository/github/aldy505/bob/badge)](https://www.codefactor.io/repository/github/aldy505/bob) [![codecov](https://codecov.io/gh/aldy505/bob/branch/master/graph/badge.svg?token=Noeexg5xEJ)](https://codecov.io/gh/aldy505/bob) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/9b78970127c74c1a923533e05f65848d)](https://www.codacy.com/gh/aldy505/bob/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=aldy505/bob&amp;utm_campaign=Badge_Grade) [![Build test](https://github.com/aldy505/bob/actions/workflows/build.yml/badge.svg)](https://github.com/aldy505/bob/actions/workflows/build.yml) [![Test and coverage](https://github.com/aldy505/bob/actions/workflows/coverage.yml/badge.svg)](https://github.com/aldy505/bob/actions/workflows/coverage.yml)

Think of this as an extension of [Squirrel](https://github.com/Masterminds/squirrel) with functionability like [Knex](https://knexjs.org/). I still use Squirrel for other types of queries (insert, select, and all that), but I needed some SQL builder for create table and some other stuffs.

Oh, and of course, heavily inspired by [Bob the Builder](https://en.wikipedia.org/wiki/Bob_the_Builder).

```go
import "github.com/aldy505/bob"
```

## Usage

It's not ready for large-scale production yet (I've already using it on one of my projects). But, the API is probably close to how you'd do things on Squirrel. 

### Create a table

```go
import "github.com/aldy505/bob"

func main() {
  // Note that CREATE TABLE don't return args params.
  sql, _, err := bob.
    CreateTable("tableName").
    // The first parameter is the column's name.
    // The second parameters and so on forth are extras.
    StringColumn("id", "NOT NULL", "PRIMARY KEY", "AUTOINCREMENT").
    StringColumn("email", "NOT NULL", "UNIQUE").
    // See the list of available column definition type through pkg.go.dev or scroll down below.
    TextColumn("password").
    ToSql()
  if err != nil {
    // handle your error
  }
}
```

Another builder of `bob.CreateTableIfNotExists()` is also available.

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

### Placeholder format

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
* `bob.Question` - `INSERT INTO "users" (name) VALUES (?)`
* `bob.Dollar` - `INSERT INTO "users" (name) VALUES ($1)`
* `bob.Colon` - `INSERT INTO "users" (name) VALUES (:1)`
* `bob.AtP` - `INSERT INTO "users" (name) VALUES (@p1)`


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

  // Check if a table is exists
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
    // Note that this will return multiple query in a single string.
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

* `bob.CreateTable(tableName)` - Basic SQL create table
* `bob.CreateTableIfNotExists(tableName)` - Create table if not exists
* `bob.HasTable(tableName)` - Checks if column exists (return error if false, check example above for error handling)
* `bob.HasColumn(columnName)` - Check if a column exists on current table

### TODO

Meaning these are some ideas for the future development of Bob.

* `bob.DropTable(tableName)` - Drop a table (`drop table "users"`)
* `bob.DropTableIfExists(tableName)` - Drop a table if exists (`drop table if exists "users"`)
* `bob.RenameTable(tableName)` - Rename a table (`rename table "users" to "old_users"`)
* `bob.Truncate(tableName)` - Truncate a table (`truncate "users"`)
* `bob.Upsert(tableName)` - UPSERT function (`insert into "users" ("name", "email") values (?, ?) on duplicate key update email = ?`)
* `bob.ExecWith()` - Just like Squirrel's [ExecWith](https://pkg.go.dev/github.com/Masterminds/squirrel?utm_source=godoc#ExecWith)
* `bob.Count(tableName, columnName)` - Count query (`select count("active") from "users"`)

## Contributing

Contributions are always welcome! As long as you add a test for your changes.

## License

Bob is licensed under [MIT license](./LICENSE)