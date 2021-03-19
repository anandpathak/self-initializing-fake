package utils

func IncludesKey(hey map[string]interface{}, key string, value interface{}) bool {
	for h := range hey {
		if h == key {
			return true
		}
	}
	return false
}
