package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

type songInfo struct {
	trackName  string
	artistName string
}

func getCurrentSongInfo() (songInfo, error) {
	script := `
		if player state is playing then
			set trackName to name of current track
			set artistName to artist of current track
			return trackName & "," & artistName
		end if
	`
	result, err := mack.Tell("Music", script)
	if err != nil {
		return songInfo{}, fmt.Errorf("error fetching song info: %w", err)
	}
	info := strings.Split(result, ",")
	if len(info) < 2 {
		return songInfo{}, nil // No song currently playing
	}
	return songInfo{
		trackName:  info[0],
		artistName: info[1],
	}, nil
}

func isMusicOpen() (bool, error) {
	script := `
		if it is running then
			return true
		else
			return false
		end if
	`
	result, err := mack.Tell("Music", script)
	if err != nil {
		return false, fmt.Errorf("error checking if Apple Music is open: %w", err)
	}
	return result == "true", nil
}

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
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
