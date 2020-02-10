package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	l := len(str)
	for i := 0; i <= l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	lrg := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			multiple := i * j
			// being super smart here
			// don't check if it's a palindrome if it is not
			// larger than the currently largest
			// also because we start from the tops
			// we're saving a lot of palindrome checking
			if multiple > lrg && isPalindrome(multiple) {
				lrg = multiple
			}
		}
	}
	fmt.Println(lrg)
}
