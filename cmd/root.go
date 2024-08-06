package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "music",
	Short: "CLI for controlling music playback",
	Long: `Music is a command-line interface (CLI) tool that allows you to control music playback.
It leverages Cobra, a CLI library for Go, to define commands and manage application flags.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Define flags and configuration settings here.

	// For example, if you need a configuration file flag:
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.music.yaml)")

	// Define local flags specific to this command.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
