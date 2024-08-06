package cmd

import (
	"fmt"
	"strconv"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Set the volume for Apple Music",
	Long: `Set the volume for Apple Music. Usage example:

music volume 50
This will set the Apple Music volume to 50%.`,

	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]
		vol, err := strconv.Atoi(input)
		if err != nil || vol < 0 || vol > 100 {
			fmt.Println("Please enter a number from 1-100")
			return
		}

		isOpen, err := isMusicOpen()
		if err != nil {
			fmt.Println("Error checking if Apple Music is open:", err)
			return
		}

		if !isOpen {
			fmt.Println("Apple Music is not open!")
			return
		}

		script := fmt.Sprintf("set sound volume to %d", vol)
		if _, err := mack.Tell("Music", script); err != nil {
			fmt.Println("Error setting volume:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)
}
