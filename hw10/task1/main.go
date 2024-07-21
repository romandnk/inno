package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
Реализуйте BFS алгоритм в представлении матрицы стоимости
*/

const pathNotExist string = "∞"

type Graph struct {
	matrix [][]string
}

func NewGraph(matrix [][]string) (*Graph, error) {
	if len(matrix) != 0 {
		for i, row := range matrix {
			if len(row) != len(matrix) {
				return nil, fmt.Errorf("matrix must be square")
			}
			for j, v := range row {
				if v != pathNotExist {
					_, err := strconv.Atoi(v)
					if err != nil {
						return nil, fmt.Errorf("matrix can contain only number or ∞: position (%d,%d)", i, j)
					}
				}
			}
		}
	}

	return &Graph{matrix: matrix}, nil
}

func (g *Graph) BFS(start int) []int {
	if len(g.matrix) == 0 {
		return []int{}
	}

	path := make([]int, 0, len(g.matrix))
	visited := make([]bool, len(g.matrix))

	queue := make([]int, 0, len(g.matrix))

	queue = append(queue, start)
	path = append(path, start)
	visited[start] = true

	vertex := start

	for len(queue) != 0 {

		row := g.matrix[vertex]

		for i, v := range row {
			if !visited[i] && v != pathNotExist {
				queue = append(queue, i)
				visited[i] = true
				path = append(path, i)
			}
		}

		queue = queue[1:]
		if len(queue) != 0 {
			vertex = queue[0]
		}
	}

	return path
}

func main() {
	// 0 - 1 - 3
	// | /     |
	// 2 ------|
	//
	// 	 0 1 2 3
	// 0 ∞ 5 1 ∞
	// 1 5 ∞ 7 4
	// 2 1 7 ∞ 2
	// 3 ∞ 4 2 ∞
	g, err := NewGraph([][]string{
		{pathNotExist, "3", "1", pathNotExist},
		{"5", pathNotExist, "7", "4"},
		{"1", "7", pathNotExist, "2"},
		{pathNotExist, "4", "2", pathNotExist},
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(g.BFS(0)) // 0 1 2 3
	fmt.Println(g.BFS(1)) // 1 0 2 3
	fmt.Println(g.BFS(2)) // 2 0 1 3
	fmt.Println(g.BFS(3)) // 3 1 2 0
}
