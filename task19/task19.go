package main

import (
	"fmt"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	//var str string
	//_, err := fmt.Scan(&str)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(str)
	str := "golang"
	//arr := []rune(str)
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
	// for
}