package main

import "fmt"

func partition(slice []int) int {
	// selecting pivot element (bar) as the first element of slice
	bar := slice[0]

	// set left and right pointers to the beginning and end of slice
	left, right := 1, len(slice)-1
	// пока левый указатель меньше или равен правому
	for left <= right {

		// если левый указатель меньше или равен правому и
		// slice[left] меньше опорного элемента - смещаем левый указатель вправо
		for left <= right && slice[left] < bar {
			left++
		}
		// если левый указатель меньше или равен правому и
		// slice[left] больше опорного элемента - смещаем правый указатель влево
		for left <= right && slice[right] > bar {
			right--
		}
		// если left все еще меньше right - мы нашли пару числе, которые
		// находятся не на своих позициях, меняем их местами и продвигаем
		// указатели к центру
		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}
	}
	// обмениваем опорный элемент с элементом на позиции right
	slice[0], slice[right] = slice[right], slice[0]
	return right
}

// func recursively sorts the passed slice slise using quick sort method
func quickSort(slice []int) {

	//if len 0 or 1 - there is nothing to sort
	if len(slice) <= 1 {
		return
	}

	// call func for spliting slice and getting index of the reference element
	bar := partition(slice)

	// recursively call quickSort for left and right parts of slice
	quickSort(slice[:bar])
	quickSort(slice[bar+1:])
}

func main() {
	slice := []int{3, 5, 1, 2, 7, 4, 3, 9, 7, 10, 0}

	quickSort(slice)
	fmt.Println(slice)

}
