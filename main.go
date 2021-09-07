package main

import (
	"billboard-scraper/auth"
	"billboard-scraper/playlist-builder"
	"fmt"
	//"billboard-scraper/scraper"
)

func main() {
	// var songList []string
	// songList = scraper.GetSongList()

	//fmt.Println(songList)

	client, ctx := authorizeUser.AuthUser()
	fmt.Println(client)
	fmt.Println(ctx)
	playlist.BuildPlaylist(client, ctx)
}
