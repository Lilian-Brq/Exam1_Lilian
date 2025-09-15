package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	start := false
	word := []rune{}
	for _, chr := range s {
		if chr == ' ' || chr == '\t' || chr == '\n' {
			if start {
				break
			}
			continue
		}
		start = true
		word = append(word, chr)
	}
	return string(word) + "\n"
}
