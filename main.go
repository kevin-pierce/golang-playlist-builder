package main

import (
	"billboard-scraper/scraper"
	"fmt"
)

func main() {
	var songList []string
	songList = scraper.GetSongList()

	fmt.Println(songList)
}
