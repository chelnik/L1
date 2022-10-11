package main

import "fmt"

func main() {
	str := "привет"
	//for _, sym := range str {
	//	fmt.Println(string(sym))
	//}
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		fmt.Println(string(arr[i]))
	}
}
