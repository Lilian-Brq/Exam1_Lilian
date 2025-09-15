package main

import "github.com/01-edu/z01"

func PrintDigits() {
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune('0' + i))
	}
}
