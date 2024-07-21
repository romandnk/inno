package main

import (
	"fmt"
	"inno/hw10/task2/tree23"
)

/*
Реализуйте 2-3-дерево: вставка, поиск, удаление
*/

func main() {
	t := tree23.NewTree()
	keys := []int{1, 2, 4, 5, 6, 9, 10}

	// 		5
	//    /  \
	// 	 2    9
	//  / \  / \
	// 1   4 6 10
	for _, key := range keys {
		t.Insert(key)
	}

	node := t.Search(9)
	for _, ch := range node.GetChildren() {
		fmt.Println(ch.GetVals())
	}
	fmt.Println(t.Search(11) == nil)

	// 		5
	//    /  \
	// 	 2    6
	//  / \    \
	// 1   4   10
	fmt.Println(t.Search(9).GetVals())
	t.Delete(9)
	fmt.Println(t.Search(9) == nil)
	node = t.Search(5)
	for _, ch := range node.GetChildren() {
		fmt.Println(ch.GetVals())
	}
}
