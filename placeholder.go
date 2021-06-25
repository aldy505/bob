package bob

import "strings"

const (
	Question = "?"
	Dollar   = "$"
	Colon    = ":"
	AtP      = "@p"
)

type PlaceholderFormat interface {
	ReplacePlaceholders(sql string) (string, error)
}

// TODO - test this one
func ReplacePlaceholder(sql string, format string) string {
	if format == "" {
		format = Question
	}
	return strings.ReplaceAll(sql, "?", format)
}
