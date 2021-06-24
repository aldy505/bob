package bob

import "github.com/lann/builder"

// TODO - The whole file is a todo
// Meant to find two things: HasTable and HasColumn(s)

type HasBuilder builder.Builder

type hasData struct {
	Name        string
	Placeholder PlaceholderFormat
}

func init() {
	builder.Register(HasBuilder{}, hasData{})
}

func (h HasBuilder) HasTable(table string) HasBuilder {
	return builder.Set(h, "Name", table).(HasBuilder)
}

func (h HasBuilder) PlaceholderFormat(f PlaceholderFormat) HasBuilder {
	return builder.Set(h, "Placeholder", f).(HasBuilder)
}
