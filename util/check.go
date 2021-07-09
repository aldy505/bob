package util

// IsIn checks if an array have a value
func IsIn(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

// FindPosition search for value position on an array
func FindPosition(arr []string, value string) int {
	for i, item := range arr {
		if item == value {
			return i
		}
	}
	return -1
}
