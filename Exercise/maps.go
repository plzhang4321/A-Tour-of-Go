package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var StringArray []string = strings.Fields(s)
	Woc := make(map[string]int)
	for _, word := range StringArray {
		_, ok := Woc[word]
		if ok {
			Woc[word] += 1
		} else {
			Woc[word] = 1
		}
	}
	return Woc
}

func main() {
	wc.Test(WordCount)
}
