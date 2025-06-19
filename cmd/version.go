package cmd

import (
	"igen/pkg"
	"igen/pkg/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "generate version",
	Long:  `生成版本`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.GenExecute(version.NewVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
