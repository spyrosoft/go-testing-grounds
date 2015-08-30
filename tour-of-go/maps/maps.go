package main

//http://tour.golang.org/moretypes/19

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	totalWords := make(map[string]int)
	eachWord := strings.Split(s, " ")
	for _, word := range eachWord {
		totalWords[ word ]++
	}
	return totalWords
}

func main() {
	wc.Test(WordCount)
}