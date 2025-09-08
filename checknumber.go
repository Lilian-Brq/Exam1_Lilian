package main

import (
	"fmt"
)

func CheckNumber(arg string) bool {
	list := "0123456789"
	for i := 0; i < len(arg); i++ {
		for j := 0; j < len(list); j++ {
			if arg[i] == list[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}
