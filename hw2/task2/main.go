package main

import (
	"fmt"
	"sort"
)

/* 2. Подсчет голосов.
Напишите функцию подсчета каждого голоса за кандидата. Входной аргумент - массив с именами кандидатов.
Результативный - массив структуры Candidate, отсортированный по убыванию количества голосов.
Пример.
Вход: ["Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"]
Вывод: [{Ann, 3}, {Kate, 2}, {Peter, 1}, {Helen, 1}]
*/

type Candidate struct {
	Name  string
	Votes int
}

func main() {
	in := []string{"Ann", "Kate", "Peter", "Kate", "Ann", "Ann", "Helen"}
	fmt.Println(countVotes(in))
}

func countVotes(names []string) []Candidate {
	votesByName := make(map[string]int, len(names))
	for _, name := range names {
		votesByName[name]++
	}

	result := make([]Candidate, 0, len(votesByName))
	for name, count := range votesByName {
		result = append(result, Candidate{name, count})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Votes > result[j].Votes
	})

	return result
}
