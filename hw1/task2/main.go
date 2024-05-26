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

	hub := NewActionsHub()

	action, path := args[0], args[1]
	err := hub.Handle(action, path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
