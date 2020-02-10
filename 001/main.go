package main

import "fmt"

func main() {
	sum := 0
	max := 1000

	for n := 1; n < max; n++ {
		if n%3 == 0 || n%5 == 0 {
			sum += n
		}
	}

	fmt.Println(sum)
}
