package main

import (
	"fmt"
	"strings"
	"sync"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

func checkStringUniq(s string) bool {
	arrRune := []rune(strings.ToLower(s))
	myMap := make(map[rune]bool)
	for _, sym := range arrRune {
		// ok если элемент существует в мапе
		if _, ok := myMap[sym]; ok {
			return false
		} else {
			myMap[sym] = true
		}
	}
	sync.RWMutex{}
	return true
}
func main() {
	fmt.Println(checkStringUniq("abcd"), "— true")
	fmt.Println(checkStringUniq("abCdefAaf"), "— false")
	fmt.Println(checkStringUniq("aabcd"), "— false")
}
