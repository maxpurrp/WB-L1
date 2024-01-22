package main

import "fmt"

func main() {
	// creating map for saving unique values
	elemsMap := make(map[string]struct{})

	elems := []string{"cat", "cat", "dog", "cat", "tree"}

	// add elements in map with struct(her size is 0)
	for _, element := range elems {
		elemsMap[element] = struct{}{}
	}
	// Создаем слайс для вывода уникальных элементов
	// create slise for unique elements
	uniqueElem := make([]string, 0, len(elemsMap))
	for key := range elemsMap {
		uniqueElem = append(uniqueElem, key)
	}

	// Выводим уникальные элементы множества (set)
	fmt.Println("Unique elements:", uniqueElem)
}
