package main

import "fmt"

type americanCharger interface {
	charge()
}

// создааем зарядку для мака
type macCharger struct {
	power int
}

// реализуем интерфейс американской зарядки
func (m *macCharger) charge() {
	fmt.Printf("i will charge with power: %d\n", m.power)
}

// реализация русской зарядки
type russianCharger interface {
	chargeInRussian()
}

// создаем переходник
type adapter struct {
	*macCharger
}

// заряжаем с переходником
func (a *adapter) chargeInRussian() {
	fmt.Printf("i will charge In Russian with power: %d\n", a.macCharger.power)
}

func main() {
	chargerFirst := &macCharger{power: 40}
	chargerSecond := adapter{
		macCharger: chargerFirst,
	}
	chargerFirst.charge()
	chargerSecond.chargeInRussian()
}

// Паттерн Адаптер (Adapter) предназначен преобразования одного класса в интерфейс, который он не поддерживает,
// к этому самому интерфейсу без изменения самогу класса. Это достигается путем создания оболочки для текущего класса,
// которая будет, в свою очередь, удовлетворяет нужному интерфейсу
