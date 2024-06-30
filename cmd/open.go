package cmd

import (
	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open Apple Music",
	Long:  "Open Apple Music",
	Run: func(cmd *cobra.Command, args []string) {
		mack.Tell("Music", "run")
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
