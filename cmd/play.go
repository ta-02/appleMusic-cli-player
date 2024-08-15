package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/andybrewer/mack"
	"github.com/qeesung/image2ascii/convert"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play current song in Apple Music",
	Long: `Plays the current song in Apple Music. 
	If Apple Music is not open, it will be launched.`,
	Run: func(cmd *cobra.Command, args []string) {
		isOpen, err := isMusicOpen()
		if err != nil {
			log.Fatalf("Failed to check if Apple Music is open: %v", err)
		}

		if !isOpen {
			mack.Tell("Music", "run")
		}

		mack.Tell("Music", "Play")
		info, err := getCurrentSongInfo()
		if err != nil {
			log.Printf("Error getting current song info: %v", err)
			return
		}

		if info.trackName == "" {
			fmt.Println("No song currently playing.")
			return
		}
		fmt.Printf("Now Playing: %s\nBy: %s\n", info.trackName, info.artistName)

		// Run the AppleScript to get the album art
		newCmd := exec.Command("osascript", "scripts/get_album_art.scpt")
		err = newCmd.Run()
		if err != nil {
			log.Fatalf("Failed to run AppleScript: %v", err)
		}

		// Convert the image to ASCII
		convertOptions := convert.DefaultOptions
		converter := convert.NewImageConverter()
		asciiArt := converter.ImageFile2ASCIIString("scripts/tmp.jpg", &convertOptions)

		// Print the ASCII art
		fmt.Println(asciiArt)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
