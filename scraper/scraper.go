package scraper

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func GetSongList() []string {
	htmlDoc, err := GetHTML()
	fmt.Println(htmlDoc)

	if err != nil {
		log.Fatalln(err)
	}
	songs, err := GetSongs(htmlDoc)
	fmt.Println(songs)
	return songs
}

func GetHTML() (*goquery.Document, error) {
	resp, err := http.Get("https://www.billboard.com/charts/hot-100")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetSongs(doc *goquery.Document) ([]string, error) {
	var songList []string

	doc.Find(".chart-list__element .display--flex").Each(func(i int, s *goquery.Selection) {
		songTitle := s.Find(".chart-element__information__song").Text()
		songArtist := s.Find(".chart-element__information__artist").Text()

		// Only search for first artists in list of artists
		splitExp := regexp.MustCompile(`&|Featuring| X `)
		firstArtist := splitExp.Split(songArtist, -1)[0]

		songInfo := songTitle + " " + firstArtist
		//fmt.Println(songInfo)

		songList = append(songList, songInfo)
	})
	return songList, nil
}
