package main

import (
	"billboard-scraper/auth"
	"billboard-scraper/playlist-builder"
	"billboard-scraper/scraper"
)

func main() {
	// Get list of songs with songName +  FIRST artist name
	songList := scraper.GetSongList()

	// Authorize user for Spotify
	client, ctx := authorizeUser.AuthUser()

	// Build and create playlist on user's account
	playlist.BuildPlaylist(client, ctx, songList)
}
