package util

func IsIn(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

func FindPosition(arr []string, value string) int {
	for i, item := range arr {
		if item == value {
			return i
		}
	}
	return -1
}
