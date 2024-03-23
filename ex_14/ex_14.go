package main

import (
	"fmt"
	"reflect"
)

func determineType(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	case chan int:
		fmt.Println("chan int")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown type")
	}
}

func determineType2(variable interface{}) {
	fmt.Println(reflect.TypeOf(variable))
}

func main() {
	var (
		test  interface{} = 10
		test1 interface{} = 10.2
		test2 interface{} = "str"
		test3 interface{} = true
		test4 interface{} = make(chan int)
	)
	determineType(test)
	determineType(test1)
	determineType(test2)
	determineType(test3)
	determineType(test4)
	fmt.Println("___________________________________________")
	determineType2(test)
	determineType2(test1)
	determineType2(test2)
	determineType2(test3)
	determineType2(test4)
}
