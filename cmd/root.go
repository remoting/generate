package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "igen",
	Short: "generate source",
	Long:  `generate source`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
