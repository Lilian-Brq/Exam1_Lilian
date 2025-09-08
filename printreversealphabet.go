package main

import "github.com/01-edu/z01"

func PrintReverseAlphabet() {
	list := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(list); i++ {
		z01.PrintRune(rune(list[len(list)-1-i]))
	}
}
