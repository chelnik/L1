package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func createNewSet(s []string) (result map[string]int) {
	result = make(map[string]int)
	for _, key := range s {
		result[key] += 1
	}
	return
}

func main() {
	mySlice := []string{"cat", "cat", "dog", "cat", "tree"}
	for str, i := range createNewSet(mySlice) {
		fmt.Println(str, i)
	}
}
