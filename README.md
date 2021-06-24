# Bob - SQL Query Builder

I really need a create table SQL builder, and I can't find one. So, like everything else, I made one. Heavily inspired by [Squirrel](https://github.com/Masterminds/squirrel) and [Knex](https://knexjs.org/).

Oh, and of course, heavily inspired by Bob the Builder.

## Usage

It's not ready for production (yet). But, the API is probably close to how you'd do things on Squirrel.

```go
import "github.com/aldy505/bob"

func main() {
  sql, _, err := bob.CreateTable("users").
    Columns("id", "email", "name", "password", "date").
    Types("varchar(36)", "varchar(255)", "varchar(255)", "text", "date").
    Primary("id").
    Unique("email")
    ToSql()
  if err != nil {
    log.Fatal(err)
  }

  // process SQL with whatever you like it
}
```