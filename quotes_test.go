package goreloaded_test

import (
	"testing"

	"goreloaded"
)

func TestQuotes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test 1",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:     "Test 2",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:     "Test 3",
			input:    "'awesome  '",
			expected: "'awesome'",
		},
		{
			name:     "Test 4",
			input:    "this is the test: ' this can't work '",
			expected: "this is the test: 'this can't work'",
		},
		{
			name:     "Tes 5",
			input:    "this is the test: 'this can't work '",
			expected: "this is the test: 'this can't work'",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := goreloaded.SortQuotes(test.input)
			want := test.expected

			if want != got {
				t.Errorf("Expected '%s' but got '%s' ", want, got)
			}
		})
	}
}
