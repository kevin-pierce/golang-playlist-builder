package scraper

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type LinkTag struct {
	linkRef  string
	linkText string
}

func GetSongList() []string {
	htmlBody, err := GetHTML()

	//fmt.Print(htmlBody)
	if err != nil {
		log.Fatalln(err)
	}
	var links []string

	ParseHTML(htmlBody)
	//links := ParseHTML(htmlBody)
	return links
}

func GetHTML() (*html.Node, error) {
	resp, err := http.Get("https://www.billboard.com/charts/hot-100")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func ParseHTML(n *html.Node) {
	var rootNode *html.Node
	rootNode = nil

	var recFindRoot func(*html.Node)

	recFindRoot = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "ol" {
			rootNode = n
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if rootNode != nil {
				break
			} else {
				recFindRoot(c)
			}
		}
	}
	recFindRoot(n)
	fmt.Println(rootNode)
}

// func GetLinks(n *html.Node) {

// }

// 	root, err := html.Parse(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var links []Link
// 	var rec func(*html.Node)
// 	rec = func(n *html.Node) {
// 		if n.Type == html.ElementNode && n.Data == "a" {
// 			for _, attr := range n.Attr {
// 				if attr.Key == "href" {
// 					var text string
// 					if n.FirstChild != nil {
// 						text = grabText(n.FirstChild)
// 					}
// 					links = append(links, Link{attr.Val, text})
// 				}
// 			}
// 		}
// 		if n.FirstChild != nil {
// 			rec(n.FirstChild)
// 		}
// 		if n.NextSibling != nil {
// 			rec(n.NextSibling)
// 		}
// 	}
// 	rec(root)

// 	return links, nil
// }
