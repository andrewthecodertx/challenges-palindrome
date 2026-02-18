package main

import (
	"strings"
)

func IsPalindrome(s string) bool {
	cleaned := strings.ToLower(s)
	var result strings.Builder
	for _, c := range cleaned {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			result.WriteRune(c)
		}
	}
	s = result.String()
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
