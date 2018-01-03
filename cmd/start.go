package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/zlypher/track/interrupt"
	"github.com/zlypher/track/store"
)

func init() {
	rootCmd.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "Creates a new tracking entry",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("oh no not enough args")
			return
		}

		label := args[0]
		entry := interrupt.Entry{Label: label, Date: time.Now()}
		store.SaveTrack(entry)
	},
}
