package scraper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type LinkTag struct {
	linkRef  string
	linkText string
}

func GetSongList() []string {
	htmlDoc, err := GetHTML()
	fmt.Println(htmlDoc)

	if err != nil {
		log.Fatalln(err)
	}

	links, err := GetLinks(htmlDoc)
	return links
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

func GetLinks(doc *goquery.Document) ([]string, error) {
	var songList []string

	doc.Find(".chart-list__element .display--flex").Each(func(i int, s *goquery.Selection) {
		songTitle := s.Find(".chart-element__information__song").Text()
		songArtist := s.Find(".chart-element__information__artist").Text()
		fmt.Println(songArtist)
		fmt.Println(songTitle)
		//fmt.Println(s)
	})

	return songList, nil
}
