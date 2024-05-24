package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrInvalidAction = errors.New("invalid action")

type Action string

const (
	Create Action = "create"
	Read   Action = "read"
	Delete Action = "delete"
)

func NewAction(action string) (Action, error) {
	if action != Create.String() && action != Read.String() && action != Delete.String() {
		return "", ErrInvalidAction
	}

	return Action(action), nil
}

func (a Action) String() string {
	return string(a)
}

func (a Action) Do(path string) error {
	var err error

	switch a {
	case Create:
		err = createFile(path)
	case Read:
		err = readFile(path)
	case Delete:
		err = deleteFile(path)
	}
	return err
}

func createFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func readFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func deleteFile(path string) error {
	return os.Remove(path)
}
