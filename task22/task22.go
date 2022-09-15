package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
func main() {
	var (
		x big.Int
		y big.Int
		z big.Int
	)
	x.SetString("210000000000000000005", 10)
	y.SetString("210000000000000000000", 10)
	fmt.Println("x:", x.String(), "y:", y.String())
	fmt.Println("Sub", z.Sub(&x, &y))
	fmt.Println("Div", z.Div(&x, &y))
	fmt.Println("Mul", z.Mul(&x, &y))
	fmt.Println("Add", z.Add(&x, &y))

	//	другой способ объявления переменной
	//a := big.NewInt(2e15)
	//b := big.NewInt(2e15)
	//c := big.NewInt(0)
	//fmt.Println("a:", a.String(), "b:", b.String())
	//fmt.Println("Sub", c.Sub(a, b))
	//fmt.Println("Div", c.Div(a, b))
	//fmt.Println("Mul", c.Mul(a, b))
	//fmt.Println("Add", c.Add(a, b))
}

//var x int64
//bigger := big.NewInt(x)
