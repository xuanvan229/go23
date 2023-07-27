package utils

import "testing"

func TestParseInt(t *testing.T) {
	var tests = []struct {
		data     []string
		expected []int
	}{
		{[]string{"1", "2", "3"}, []int{1, 2, 3}},
		{[]string{"1", "2", "33", "4"}, []int{1, 2, 33, 4}},
		{[]string{"1", "2", "-33", "4"}, []int{1, 2, -33, 4}},
	}

	var failTests = []struct {
		data []string
	}{
		{[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10.1"}},
		{[]string{"a", "b", "c"}},
		{[]string{"1", "2", "acb"}},
	}

	for _, value := range tests {
		if output, _ := ParseInt(value.data); !EqualInt(output, value.expected) {
			t.Errorf("Output %v not equal to expected %v", output, value.expected)
		}
	}

	for _, value := range failTests {
		if _, err := ParseInt(value.data); err == nil {
			t.Errorf("Expected error but got none")
		}
	}
}

func TestParseFloat(t *testing.T) {
	var tests = []struct {
		data     []string
		expected []float64
	}{
		{[]string{"1.1", "2.2", "3.3"}, []float64{1.1, 2.2, 3.3}},
		{[]string{"10", "-20", "30.33"}, []float64{10, -20, 30.33}},
	}

	var failTests = []struct {
		data []string
	}{
		{
			[]string{"abc", "1,", "3"},
		},
		{
			[]string{"1", "2", "3", "4", "5", "6", "7.2.2", "8", "9", "10.a"},
		},
	}

	for _, value := range tests {
		if output, _ := ParseFloat(value.data); !EqualFloat(output, value.expected) {
			t.Errorf("Output %v not equal to expected %v", output, value.expected)
		}
	}

	for _, value := range failTests {
		if _, err := ParseFloat(value.data); err == nil {
			t.Errorf("Expected error but got none")
		}
	}

}

func EqualInt(output []int, expected []int) bool {
	if len(output) != len(expected) {
		return false
	}

	for i, value := range output {
		if value != expected[i] {
			return false
		}
	}

	return true
}

func EqualFloat(output []float64, expected []float64) bool {
	if len(output) != len(expected) {
		return false
	}

	for i, value := range output {
		if value != expected[i] {
			return false
		}
	}

	return true
}
