package main

import "fmt"

type Human struct {
	name     string
	activity Action
}

type Action struct {
	age int
}

func main() {
	var chelovek Human = Human{
		name: "boba",
	}
	chelovek.activity = Action{age: 5}
	fmt.Println(chelovek)
}
