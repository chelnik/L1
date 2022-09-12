package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 12, 15, 19, 25, 29, 33, 35, 36, 39, 41, 1000}
	item := 1000
	fmt.Println(binarySearch(mySlice, item))

}

func binarySearch(mySlice []int, item int) int {
	low := 0
	high := len(mySlice) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := mySlice[mid]
		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
