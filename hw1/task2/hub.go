package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrInvalidAction = errors.New("invalid action")

type actionFunc func(path string) error

type ActionsHub struct {
	actions map[Action]actionFunc
}

func NewActionsHub() *ActionsHub {
	hub := &ActionsHub{}
	hub.register()
	return hub
}

func (hub *ActionsHub) register() {
	hub.actions = map[Action]actionFunc{
		Create: createFile,
		Read:   readFile,
		Delete: deleteFile,
	}
}

func (hub *ActionsHub) Handle(action, path string) error {
	f, ok := hub.actions[Action(action)]
	if !ok {
		return ErrInvalidAction
	}

	err := f(path)
	if err != nil {
		return err
	}

	return nil
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
