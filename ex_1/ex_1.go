package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) PrintAge() {
	fmt.Println(h.Age)
}

func (h *Human) PrintName() {
	fmt.Println(h.Name)
}

// В структуру Action встраиваем структуру Human

type Action struct {
	Human
}

func (a *Action) Eat() {
	fmt.Printf("%s are eating", a.Name)
}

func main() {
	// Инициализируем структуру с вложенной в нее структурой
	Bob := Action{Human{"Bob", 22}}

	// Вызываем методы принадлежащие к стуктуре Human
	Bob.PrintAge()
	Bob.PrintName()

	// Вызываем метод принадлежащий к структуре Action
	Bob.Eat()
}
