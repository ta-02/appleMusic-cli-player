package cmd

import (
	"fmt"
	"strings"

	"github.com/andybrewer/mack"
)

type songInfo struct {
	trackName  string
	artistName string
}

type fullSongInfo struct {
	trackName     string
	artistName    string
	trackDuration string
	trackPosition string
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
		return songInfo{}, nil
	}

	return songInfo{
		trackName:  info[0],
		artistName: info[1],
	}, nil
}

func escapeQuotes(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}
