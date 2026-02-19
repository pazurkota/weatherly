package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "weatherly",
	Short: "Simple and beautiful CLI app for weather",
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
