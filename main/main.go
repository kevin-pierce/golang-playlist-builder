package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%d", len(os.Args))
	fmt.Println("What artist would you like to search?")
}
