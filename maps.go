package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	ret := make(map[string]int)
	for _, key := range arr {
		if _, ok := ret[key]; ok {
			ret[key] = ret[key] + 1
		} else {
			ret[key] = 1
		}
	}

	return ret
}

func main() {
	wc.Test(WordCount)
}
