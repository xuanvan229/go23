package utils

import (
	"fmt"
	"strconv"
)

// ParseInt parse a list of string to integer
func ParseInt(args []string) ([]int, error) {
	var result []int

	for _, str := range args {
		// Attempt to parse the string to an integer
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %s", str)
		}

		// Check if the parsed integer is the same as the original string
		// This is to handle cases where strings like "1.23" are mistakenly parsed as integers.
		if strconv.Itoa(num) != str {
			return nil, fmt.Errorf("invalid integer: %s", str)
		}

		// Append the valid integer to the result list
		result = append(result, num)
	}

	return result, nil
}

// ParseFloat parse a list of string to float
func ParseFloat(args []string) ([]float64, error) {
	var result []float64

	for _, str := range args {
		// Attempt to parse the string to a float
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float or integer: %s", str)
		}

		// Check if the parsed float is the same as the original string
		// This is to handle cases where strings like "123.00" are mistakenly parsed as floats.
		if strconv.FormatFloat(num, 'f', -1, 64) != str {
			return nil, fmt.Errorf("invalid float or integer: %s", str)
		}

		// Append the valid float to the result list
		result = append(result, num)
	}

	return result, nil
}

// ParseString parse a list of string to string, if the string is not valid return error
func ParseString(args []string) ([]string, error) {

	var result []string
	for _, str := range args {
		if isNumber(str) {
			return nil, fmt.Errorf("invalid string: %s", str)
		}
		result = append(result, str)
	}
	return result, nil
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
