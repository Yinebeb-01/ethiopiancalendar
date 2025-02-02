package cmd

import (
	"fmt"
	"github.com/yinebebt/ethiocal/handler"
	"os"

	"github.com/spf13/cobra"
)

var cli bool

var rootCmd = &cobra.Command{
	Use:   "ethiocal",
	Short: "Ethiopian Calendar (ባሕረ-ሐሳብ) and date converter",
	Long: `ethiocal is used to get fasting and holidays date with in a year based on Ethiopian Orthodox 
church calendar. It also do date conversion between Ethiopian and Gregorian in the format yy-mm-dd.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cli {
			fmt.Println("Running in server mode")
			handler.Init()
		}
	},
}

func Execute() {
	if len(os.Args) == 1 || (len(os.Args) == 2 && os.Args[1] == "--cli") {
		_ = rootCmd.Help()
		return
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&cli, "cli", true, "Run ethiocal in CLI mode")

	rootCmd.AddCommand(bahirCmd)
	rootCmd.AddCommand(convertCmd)
}
