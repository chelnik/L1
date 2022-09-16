package main

import (
	"log"
	"math/rand"
	"time"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}
// минус 1 использование глобальной переменной
var justString string

// создаем строку размером size
func createHugeString(size int) string {
	res := make([]byte, size)
	// создаем рандом
	rand.Seed(time.Now().UnixNano())
	// заполняем слайс рун
	for n := range res {
		res[n] = byte(rand.Intn(70) + 48)
	}
	return string(res)
}
func someFunc() {
	// неговорящее название переменной
	v := createHugeString(1 << 10) //1024
	// никак не контролируется правая граница слайса, на примере с размером строки < 100, словим панику по выходу
	// за границу
	switch {
	case 100 <= len(v)-1:
		justString = v[:100]
	default:
		log.Println("Слишком маленькая строка  slice bounds out of range [:100]")
	}
}

func main() {
	someFunc()
	println(justString)
}
