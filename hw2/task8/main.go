package main

/* 8. Напишите функцию-дженерик IsEqualArrays для comparable типов, которая сравнивает два неотсортированных массива.
Функция выдает булевое значение как результат. true - если массивы равны, false - если нет.
Массивы считаются равными, если в элемент из первого массива существует в другом, и наоборот.
Вне зависимости от расположения.
*/

func IsEqualArrays[T comparable](first, second []T) bool {
	firstElems := make(map[T]int, len(first))
	for _, elem := range first {
		firstElems[elem]++
	}

	for _, elem := range second {
		if _, ok := firstElems[elem]; !ok {
			return false
		} else {
			firstElems[elem]--
		}
	}

	return true
}
