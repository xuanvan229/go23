package main

import "testing"

type SCheckEasternName struct {
	CountryCode string
	Expected    bool
}

type SFormatName struct {
	data        []string
	countryCode string
	expected    string
}

type SCheckIsValidCountry struct {
	countryCode string
	expected    bool
}

func TestCheckEasternNameOrder(t *testing.T) {

	var tests = []SCheckEasternName{
		{"VN", true},
		{"US", false},
		{"CN", true},
		{"JP", true},
		{"KR", true},
		{"wrong", false},
	}
	var countries = ReadJson()

	for _, test := range tests {
		if output := CheckEasternNameOrder(countries, test.CountryCode); output != test.Expected {
			t.Errorf("Output %t not equal to expected %t", output, test.Expected)
		}
	}

}

func TestFormatName(t *testing.T) {
	var countries = ReadJson()
	var test = []SFormatName{
		{[]string{"Xuan", "Van", "Hong"}, "VN", "Van Hong Xuan"},
		{[]string{"Xuan", "Van", "Hong"}, "US", "Xuan Hong Van"},
		{[]string{"Xuan", "Van"}, "VN", "Van Xuan"},
		{[]string{"Xuan", "Van"}, "us", "Xuan Van"},
		{[]string{"Xuan", "Van", "Hong", "Jr."}, "VN", "Van Hong Jr. Xuan"},
		{[]string{"Xuan", "Van", "Hong", "Jr."}, "US", "Xuan Hong Jr. Van"},
	}

	for _, value := range test {
		if output, _ := FormatName(value.data, value.countryCode, countries); output != value.expected {
			t.Errorf("Output %s not equal to expected %s with country %s", output, value.expected, value.countryCode)
		}
	}
}

func TestCheckIsValidCountry(t *testing.T) {
	var countries = ReadJson()
	var test = []SCheckIsValidCountry{
		{"VN", true},
		{"US", true},
		{"CN", true},
		{"JP", true},
		{"KR", true},
		{"wrong", false},
		{"", false},
	}

	for _, value := range test {
		if output, _ := CheckIsValidCountry(countries, value.countryCode); output != value.expected {
			t.Errorf("Output %t not equal to expected %t with country %s", output, value.expected, value.countryCode)
		}
	}
}
