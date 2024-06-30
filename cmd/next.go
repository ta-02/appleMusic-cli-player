package cmd

import (
	"fmt"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		isOpen, err := isMusicOpen()
		if err != nil {
			fmt.Println("Error checking if Apple Music is open:", err)
			return
		}

		if !isOpen {
			fmt.Println("Apple Music is not open!")
			return
		}

		if _, err := mack.Tell("Music", "next track"); err != nil {
			fmt.Println("Error skipping the song:", err)
		}

		info, err := getCurrentSongInfo()
		if err != nil {
			fmt.Printf("Error getting current song info: %v", err)
			return
		}

		if info.trackName == "" {
			fmt.Println("Song Skipped")
			return
		}
		fmt.Printf("Now Playing: %s\nBy: %s\n", info.trackName, info.artistName)
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
