package cmd

import (
	"fmt"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

var shuffleCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "Toggle shuffle mode in Apple Music",
	Long: `This command toggles the shuffle mode in Apple Music.
If shuffle is currently enabled, this command will disable it, and vice versa.

Example usage:
  music-cli shuffle`,
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

		if _, err := mack.Tell("Music", "set shuffle enabled to not shuffle enabled"); err != nil {
			fmt.Println("Error toggling shuffle mode:", err)
			return
		}

		fmt.Println("Shuffle mode has been toggled")
	},
}

func init() {
	rootCmd.AddCommand(shuffleCmd)
}
