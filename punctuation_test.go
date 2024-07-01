package goreloaded_test

import (
	"testing"

	"goreloaded"
)

func TestPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test 1",
			input:    "This is the first ,test",
			expected: "This is the first, test",
		},
		{
			name:     "Test 2",
			input:    "This is the second  ,  test",
			expected: "This is the second, test",
		},
		{
			name:     "Test 3",
			input:    "This is ,the fourth test ?",
			expected: "This is, the fourth test?",
		},
		{
			name:     "Test 4",
			input:    "This should entail : business and more business .",
			expected: "This should entail: business and more business.",
		},
		{
			name:     "Test 5",
			input:    "Should this work ....",
			expected: "Should this work....",
		},
		{
			name:     "Test 6",
			input:    "Even this !?!?",
			expected: "Even this!?!?",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := goreloaded.Punct(test.input)
			want := test.expected

			if want != got {
				t.Errorf("Expected '%s' but got '%s' ", want, got)
			}
		})
	}
}
