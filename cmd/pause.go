package cmd

import (
	"fmt"
	"log"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause playback of music",
	Long: `Pause command allows you to pause the playback of music using the command line.
This command sends a pause instruction to control music playback.`,
	Run: func(cmd *cobra.Command, args []string) {
		isOpen, err := isMusicOpen()
		if err != nil {
			log.Fatalf("Failed to check if Apple Music is open: %v", err)
		}

		if !isOpen {
			fmt.Println("Apple Music is not open nothing to pause!")
			return
		}

		mack.Tell("Music", "Pause")
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
