package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/zlypher/track/interrupt"
	"github.com/zlypher/track/store"
)

func init() {
	rootCmd.AddCommand(interruptCmd)
}

var interruptCmd = &cobra.Command{
	Use:   "interrupt",
	Short: "Creates a new interrupt entry",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("oh no not enough args")
			return
		}

		label := args[0]
		entry := interrupt.Entry{Label: label, Date: time.Now()}
		store.SaveInterrupt(entry)
	},
}
