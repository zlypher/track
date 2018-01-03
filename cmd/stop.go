package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/zlypher/track/interrupt"
	"github.com/zlypher/track/store"
)

func init() {
	rootCmd.AddCommand(stopCommand)
}

var stopCommand = &cobra.Command{
	Use:   "stop",
	Short: "Creates a new stop entry",
	Run: func(cmd *cobra.Command, args []string) {
		entry := interrupt.Entry{Label: "stop", Date: time.Now()}
		store.SaveStop(entry)
	},
}
