package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func LoadCountries() []Country {
	fileContent, err := os.Open("countries.json")

	if err != nil {
		log.Fatal(err)
		return []Country{}
	}

	defer fileContent.Close()

	byteResult, _ := io.ReadAll(fileContent)

	var countries []Country

	err = json.Unmarshal(byteResult, &countries)
	if err != nil {
		log.Fatal(err)
		return []Country{}
	}

	return countries
}
