package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/andybrewer/mack"
	"github.com/spf13/cobra"
)

func showPlaylist() ([]string, error) {
	script := `
    set playlistNames to name of playlists
    return playlistNames
  `
	result, err := mack.Tell("Music", script)
	if err != nil {
		return nil, fmt.Errorf("error fetching playlists: %w", err)
	}
	playlistNames := strings.Split(result, ", ")
	return playlistNames, nil
}

func playPlaylist(playlist string) error {
	fmt.Printf("Attempting to play playlist: %s \n", playlist)
	script := fmt.Sprintf(`play playlist named "%s"`, playlist)
	_, err := mack.Tell("Music", script)
	if err != nil {
		return fmt.Errorf("error playing the playlist: %w", err)
	}
	return nil
}

func escapeQuotes(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}

var playlistsCmd = &cobra.Command{
	Use:   "playlists",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(0, 1),
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

		if len(args) == 0 {
			playlistNames, err := showPlaylist()
			if err != nil {
				fmt.Println("Error gettings playlust names:", err)
				return
			}

			if len(playlistNames) == 0 {
				fmt.Println("No playlists found.")
				return
			}

			fzfCmd := exec.Command("fzf")
			fzfCmd.Stdin = strings.NewReader(strings.Join(playlistNames, "\n"))
			fzfCmd.Stderr = os.Stderr

			playlistBytes, err := fzfCmd.Output()
			if err != nil {
				fmt.Println("Error running fzf:", err)
				return
			}

			selectedPlaylist := strings.TrimSpace(string(playlistBytes))
			errr := playPlaylist(selectedPlaylist)
			if errr != nil {
				fmt.Println("Error playing the chosen playlist:", err)
				return
			}
			fmt.Println("%s has been played", selectedPlaylist)
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

		} else {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(playlistsCmd)
}
