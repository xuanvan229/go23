/*
Copyright © 2023 Xuan Van xuanvan229@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xuanvan229/go23/exercise-02/utils"
	"os"
)

type FlagType string

const (
	INTEGER FlagType = "int"
	NUMBER  FlagType = "number"
	STRING  FlagType = "string"
)

var FlagList []FlagType = []FlagType{INTEGER, NUMBER, STRING}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exercise-02",
	Short: "Sorts the input array in ascending order.",
	Long:  `Sorts the input array in ascending order.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		inputType, err := CheckWhatIsCurrentFlag(cmd)
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
		case NUMBER:
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
	},
}

func CheckWhatIsCurrentFlag(cmd *cobra.Command) (FlagType, error) {
	for _, flag := range FlagList {
		if cmd.Flags().Changed(string(flag)) {
			return flag, nil
		}
	}

	return "", fmt.Errorf("No flag is set")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("int", "i", false, "For integer input")
	rootCmd.Flags().BoolP("number", "n", false, "For number input")
	rootCmd.Flags().BoolP("string", "s", false, "For string input")

}
