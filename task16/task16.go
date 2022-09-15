package main

import (
	"fmt"
	"sort"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quicksort(mySlice []int) {
	left := 0
	right := len(mySlice) - 1

	if len(mySlice) < 2 {
		return
	} else {
		// pivot опорная точка
		pivot := mySlice[right-left/2]
		for left <= right {
			if mySlice[left] >= pivot && mySlice[right] <= pivot {
				mySlice[left], mySlice[right] = mySlice[right], mySlice[left]
				left++
				right--
				continue
			}
			//если эл-нт слева меньше либо равен центральному оставляем его на месте и сдвигаем указатель к центру
			if mySlice[left] <= pivot {
				left++
			}
			if mySlice[right] > pivot {
				right--
			}
		}
	}
	quicksort(mySlice[:right+1])
	quicksort(mySlice[left:])

}
func main() {
	mySlice := []int{
		1, 2, 4, 4, 4, 4, 4, 4, 4, 4, 2, 2, 2, 2, 2, 25, 5, 5, 5, 6, 6, 6, 6, 67, 77, 8, 8, 99, 9, 9, 9, 9, 9, 7, 1000}
	quicksort(mySlice)
	fmt.Println(sort.IntsAreSorted(mySlice))
	fmt.Println(mySlice)
}
