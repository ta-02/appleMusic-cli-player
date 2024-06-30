package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func showPlaylist() {

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
		fmt.Println("playlists called")
		if len(args) == 0 {

		} else {

		}
	},
}

func init() {
	rootCmd.AddCommand(playlistsCmd)
}
