package main

import (
	"fmt"
)

func main() {
	number := 3
	bit := 1
	choice := 1
	if choice == 1 {
		fmt.Println(number | (1 << bit))
	} else {
		// &^ исключающее и (~)
		fmt.Println(number &^ (1 << bit))
	}

}
