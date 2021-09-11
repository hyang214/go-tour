package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	tmp := strings.Split(s, " ")
	for _, e := range tmp {
		m[e] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}