package main

import (
	"go.uber.org/fx"

	app "chat/internal/app/fx"
)

func main() {
	fx.New(app.New()).Run()
}
