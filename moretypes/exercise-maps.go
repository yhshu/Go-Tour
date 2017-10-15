package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	strs := strings.Fields(s) //func Fields(s string) []string
	length := len(strs)
	for i := 0; i < length; i++ {
		(count[strs[i]])++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
