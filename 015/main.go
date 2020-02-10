package main

import (
	"fmt"
	"math/big"
)

func fac(n *big.Int) *big.Int {
	if n.Cmp(&big.Int{}) <= 0 {
		return big.NewInt(1)
	}

	res := new(big.Int)
	res.Set(n)
	one := big.NewInt(1)

	return res.Mul(res, fac(n.Sub(n, one)))
}

func ways(n *big.Int) *big.Int {
	two := big.NewInt(2)

	// double n
	dblN := new(big.Int)
	dblN.Set(n)
	dblN.Mul(dblN, two)
	// fac(double n)
	top := fac(dblN)

	facN := fac(n)
	// facN * facN
	btm := facN.Mul(facN, facN)

	return top.Div(top, btm)
}

func main() {
	fmt.Printf("%#v\n", ways(big.NewInt(20)))
}
