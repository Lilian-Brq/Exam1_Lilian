package main

import (
	"fmt"
)

func CountChar(str string, c rune) int {
	if str == '' {
		return 0
	}
	if c < 'a' || c > 'z' && c < 'A' || c > 'Z' && c < '0' || c > '9' && c != ' ' {
		return 0
	}
	counter := 0

	for i := 0; i < len(str); i++ {
		if str[i] == string(c)[0] {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}
