package utils

import "strings"

func NotIn(key string, slice []string) bool {
	for _, s := range slice {
		if key == s {
			return false
		}
	}
	return true
}

func In(key string, slice []string) bool {
	for _, s := range slice {
		if key == s {
			return true
		}
	}
	return false
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func SplitLast(str, spilt string) string {
	arr := strings.Split(str, spilt)
	return arr[len(arr)-1]
}
