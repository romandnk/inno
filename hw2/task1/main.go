package main

import (
	"fmt"
	"sort"
)

/* 1. Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.
Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
Примеры:
1. Если на вход подается только 1 слайс [1, 2, 3, 2], результатом должен быть слайс [1, 2, 3].
2. Вход: 2 слайса [1, 2, 3, 2] и [3, 2], результат - [2, 3].
3. Вход: 3 слайса [1, 2, 3, 2], [3, 2] и [], результат - [].
*/

func main() {
	fmt.Println(findIntersections([]int{1, 2, 3, 2}, []int{3, 2}))
}

func findIntersections(slices ...[]int) []int {
	var result []int

	if len(slices) == 0 {
		return result
	}

	intersections := make(map[int]int)
	for _, slice := range slices {
		seen := make(map[int]struct{})
		for _, el := range slice {
			if _, ok := seen[el]; !ok {
				seen[el] = struct{}{}
				intersections[el]++
			}
		}
	}

	numSlices := len(slices)
	for num, inter := range intersections {
		if inter == numSlices {
			result = append(result, num)
		}
	}

	sort.Ints(result)

	return result
}
