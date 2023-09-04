package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("loremipsum.txt") // Type your file name here
	if err != nil {
		fmt.Println("error: ", err)
	}
	text := string(f)

	words := strings.Fields(text)
	wordCount := make(map[string]int)
	for _, word := range words {
		word = strings.ReplaceAll(word, ",", "")
		wordCount[word] += 1
	}

	for key, value := range wordCount {
		fmt.Printf("%s: %d\n", key, value)
	}
	fmt.Println("Total word count is:", len(words))
}
