package bob

// TODO

type PlaceholderFormat interface {
	ReplacePlaceholders(sql string) (string, error)
}
