package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	sliceee := strings.Fields(s)
	imTheMap := make(map[string]int)

	for _, v := range sliceee {
		imTheMap[v]++
	}

	return imTheMap
}

func main() {
	wc.Test(wordCount)
}
