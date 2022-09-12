package main

import "fmt"

func removeIndex(arr []int, pos int) []int {
	// просто создаем новый слайс в куче
	newArr := make([]int, len(arr))
	k := 0
	// выполняем копирование без нужного нам элемента
	for i := 0; i < len(arr); {
		if i != pos {
			newArr[i] = arr[k]
			k++
		} else {
			k++
		}
		i++
	}
	return newArr
}

func removeIndexNew(s []int, index int) []int {
	// создаем новые сегмент
	ret := make([]int, 0)
	// делаем распаковку сегмента и одновременно поэлементное копирование
	ret = append(ret, s[:index]...)
	// до удаляемого элемента и после
	return append(ret, s[index+1:]...)
}
func main() {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	newSlice := removeIndexNew(all, 5)
	fmt.Println(newSlice)
}

// говорит автоматически найти длину массива
//mySlice := [...]int{1, 2, 3, 4, 5}
