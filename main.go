package main

import (
	"billboard-scraper/scraper"
	"fmt"
	"os"
)

func main() {
	fmt.Println("%d", len(os.Args))
	fmt.Println("What artist would you like to search?")

	scraper.GetHTML()
}
