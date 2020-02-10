package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	// this is the trick for a fast output
	l := int(math.Sqrt(float64(n)))
	for i := 2; i <= l; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

const limit = 2000000

func main() {
	sum := 0
	for i := 2; i < limit; i++ {
		if isPrime(i) {
			//fmt.Println(i)
			sum += i
		}
	}
	fmt.Println(sum)
}
