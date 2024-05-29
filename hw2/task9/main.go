package main

import "golang.org/x/exp/constraints"

/* 9. Реализуйте тип-дженерик Numbers, который является слайсом численных типов.
Реализуйте следующие методы для этого типа:
* суммирование всех элементов,
* произведение всех элементов,
* сравнение с другим слайсом на равность,
* проверка аргумента, является ли он элементом массива, если да - вывести индекс первого найденного элемента,
* удаление элемента массива по значению,
* удаление элемента массива по индексу.
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

func (nums Numbers[T]) Compare() bool {
	if len(nums) == 0 {
		return 0
	}

	var pr T = 1
	for _, num := range nums {
		pr *= num
	}
	return pr
}
