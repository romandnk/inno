package main

import (
	"go.uber.org/fx"
	"zoo/internal/app"
)

func main() {
	fx.New(app.NewApp()).Run()
}
