package main

import (
	"billboard-scraper/auth"
	"billboard-scraper/playlist-builder"
	"billboard-scraper/scraper"
)

func main() {
	songList := scraper.GetSongList()

	client, ctx := authorizeUser.AuthUser()
	playlist.BuildPlaylist(client, ctx, songList)
}
