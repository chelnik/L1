package main

import "fmt"

func interect(s1, s2 []int) map[int]string {
	result := make(map[int]string)
	for key := range s1 {
		result[key] = "no"
	}
	for key := range s2 {
		result[key] = "yes"
	}
	return result
}
func main() {
	sliceFirst := []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 9}
	sliceSecond := []int{3, 5, 1, 9}
	// из функции приходит пересечение в виде мапы и мы ее сразу же выводим
	for i, str := range interect(sliceFirst, sliceSecond) {
		fmt.Println(i, str)
	}

}
