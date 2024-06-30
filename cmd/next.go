package cmd

import (
	"fmt"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Skip to the next track in Apple Music",
	Long: `Skip to the next track in Apple Music. This command will check if Apple Music is open,
and if it is, it will skip to the next track. After skipping, it will display the 
current song information.

Usage example:

music next
This will skip the currently playing song and display the new song's information.`,
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
