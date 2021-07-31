package bob

import (
	"bytes"
	"errors"
	"strings"

	"github.com/lann/builder"
)

type UpsertBuilder builder.Builder

type upsertData struct {
	Dialect     int
	Into        string
	Columns     []string
	Values      [][]interface{}
	Key         []interface{}
	Replace     [][]interface{}
	Placeholder string
}

func init() {
	builder.Register(UpsertBuilder{}, upsertData{})
}

// dialect specifies database dialect used.
func (u UpsertBuilder) dialect(db int) UpsertBuilder {
	return builder.Set(u, "Dialect", db).(UpsertBuilder)
}

// Table sets which table to be dropped.
func (u UpsertBuilder) into(name string) UpsertBuilder {
	return builder.Set(u, "Into", name).(UpsertBuilder)
}

// Columns sets the columns for the data to be inserted.
func (u UpsertBuilder) Columns(columns ...string) UpsertBuilder {
	return builder.Extend(u, "Columns", columns).(UpsertBuilder)
}

// Values sets the values in relation with the columns.
// Please not that only string, int, and bool type are supported.
// Inputting other types other than those might result in your SQL not working properly.
func (u UpsertBuilder) Values(values ...interface{}) UpsertBuilder {
	return builder.Append(u, "Values", values).(UpsertBuilder)
}

// Key specifies which key to be checked on conflict.
// Must be used on PostgreSQL and SQLite.
func (u UpsertBuilder) Key(key ...interface{}) UpsertBuilder {
	var value interface{}
	column := key[0]
	if len(key) > 1 && key[0] != nil {
		value = key[1]
	} else {
		value = ""
	}
	return builder.Extend(u, "Key", []interface{}{column, value}).(UpsertBuilder)
}

// Replace sets the column and value respectively for the data to be changed on
// a specific row.
func (u UpsertBuilder) Replace(column interface{}, value interface{}) UpsertBuilder {
	return builder.Append(u, "Replace", []interface{}{column, value}).(UpsertBuilder)
}

// PlaceholderFormat changes the default placeholder (?) to desired placeholder.
func (u UpsertBuilder) PlaceholderFormat(f string) UpsertBuilder {
	return builder.Set(u, "Placeholder", f).(UpsertBuilder)
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (u UpsertBuilder) ToSql() (string, []interface{}, error) {
	data := builder.GetStruct(u).(upsertData)
	return data.ToSql()
}

// ToSql returns 3 variables filled out with the correct values based on bindings, etc.
func (d *upsertData) ToSql() (sqlStr string, args []interface{}, err error) {
	if len(d.Into) == 0 || d.Into == "" {
		err = errors.New("upsert statements must specify a table")
		return
	}

	if len(d.Columns) == 0 || d.Columns[0] == "" {
		err = errors.New("upsert statement must have at least one column")
		return
	}

	if len(d.Values) == 0 {
		err = errors.New("upsert statements must have at least one set of values")
		return
	}

	if len(d.Replace) == 0 {
		err = errors.New("upsert statement must have at least one key value pair to be replaced")
		return
	}

	sql := &bytes.Buffer{}

	if d.Dialect == MSSQL {
		if len(d.Key) == 0 {
			err = errors.New("unique key and value must be provided for MS SQL")
			return
		}

		sql.WriteString("IF NOT EXISTS (SELECT * FROM \"" + d.Into + "\" WHERE \"" + d.Key[0].(string) + "\" = ?) ")
		args = append(args, d.Key[1])
	}

	sql.WriteString("INSERT INTO ")
	sql.WriteString("\"" + d.Into + "\"")
	sql.WriteString(" ")

	var columns []string
	for _, v := range d.Columns {
		columns = append(columns, "\""+v+"\"")
	}

	sql.WriteString("(")
	sql.WriteString(strings.Join(columns, ", "))
	sql.WriteString(") ")

	sql.WriteString("VALUES ")

	var values []string
	for i := 0; i < len(d.Values); i++ {
		var tempValues []string
		for _, v := range d.Values[i] {
			args = append(args, v)
			tempValues = append(tempValues, "?")
		}
		values = append(values, "("+strings.Join(tempValues, ", ")+")")
	}

	sql.WriteString(strings.Join(values, ", "))
	sql.WriteString(" ")

	var replaces []string
	for i := 0; i < len(d.Replace); i++ {
		args = append(args, d.Replace[i][1])
		replace := "\"" + d.Replace[i][0].(string) + "\" = ?"
		replaces = append(replaces, replace)
	}

	if d.Dialect == MySQL {
		// INSERT INTO table (col) VALUES (values) ON DUPLICATE KEY UPDATE col = value

		sql.WriteString("ON DUPLICATE KEY UPDATE ")
		sql.WriteString(strings.Join(replaces, ", "))
	} else if d.Dialect == PostgreSQL || d.Dialect == SQLite {
		// INSERT INTO players (user_name, age) VALUES('steven', 32) ON CONFLICT(user_name) DO UPDATE SET age=excluded.age;

		if len(d.Key) == 0 {
			err = errors.New("unique key must be provided for PostgreSQL and SQLite")
			return
		}

		sql.WriteString("ON CONFLICT ")
		sql.WriteString("(\"" + d.Key[0].(string) + "\") ")
		sql.WriteString("DO UPDATE SET ")
		sql.WriteString(strings.Join(replaces, ", "))

	} else if d.Dialect == MSSQL {
		// IF NOT EXISTS (SELECT * FROM dbo.Table1 WHERE ID = @ID)
		//    INSERT INTO dbo.Table1(ID, Name, ItemName, ItemCatName, ItemQty)
		//    VALUES(@ID, @Name, @ItemName, @ItemCatName, @ItemQty)
		// ELSE
		//    UPDATE dbo.Table1
		//    SET Name = @Name,
		//        ItemName = @ItemName,
		//        ItemCatName = @ItemCatName,
		//        ItemQty = @ItemQty
		//    WHERE ID = @ID

		sql.WriteString("ELSE ")
		sql.WriteString("UPDATE \"" + d.Into + "\" SET ")
		sql.WriteString(strings.Join(replaces, ", "))
		sql.WriteString(" WHERE \"" + d.Key[0].(string) + "\" = ?")
		args = append(args, d.Key[1])

	} else {
		err = ErrDialectNotSupported
		return
	}

	sql.WriteString(";")

	sqlStr = ReplacePlaceholder(sql.String(), d.Placeholder)
	return
}
