package util

// createArgs should create an argument []interface{} for SQL query
// I'm using the idiot approach for creating args
func CreateArgs(keys ...interface{}) []interface{} {
	var args []interface{}
	for _, v := range keys {
		if v == "" {
			continue
		}
		args = append(args, v)
	}
	return args
}
