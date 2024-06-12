package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

/* 9. Реализуйте тип-дженерик Numbers, который является слайсом численных типов.
Реализуйте следующие методы для этого типа:
* суммирование всех элементов, +
* произведение всех элементов, +
* сравнение с другим слайсом на равность, +
* проверка аргумента, является ли он элементом массива, если да - вывести индекс первого найденного элемента, +
* удаление элемента массива по значению, +
* удаление элемента массива по индексу. +
*/

type numbers interface {
	constraints.Integer | constraints.Float
}

type Numbers[T numbers] []T

func (nums Numbers[T]) Sum() T {
	var sum T
	for _, num := range nums {
		sum += num
	}
	return sum
}

func (nums Numbers[T]) MultiplyElems() T {
	if len(nums) == 0 {
		return 0
	}

	var pr T = 1
	for _, num := range nums {
		pr *= num
	}
	return pr
}

func (nums *Numbers[T]) RemoveElementByElement(toRemove T) {
	j := 0
	for i := 0; i < len(*nums); i++ {
		el := (*nums)[i]
		if el != toRemove {
			(*nums)[j] = (*nums)[i]
			j++
		}
	}

	*nums = (*nums)[:j]
}

func (nums *Numbers[T]) RemoveElementByIndex(i uint32) {
	if int(i) >= len(*nums) {
		return
	}
	*nums = append((*nums)[:i], (*nums)[i+1:]...)
}

func (nums Numbers[T]) IsElem(toFind T) int {
	for i, el := range nums {
		if el == toFind {
			return i
		}
	}
	return -1
}

func (nums Numbers[T]) Compare(sl []T) bool {
	if len(nums) != len(sl) {
		return false
	}

	firstElems := make(map[T]int, len(nums))
	for _, elem := range nums {
		firstElems[elem]++
	}

	secondElems := make(map[T]int, len(sl))
	for _, elem := range sl {
		secondElems[elem]++
	}

	for elem, count := range firstElems {
		if count != secondElems[elem] {
			return false
		}
	}

	return true
}

func main() {
	nums := Numbers[int]{5, 2, 2, 4, 6, 6}
	fmt.Println("Initial slice:", nums)
	fmt.Println("Multiply:", nums.MultiplyElems())
	fmt.Println("Sum:", nums.Sum())

	nums.RemoveElementByElement(2)
	fmt.Println("Remove element 2:", nums)

	nums.RemoveElementByIndex(0)
	fmt.Println("Remove by index 0:", nums)

	fmt.Println("Index the first element 6:", nums.IsElem(6))

	fmt.Println("Compare:", nums.Compare([]int{6, 5, 2, 2, 4, 6}))
	fmt.Println("Compare:", nums.Compare([]int{6, 4, 6}))
}
