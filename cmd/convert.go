package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yinebebt/ethiocal/v1/dateconverter"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert dates between Ethiopian and Gregorian calendars",
}

var gtoeCmd = &cobra.Command{
	Use:   "gtoe [year] [month] [day]",
	Short: "Convert Gregorian date to Ethiopian date",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := strconv.Atoi(args[0])
		month, _ := strconv.Atoi(args[1])
		day, _ := strconv.Atoi(args[2])

		etDate, err := dateconverter.Ethiopian(year, month, day)
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}

		fmt.Printf("\nGregorian Date: %04d-%02d-%02d\n", year, month, day)
		fmt.Println("Converted Ethiopian Date:", etDate.Format("2006-01-02"))
	},
}

var etogCmd = &cobra.Command{
	Use:   "etog [year] [month] [day]",
	Short: "Convert Ethiopian date to Gregorian date",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		year, _ := strconv.Atoi(args[0])
		month, _ := strconv.Atoi(args[1])
		day, _ := strconv.Atoi(args[2])

		gregDate, err := dateconverter.Gregorian(year, month, day)
		if err != nil {
			fmt.Println("Error converting date:", err)
			return
		}
		fmt.Printf("\nEthiopian Date: %04d-%02d-%02d\n", year, month, day)
		fmt.Println("Converted Gregorian Date:", gregDate.Format("2006-01-02"))
	},
}

func init() {
	convertCmd.AddCommand(gtoeCmd)
	convertCmd.AddCommand(etogCmd)
}
