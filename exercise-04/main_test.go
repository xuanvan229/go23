package main

import "testing"

type TestCase struct {
	value    string
	result   int
	expected int
}

func TestNumbDifferentInteriors(t *testing.T) {
	tests := []TestCase{
		{
			value:    "ab123uii1j01mmh024zxc01bbb900qq000001",
			expected: 4,
		},
		{
			value:    "a1b2b3b4b5n1b2",
			expected: 5,
		},
		{
			value:    "amdjskjashdkjaskdjhaskdha",
			expected: 0,
		},
		{
			value:    "1234567890",
			expected: 1,
		},
		{
			value:    "1234567890a",
			expected: 1,
		},
		{
			value:    "a1234567890",
			expected: 1,
		},
		{
			value:    "a12345712312,67890a",
			expected: 2,
		},
	}

	for _, test := range tests {
		result := NumbDifferentInteriors(test.value)
		if result != test.expected {
			t.Errorf("NumbDifferentInteriors of %v Expected %d, got %d", test.value, test.expected, result)
		}
	}
}
