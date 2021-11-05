package bob

// createArgs should create an argument []interface{} for SQL query
// I'm using the idiot approach for creating args
func createArgs(keys ...interface{}) []interface{} {
	var args []interface{}
	for _, v := range keys {
		if v == "" {
			continue
		}
		args = append(args, v)
	}
	return args
}

// isIn checks if an array have a value
func isIn(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

// findPosition search for value position on an array
func findPosition(arr []string, value string) int {
	for i, item := range arr {
		if item == value {
			return i
		}
	}
	return -1
}
