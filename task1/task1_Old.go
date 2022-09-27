//package main
//
//import "fmt"
//
//// создаем структуру в которую агрегируем структуру Action
//type Human struct {
//	name     string
//	activity Action
//}
//
//// встраиваем структуру Action чтобы обращаться напрямую
//type AnotherHuman struct {
//	name string
//	Action
//}
//
//type Action struct {
//	age      int
//	position int
//}
//
//func (a *Action) run() {
//	a.position++
//}
//
//func main() {
//	var dave Human = Human{
//		name: "dave",
//	}
//	stive := AnotherHuman{
//		name: "stive",
//		Action: Action{
//			age:      27,
//			position: 3,
//		},
//	}
//
//	dave.activity = Action{age: 24}
//	dave.activity.run()
//
//	stive.run()
//	fmt.Printf("%#v\n%#v", stive, dave)
//}
