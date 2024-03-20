package main

import (
	"fmt"
	"log"
)

func main() {
	ch := make(chan string)
	var w int
	fmt.Println("Input amount workers")
	count, err := fmt.Scanln(&w)
	if err != nil || count != 1 {
		log.Fatal("Wrong input")
	}

	var i int
	for i < w {
		go func() {
			fmt.Println(<-ch)
		}()
		i++
	}

	for {
		var input string
		fmt.Scanln(&input)
		ch <- input
	}
}
