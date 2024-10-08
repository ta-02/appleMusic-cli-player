package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

func showCurrent() (fullSongInfo, error) {
	script := `
		if player state is playing then
			set trackName to name of current track
			set artistName to artist of current track
			set trackDuration to duration of current track
			set trackPosition to player position
			return trackName & "," & artistName & "," & trackDuration & "," & trackPosition
		end if
	`
	result, err := mack.Tell("Music", script)
	if err != nil {
		return fullSongInfo{}, fmt.Errorf("error fetching song info: %w", err)
	}
	info := strings.Split(result, ",")
	if len(info) < 4 {
		return fullSongInfo{}, fmt.Errorf("unexpected result format: %s", result)
	}

	trackDuration, _ := strconv.ParseFloat(info[2], 64)
	trackPosition, _ := strconv.ParseFloat(info[3], 64)

	durationMin := int(trackDuration / 60)
	durationSec := int(trackDuration) % 60

	positionMin := int(trackPosition / 60)
	positionSec := int(trackPosition) % 60

	return fullSongInfo{
		trackName:     info[0],
		artistName:    info[1],
		trackDuration: fmt.Sprintf("%d:%02d", durationMin, durationSec),
		trackPosition: fmt.Sprintf("%d:%02d", positionMin, positionSec),
	}, nil
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Displays currently playing song information",
	Long: `Displays information about the song currently playing in the default music player.
	
	Example:
	  musicapp current`,
	Run: func(cmd *cobra.Command, args []string) {
		info, err := showCurrent()
		if err != nil {
			fmt.Println("Nothing is currently playing")
			return
		}
		fmt.Printf("Now Playing: %s\n", info.trackName)
		fmt.Printf("Artist: %s\n", info.artistName)
		fmt.Printf("Time: %s / %s\n", info.trackPosition, info.trackDuration)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
