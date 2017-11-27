// https://go-tour-jp.appspot.com/moretypes/23
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) (result map[string]int) {
	result = make(map[string]int)

	for _, word := range strings.Fields(s) {
		result[word]++
	}

	return
}

func main() {
	wc.Test(WordCount)
}
