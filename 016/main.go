package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	b := big.NewInt(2)
	exp := big.NewInt(1000)

	res := big.NewInt(0).Exp(b, exp, nil)

	s := res.String()

	sum := 0

	for _, n := range s {
		i, err := strconv.Atoi(string(n))
		if err != nil {
			panic("Not possible")
		}
		sum += i
	}

	fmt.Println(sum)
}
