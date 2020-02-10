package main

import (
	"fmt"
)

func chains(start int) int {
	chains := 0
	n := start
	for i := n; i < 1000000; i++ {
		chains++
		if n == 1 {
			return chains
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	return chains
}

func main() {
	maxN := 1
	max := 1

	for i := 1; i < 1000000; i++ {
		chns := chains(i)
		if chns > max {
			maxN = i
			max = chns
		}
	}

	fmt.Println(maxN)
}
