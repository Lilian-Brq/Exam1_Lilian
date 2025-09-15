package main

import (
	"fmt"
)

func CountAlpha(s string) int {
	list := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	counter := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(list); j++ {
			if s[i] == list[j] {
				counter++
			}
		}
	}
	return counter
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
