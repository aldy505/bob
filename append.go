package bob

import "io"

func appendToSql(parts []BobBuilder, w io.Writer, sep string, args []interface{}) ([]interface{}, error) {
	for i, p := range parts {
		partSql, partArgs, err := p.ToSql()
		if err != nil {
			return nil, err
		} else if len(partSql) == 0 {
			continue
		}

		if i > 0 {
			_, err := io.WriteString(w, sep)
			if err != nil {
				return nil, err
			}
		}

		_, err = io.WriteString(w, partSql)
		if err != nil {
			return nil, err
		}
		args = append(args, partArgs...)
	}
	return args, nil
}

// createArgs should create an argument []interface{} for SQL query
// I'm using the idiot approach for creating args
func createArgs(keys ...string) []interface{} {
	var args []interface{}
	for _, v := range keys {
		if v == "" {
			continue
		}
		args = append(args, v)
	}
	return args
}
