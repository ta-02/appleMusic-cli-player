package cmd

import (
	"fmt"
	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
	"strings"
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
		else
			return ""
		end if
	`
	result, err := mack.Tell("Music", script)
	if err != nil {
		return songInfo{}, fmt.Errorf("error fetching song info: %w", err)
	}
	info := strings.Split(result, ",")
	return songInfo{
		trackName:  info[0],
		artistName: info[1],
	}, nil
}

func playMusic() {
	mack.Tell("Music", "play")
}

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		playMusic()
		info, err := getCurrentSongInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Now Playing: %s \nBy: %s \n", info.trackName, info.artistName)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
