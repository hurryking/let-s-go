package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	statics := make(map[string]int)

	for _, v := range fields {
		_, ok := statics[v]
		if ok {
			statics[v]++
		} else {
			statics[v] = 1
		}

	}

	return statics

}

func main() {
	wc.Test(WordCount)
}
