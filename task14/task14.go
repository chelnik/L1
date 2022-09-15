package main

import (
	"fmt"
)

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func takeType(variable interface{}) {
	fmt.Printf("%T\n", variable)
}

func main() {
	value1 := 5
	value2 := "string"
	value3 := make(chan interface{})
	value4 := false
	takeType(value1)
	takeType(value2)
	takeType(value3)
	takeType(value4)
}
