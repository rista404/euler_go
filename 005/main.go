package main

import (
	"fmt"
)

var divMax = 20

// see https://projecteuler.net/problem=5
func isDiv(n int) bool {
	for i := 1; i <= divMax; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	i := 1
	for {
		if isDiv(i) {
			fmt.Println(i)
			return
		}
		i++
	}
}
