package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("invalid nums of arguments ('action' 'path')")
		return
	}

	action, err := NewAction(args[0])
	if err != nil {
		fmt.Println("invalid action ('create', 'read', 'delete')")
		return
	}

	path := args[1]
	err = action.Do(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
