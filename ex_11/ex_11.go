package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	// множество в языке golang
	set1 := map[interface{}]bool{
		5: true,
		1: true,
		3: true,
		2: true,
	}

	set2 := map[interface{}]bool{
		4: true,
		1: true,
		3: true,
		2: true,
	}
	fmt.Println(Intersection(set1, set2))
}

func Intersection(set1, set2 map[interface{}]bool) map[interface{}]bool {
	result := map[interface{}]bool{}
	// идем по первому множеству
	for key := range set1 {
		// если ключ есть во втором множестве, то добавляем этот ключ в result множество
		if set2[key] {
			result[key] = true
		}
	}
	return result
}
