package main

import "fmt"

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	a, b := 1, 2
	sum := 0

	for {
		if a > 4e6 {
			break
		}
		if isEven(a) {
			sum += a
		}
		a, b = b, a+b
	}

	fmt.Println(sum)
}
