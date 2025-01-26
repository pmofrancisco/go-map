package main

import (
	"fmt"
)

func main() {
	// Option #1
	// colors := make(map[string]string)
	// colors["black"] = "#000"
	// colors["white"] = "#fff"

	// Option #2
	colors := map[string]string{
		"black": "#000",
		"white": "#fff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}