package main

/* 8. Напишите функцию-дженерик IsEqualArrays для comparable типов, которая сравнивает два неотсортированных массива.
Функция выдает булевое значение как результат. true - если массивы равны, false - если нет.
Массивы считаются равными, если в элемент из первого массива существует в другом, и наоборот.
Вне зависимости от расположения.
*/

func IsEqualArrays[T comparable](first, second []T) bool {
	if len(first) != len(second) {
		return false
	}

	firstElems := make(map[T]int, len(first))
	for _, elem := range first {
		firstElems[elem]++
	}

	secondElems := make(map[T]int, len(second))
	for _, elem := range second {
		secondElems[elem]++
	}

	for elem, count := range firstElems {
		if count != secondElems[elem] {
			return false
		}
	}

	return true
}
