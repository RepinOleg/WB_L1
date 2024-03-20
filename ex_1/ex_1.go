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

type Action struct {
	Human
}

func (a *Action) Eat() {
	fmt.Printf("%s are eating", a.Name)
}

func main() {
	Bob := Action{Human{"Bob", 22}}
	Bob.PrintAge()
	Bob.PrintName()
	Bob.Eat()
}
