package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is experimental repo, have fun!")
	for {
	}
}

func forTest(a, b int) int {
	r := 10
	r = a + b*r
	if r == 10 {
		return r
	}
	return r
}
