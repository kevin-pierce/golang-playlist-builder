package scraper

import (
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// Returns the main list of songs
func GetSongList() []string {
	htmlDoc, err := GetHTML()

	if err != nil {
		log.Fatalln(err)
	}
	songs := GetSongs(htmlDoc)
	return songs
}

// Makes request to Billboard website, returns html document that can be parsed
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

// Parses document for songs, returns slice of song title + first artist name
func GetSongs(doc *goquery.Document) []string {
	var songList []string

	doc.Find(".chart-list__element .display--flex").Each(func(i int, s *goquery.Selection) {
		songTitle := s.Find(".chart-element__information__song").Text()
		songArtist := s.Find(".chart-element__information__artist").Text()

		// Only search for first artists in list of artists
		splitExp := regexp.MustCompile(`&|Featuring| X | x`)
		firstArtist := splitExp.Split(songArtist, -1)[0]

		songInfo := songTitle + " " + firstArtist
		songList = append(songList, songInfo)
	})
	return songList
}
