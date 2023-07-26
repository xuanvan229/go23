package utils

import (
	"strconv"
	"strings"
)

// ConvertIntToString convert a list of integer to string
func ConvertIntToString(args []int) []string {
	var result []string

	for _, value := range args {
		result = append(result, strconv.Itoa(value))
	}

	return result
}

// ConvertFloatToString convert a list of float to string
func ConvertFloatToString(args []float64) []string {
	var result []string

	for _, value := range args {
		result = append(result, strconv.FormatFloat(value, 'f', -1, 64))
	}

	return result
}

// ConvertArrayToString convert a list of string to string
func ConvertArrayToString(arr []string) string {
	return strings.Join(arr, " ")
}
