package bob

import (
	"strconv"
	"strings"
)

const (
	// Question is the format used in MySQL
	Question = "?"
	// Dollar is the format used in PostgreSQL
	Dollar = "$"
	// Colon is the format used in Oracle Database, but here I implemented it wrong.
	// I will either fix it or remove it in the future.
	Colon = ":"
	// AtP comes in the documentation of Squirrel but I don't know what database uses it.
	AtP = "@p"
)

// PlaceholderFormat is an interface for placeholder formattings.
type PlaceholderFormat interface {
	ReplacePlaceholders(sql string) (string, error)
}

// ReplacePlaceholder converts default placeholder format to a specific format.
func ReplacePlaceholder(sql string, format string) string {
	if format == "" {
		format = Question
	}

	if format == Dollar || format == Colon || format == AtP {
		separate := strings.SplitAfter(sql, "?")
		for i := 0; i < len(separate); i++ {
			separate[i] = strings.Replace(separate[i], "?", format+strconv.Itoa(i+1), 1)
		}
		return strings.Join(separate, "")
	}

	return strings.ReplaceAll(sql, "?", format)
}
