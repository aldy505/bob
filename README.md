# Bob - SQL Query Builder

[![Go Reference](https://pkg.go.dev/badge/github.com/aldy505/bob.svg)](https://pkg.go.dev/github.com/aldy505/bob) [![Go Report Card](https://goreportcard.com/badge/github.com/aldy505/bob)](https://goreportcard.com/report/github.com/aldy505/bob) ![GitHub](https://img.shields.io/github/license/aldy505/bob) [![CodeFactor](https://www.codefactor.io/repository/github/aldy505/bob/badge)](https://www.codefactor.io/repository/github/aldy505/bob) [![codecov](https://codecov.io/gh/aldy505/bob/branch/master/graph/badge.svg?token=Noeexg5xEJ)](https://codecov.io/gh/aldy505/bob) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/9b78970127c74c1a923533e05f65848d)](https://www.codacy.com/gh/aldy505/bob/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=aldy505/bob&amp;utm_campaign=Badge_Grade) [![Build test](https://github.com/aldy505/bob/actions/workflows/build.yml/badge.svg)](https://github.com/aldy505/bob/actions/workflows/build.yml) [![Test and coverage](https://github.com/aldy505/bob/actions/workflows/coverage.yml/badge.svg)](https://github.com/aldy505/bob/actions/workflows/coverage.yml)

I really need a create table SQL builder, and I can't find one. So, like everything else, I made one. Heavily inspired by [Squirrel](https://github.com/Masterminds/squirrel) and [Knex](https://knexjs.org/).

Oh, and of course, heavily inspired by Bob the Builder.

## Usage

It's not ready for production yet. But, the API is probably close to how you'd do things on Squirrel. This is an example for using with pgx.

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
  sql, args, err = bob.HasTable("users").PlaceholderFormat(bob.Dollar).ToSQL()
	if err != nil {
		log.Fatal(err)
	}

  var hasTableUsers bool
  err = db.QueryRow(context.Background(), sql, args...).Scan(&hasTableUsers)
  if err != nil {
    if err.Error() == "no rows in result set" {
			hasTableUsers = false
		} else {
			log.Fatal(err)
		}
  }

  if !hasTableUsers {
    // Create "users" table
    // Note that this will return multiple query in a single string.
    sql, _, err := bob.CreateTable("users").
      Columns("id", "email", "name", "password", "date").
      Types("varchar(36)", "varchar(255)", "varchar(255)", "text", "date").
      Primary("id").
      Unique("email")
      ToSQL()
    if err != nil {
      log.Fatal(err)
    }

    // If you don't do this, you will get the error:
    // ERROR: cannot insert multiple commands into a prepared statement (SQLSTATE 42601)
    splitQuery := strings.Split(sql, ";")
    for i := range splitQuery {
      _, err = db.Query(context.Background(), splitQuery[i])
			if err != nil {
        log.Fatal(err)
			}
    }
  }
}
```