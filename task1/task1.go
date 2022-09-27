package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) String() string {
	return fmt.Sprintf("name: %s age: %d", h.Name, h.Age)
}

type Action struct {
	Human // встраиваем поле Human в структуру action
}

// теперь action тоже реализует интерфейс stringer и может использовать метод string
func main() {
	fmt.Println(Action{Human{
		Name: "Max",
		Age:  25,
	}})
}
