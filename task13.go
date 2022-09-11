package main

import "fmt"

func main() {
	a := 80
	b := 35
	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)
}
