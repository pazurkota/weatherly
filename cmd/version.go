package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return current app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("weatherly v0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
