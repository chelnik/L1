package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

//функция для разворота слайса
func reverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	str1 := "snow dog sun"
	//str2 := "The quick brown fox jumps over the lazy dog"
	mySlice := strings.Split(str1, " ")
	fmt.Println(mySlice)
	reverseSlice(mySlice)
	fmt.Println(mySlice)

}
