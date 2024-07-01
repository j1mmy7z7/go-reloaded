package goreloaded_test

import (
	"testing"

	"goreloaded"
)

func TestVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test 1",
			input:    "This is a amazing Experience",
			expected: "This is an amazing Experience",
		},
		{
			name:     "Test 2",
			input:    "What a odd orange",
			expected: "What an odd orange",
		},
		{
			name:     "Test 3",
			input:    "A expert in lying",
			expected: "An expert in lying",
		},
		{
			name:     "Test 4",
			input:    "This is a Anchor",
			expected: "This is an Anchor",
		},
		{
			name:     "Test 5",
			input:    "its been a hour",
			expected: "its been an hour",
		},
		{
			name:     "Test 6",
			input:    "Its been a Hour",
			expected: "Its been an Hour",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := goreloaded.SortVowels(test.input)
			want := test.expected

			if want != got {
				t.Errorf("Expected '%s' but got '%s' ", want, got)
			}
		})
	}
}
