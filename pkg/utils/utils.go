package utils

import "self_initializing_fake/internal/model"

func IncludesKey(hey model.Header, needle interface{}) bool {
	for h := range hey {
		if h == needle {
			return true
		}
	}
	return false
}
