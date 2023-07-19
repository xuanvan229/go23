package main

import (
	"fmt"
	"os"
	"strings"
)

type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dial_code"`
	Code     string `json:"code"`
}

func main() {
	var countries = LoadCountries()

	args := os.Args[1:]
	var length = len(args)

	if length < 3 {
		fmt.Println("Not enough arguments")
		return
	}

	var countryCode = args[length-1]
	countryCode = strings.ToUpper(countryCode)

	var isValidCountry, _ = CheckIsValidCountry(countries, countryCode)

	if !isValidCountry {
		fmt.Println("Invalid country code")
		return
	}

	var stringArray = args[:length-1]
	var finalName, err = FormatName(stringArray, countryCode, countries)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Output: " + finalName)

}

func FormatName(array []string, countryCode string, countries []Country) (string, error) {
	var firstName = array[0]
	var lastName = array[1]
	var middleNameArray = array[2:]
	var middleName = ""

	for _, value := range middleNameArray {
		middleName += value + " "
	}

	var isEasternNameOrder = CheckEasternNameOrder(countries, countryCode)
	var name string

	if isEasternNameOrder {
		name = lastName + " " + middleName + firstName
	} else {
		name = firstName + " " + middleName + lastName
	}

	return name, nil

}

func CheckIsValidCountry(countries []Country, countryCode string) (bool, Country) {
	for _, value := range countries {
		if value.Code == countryCode {
			return true, value
		}
	}

	return false, Country{}
}

func CheckEasternNameOrder(countries []Country, countryCode string) bool {
	var EasternNameOrder = []string{"VN", "CN", "JP", "KR", "KP", "SG", "HU"}

	var _, country = CheckIsValidCountry(countries, countryCode)

	for _, value := range EasternNameOrder {
		if value == country.Code {
			return true
		}
	}

	return false
}
