package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
func rounding(num float64) (res float64) {
	res = float64(int(num) / 10 * 10)
	return
}
func main() {
	myMap := make(map[float64][]float64)
	myArr := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, num := range myArr {
		//округляем число (сначала убираем последнюю цифру потом добавляем)
		key := rounding(num)
		myMap[key] = append(myMap[key], num)
	}
	fmt.Println(myMap)
}
