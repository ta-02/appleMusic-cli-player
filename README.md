<h1 align="center">appleMusic-cli-player</h1>

<img width="622" alt="image" src="https://github.com/user-attachments/assets/e93089ac-8bfb-4c91-ad2e-ce7d80420994">

<h3 align="center">Control your music directly from the terminal</h3>

## Overview

`appleMusic-cli-player` is a command-line interface (CLI) tool designed to control music playback on Apple Music. Built using [Go](https://golang.org/) and the [Cobra](https://github.com/spf13/cobra) CLI library, it allows you to manage your music without leaving the terminal. With commands for playing, pausing, skipping tracks, and more, this tool offers seamless control over your Apple Music experience.

## Usage

- **List of Commands:**

  ```bash
  music [command]
  ```

  - `completion` - Generate the autocompletion script for the specified shell
  - `current` - Display currently playing song information
  - `help` - Get help about any command
  - `next` - Skip to the next track in Apple Music
  - `open` - Open Apple Music
  - `pause` - Pause music playback
  - `play` - Play the current song in Apple Music
  - `playlists` - Choose a playlist to play
  - `shuffle` - Toggle shuffle mode in Apple Music
  - `volume` - Set the volume for Apple Music

- **Flags:**

  ```bash
  -h, --help     Help for music
  -t, --toggle   Help message for toggle
  ```

Use `music [command] --help` for more information about a specific command.

## Run locally

- Clone the repository

  ```bash
  git clone https://github.com/ta-02/appleMusic-cli-player.git
  ```

- Navigate to the project directory

  ```bash
  cd appleMusic-cli-player
  ```

- Build the project

  ```bash
  go build
  ```

- Run the CLI

  ```bash
  ./music
  ```
