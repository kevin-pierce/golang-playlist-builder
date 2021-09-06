package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type LinkTag struct {
	linkRef  string
	linkText string
}

func GetHTML() (bool, error) {
	resp, err := http.Get("https://www.billboard.com/charts/hot-100")

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(string(body))
	return true, nil

}

// func ParseHTML(r io.Reader) ([]LinkTag, error) {
// 	doc, err := html.Parse(r)

// 	if err != nil {
// 		return nil, err
// 	}
// 	return nil, nil

// 	return doc
// }
