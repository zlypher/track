package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zlypher/track/store"
)

func init() {
	rootCmd.AddCommand(locationCmd)
}

var locationCmd = &cobra.Command{
	Use:   "location",
	Short: "Print the location where track stores its data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(store.Location())
	},
}
