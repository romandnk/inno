package main

import (
	"go.uber.org/fx"
	"inno/hw11/internal/app"
)

func main() {
	fx.New(app.NewApp()).Run()
}
