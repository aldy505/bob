package bob_test

import (
	"fmt"
	"log"

	"github.com/aldy505/bob"
)

func ExampleCreateIndex() {
	sql, _, err := bob.
		CreateIndex("idx_email").
		On("users").
		Unique().
		Columns(bob.IndexColumn{Name: "email", Collate: "DEFAULT", Extras: []string{"ASC"}}).
		ToSql()
	if err != nil {
		fmt.Printf("Handle this error: %v", err)
	}

	fmt.Print(sql)
	// Output: CREATE UNIQUE INDEX idx_email ON users (email COLLATE DEFAULT ASC);
}

func ExampleHasTable() {
	sql, args, err := bob.HasTable("users").ToSql()
	if err != nil {
		fmt.Printf("Handle this error: %v", err)
	}

	fmt.Printf("sql: %s, args: %v", sql, args)
	// Output: sql: SELECT * FROM information_schema.tables WHERE table_name = ? AND table_schema = current_schema();, args: [users]
}

func ExampleHasColumn() {
	sql, args, err := bob.HasColumn("email").ToSql()
	if err != nil {
		fmt.Printf("Handle this error: %v", err)
	}

	fmt.Printf("sql: %s, args: %v", sql, args)
	// Output: sql: SELECT * FROM information_schema.columns WHERE table_name = ? AND table_schema = current_schema();, args: [users]
}

func ExampleDropTable() {
	sql, _, err := bob.DropTable("users").Cascade().ToSql()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)
	// Output: DROP TABLE users CASCADE;
}

func ExampleDropTableIfExists() {
	sql, _, err := bob.DropTableIfExists("users").ToSql()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)
	// Output: DROP TABLE IF EXISTS users;
}

func ExampleTruncate() {
	sql, _, err := bob.Truncate("users").ToSql()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)
	// Output: TRUNCATE TABLE users;
}

func ExampleRenameTable() {
	sql, _, err := bob.RenameTable("users", "people").ToSql()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)
	// Output: ALTER TABLE users RENAME TO people;
}

func ExampleUpsert() {
	// Example for MYSQL
	mysql, myArgs, err := bob.
		// Notice that you should give database dialect on the second params.
		// Available database dialect are MySQL, PostgreSQL, SQLite, and MSSQL.
		Upsert("users", bob.MySQL).
		Columns("name", "email", "age").
		// You could do multiple Values() call, but I'd suggest to not do it.
		// Because this is an upsert function, not an insert one.
		Values("Thomas Mueler", "tmueler@something.com", 25).
		Replace("age", 25).
		ToSql()
	if err != nil {
		fmt.Printf("Handle this error: %v", err)
	}

	// Another example for PostgreSQL
	pgsql, pgArgs, err := bob.
		Upsert("users", bob.PostgreSQL).
		Columns("name", "email", "age").
		Values("Billy Urtha", "billu@something.com", 30).
		Key("email").
		Replace("age", 40).
		ToSql()
	if err != nil {
		fmt.Printf("Handle this error: %v", err)
	}

	// One more time, for MSSQL / SQL Server.
	mssql, msArgs, err := bob.
		Upsert("users", bob.MSSQL).
		Columns("name", "email", "age").
		Values("George Rust", "georgee@something.com", 19).
		Key("email", "georgee@something.com").
		Replace("age", 18).
		ToSql()
	if err != nil {
		fmt.Printf("Handle this error: %v", err)
	}

	fmt.Printf("MySQL: %s, %v\n", mysql, myArgs)
	fmt.Printf("PostgreSQL: %s, %v\n", pgsql, pgArgs)
	fmt.Printf("MSSQL: %s, %v\n", mssql, msArgs)
	// Output:
	// MySQL: INSERT INTO "users" ("name", "email", "age") VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE "age" = ?;, [Thomas Mueler tmueler@something.com 25 25]
	// PostgreSQL: INSERT INTO "users" ("name", "email", "age") VALUES ($1, $2, $3) ON CONFLICT ("email") DO UPDATE SET "age" = $4;, [Billy Urtha billu@something.com 30 40]
	// MSSQL: IF NOT EXISTS (SELECT * FROM "users" WHERE "email" = @p1) INSERT INTO "users" ("name", "email", "age") VALUES (@p2, @p3, @p4) ELSE UPDATE "users" SET "age" = @p5 WHERE "email" = @p6;, [georgee@something.com George Rust georgee@something.com 19 18 georgee@something.com]
}
