/*
Copyright © 2023 Xuan Van xuanvan229@gmail.com
*/
package cmd

import (
	"flag"
	"fmt"
	"github.com/xuanvan229/go23/exercise-02/utils"
	"os"
)

type FlagType string

const (
	INTEGER FlagType = "int"
	FLOAT   FlagType = "float"
	STRING  FlagType = "string"
)

var FlagList []FlagType = []FlagType{INTEGER, FLOAT, STRING}

func CheckWhatIsCurrentFlag() (FlagType, error) {
	for _, item := range FlagList {
		if t := flag.Lookup(string(item)); t != nil && t.Value.String() == "true" {
			return item, nil
		}

	}

	return "", fmt.Errorf("no flag is set")
}

func Execute() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("requires at least 1 arg(s), only received 0")
		os.Exit(1)
	}

	inputType, err := CheckWhatIsCurrentFlag()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch inputType {
	case INTEGER:
		result, err := utils.ParseInt(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sortedNumber := utils.SortInt(result)
		fmt.Println("OUTPUT: ", utils.ConvertArrayToString(utils.ConvertIntToString(sortedNumber)))
		break
	case FLOAT:
		result, err := utils.ParseFloat(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sortedNumber := utils.SortFloat(result)

		fmt.Println("OUTPUT: ", utils.ConvertArrayToString(utils.ConvertFloatToString(sortedNumber)))
		break
	case STRING:
		result, err := utils.ParseString(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sortedString := utils.SortString(result)
		fmt.Println("OUTPUT: ", utils.ConvertArrayToString(sortedString))
		break
	}

}

func init() {
	flag.Bool("int", false, "For integer input")
	flag.Bool("float", false, "For number input")
	flag.Bool("string", false, "For string input")
}
