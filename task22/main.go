package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).MulRange(1, 50)
	b := new(big.Int).MulRange(1, 10)

	dv := new(big.Int)

	// Mul - умножение
	// Add - сложение
	// Sub - вычитание
	fmt.Printf("%v\n+\n%v\n=\n%v", a, b, dv.Div(a, b))
}