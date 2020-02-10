package main

import (
	"fmt"
	"math"
)

// note: doesn't work for n=1
// but for this usage it's not important
func nDiv(n int) int {
	div := 0
	sqrt := int(math.Sqrt(float64(n)))

	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			div++
		}
	}

	// This did the trick
	// cause we're going up to sqrt(n)
	return div * 2
}

const limit = 500

func main() {
	curr := 1
	nth := 1

	for {
		nth++
		curr += nth

		divs := nDiv(curr)

		if divs > limit {
			fmt.Println(curr)
			break
		}
	}
}
