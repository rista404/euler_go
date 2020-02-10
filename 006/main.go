package main

import (
	"fmt"
)

func main() {
	max := 100

	sumOfSquares := 0
	for i := 1; i <= max; i++ {
		sumOfSquares += i * i
	}

	sumSquared := 0
	for i := 1; i <= max; i++ {
		sumSquared += i
	}
	sumSquared *= sumSquared

	fmt.Println(sumSquared, sumOfSquares)
	fmt.Println(sumSquared - sumOfSquares)
}
