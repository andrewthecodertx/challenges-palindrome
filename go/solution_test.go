package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"", true},
		{"a", true},
		{"ab", false},
	}

	for _, tt := range tests {
		result := IsPalindrome(tt.input)
		if result != tt.expected {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}
