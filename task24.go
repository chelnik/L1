package main

import "fmt"

type Point struct {
	x int
	y int
}

func newThing() *Point {
	return &Point{x: 10, y: 5}
	//	либо
	//	return &Point{ 10,  5}

}

func (p *Point) takeDistance() {
	fmt.Println(p.x - p.y)
}
func main() {
	dot := newThing()
	dot.takeDistance()
}

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором
