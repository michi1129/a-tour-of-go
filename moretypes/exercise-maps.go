package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	ws := make(map[string]int)
	for i := 0; i < len(fields); i++ {
		ws[fields[i]] += 1
	}

	return ws
}

func main() {
	wc.Test(WordCount)
}
