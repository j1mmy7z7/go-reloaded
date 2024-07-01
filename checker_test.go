package goreloaded_test

import (
	"testing"

	"goreloaded"
)

func TestPatterns(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test 1",
			input:    "this is a test (cap, 4) , this (cap) IS A TEST (low, 3).",
			expected: "This Is A Test , This is a test.",
		},
		{
			name:     "Test 2",
			input:    "THIS IS A TEST (low, 3) , THIS (low) is a test (up, 3).",
			expected: "THIS is a test , this IS A TEST.",
		},
		{
			name:     "Test 3",
			input:    "this is (cap, 2) , shocking (up)",
			expected: "This Is , SHOCKING",
		},
		{
			name:     "test 4",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := goreloaded.CheckString(test.input)
			want := test.expected

			if want != got {
				t.Errorf("Expected '%s' but got '%s' ", want, got)
			}
		})
	}
}
