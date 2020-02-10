package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fac := -1
	max := 600851475143

	for n := 5; n <= max; n++ {
		if max%n == 0 && isPrime(n) {
			fmt.Println(n)
			fac = n
		}
	}
	fmt.Println(fac)
}
