package main

type Action string

const (
	Create Action = "create"
	Read   Action = "read"
	Delete Action = "delete"
)

func (a Action) String() string {
	return string(a)
}
