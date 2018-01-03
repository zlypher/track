package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "track",
	Short: "Track helps you keep track of tasks and interrupts",
}

var (
	appVersion string
)

// Execute executes the root command
func Execute(version string) {
	appVersion = version

	rootCmd.Execute()
}
