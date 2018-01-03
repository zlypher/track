package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zlypher/track/interrupt"
	"github.com/zlypher/track/store"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print a statistic of all interrupts",
	Run: func(cmd *cobra.Command, args []string) {
		data := store.LoadInterrupts()
		statistic := interrupt.DoAnalyze(data)
		interrupt.PrintStatistic(statistic)
	},
}
