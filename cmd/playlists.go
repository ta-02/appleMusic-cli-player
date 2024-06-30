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
			// fzfCmd.Stdout = os.Stdout
			fzfCmd.Stderr = os.Stderr

			playlistBytes, err := fzfCmd.Output()
			if err != nil {
				fmt.Println("Error running fzf:", err)
				return
			}

			selectedPlaylist := strings.TrimSpace(string(playlistBytes))
			fmt.Println("Selected Playlist:", selectedPlaylist)
		} else {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(playlistsCmd)
}
